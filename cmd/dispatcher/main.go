package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"

	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/crypto"
	"github.com/hetu-project/hetu-checkpoint/logger"
	"github.com/hetu-project/hetu-checkpoint/store"
)

var (
	configFile string
	logLevel   string
	enableDB   bool
	keyFile    string
	keyPwd     string	// password
	keyPair    *crypto.CombinedKeyPair
	validators struct {
		sync.RWMutex
		connections map[net.Conn]bool
		addresses   map[net.Conn]string
	}
	dbClient *store.DBClient
)

func init() {
	validators.connections = make(map[net.Conn]bool)
	validators.addresses = make(map[net.Conn]string)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./config.json)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "INFO", "log level (DEBUG, INFO, WARN, ERROR, FATAL)")
	rootCmd.PersistentFlags().BoolVar(&enableDB, "enable-db", false, "enable database persistence")
	rootCmd.PersistentFlags().StringVar(&keyFile, "keys", "", "path to the key file")
	rootCmd.PersistentFlags().StringVar(&keyPwd, "key-password", "", "password for the key file")
}

var rootCmd = &cobra.Command{
	Use:   "dispatcher",
	Short: "Dispatcher service for BLS signing",
	Long:  `Dispatcher service that coordinates validators for BLS signing operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set log level
		logger.SetLevel(logger.GetLevelFromString(logLevel))
		logger.Info("Setting log level to %s", logLevel)

		// Load configuration
		cfg, err := config.LoadDispatcherConfig(configFile)
		if err != nil {
			logger.Fatal("Failed to load configuration: %v", err)
		}

		// Load key pair if specified
		if keyFile != "" {
			logger.Info("Loading key pair from %s", keyFile)
			keyPair, err = crypto.LoadKeyPair(keyFile, keyPwd)
			if err != nil {
				logger.Fatal("Failed to load key pair: %v", err)
			}
			logger.Info("Loaded key pair with Ethereum address: %s", keyPair.ETH.Address)
		}

		// Initialize database client only if enabled
		if enableDB {
			logger.Info("Database persistence enabled, initializing connection...")
			dbClient, err = store.NewDBClient(store.Config{
				Host:     cfg.DBHost,
				Port:     cfg.DBPort,
				User:     cfg.DBUser,
				Password: cfg.DBPassword,
				DBName:   cfg.DBName,
			})
			if err != nil {
				logger.Fatal("Failed to initialize database client: %v", err)
			}
			defer dbClient.Close()

			// Create database tables
			if err := dbClient.CreateDispatcherTables(); err != nil {
				logger.Fatal("Failed to create database tables: %v", err)
			}
			logger.Info("Database initialized successfully")
		}

		// Start the server
		startServer(cfg)
	},
}

func startServer(cfg *config.DispatcherConfig) {
	httpPort := fmt.Sprintf(":%d", cfg.HTTPPort)
	tcpPort := fmt.Sprintf(":%d", cfg.TCPPort)

	http.HandleFunc("/reqblssign", handleRequest)
	go func() {
		logger.Info("Starting HTTP server on port %d", cfg.HTTPPort)
		if err := http.ListenAndServe(httpPort, nil); err != nil {
			logger.Fatal("Error starting HTTP server: %v", err)
		}
	}()

	listener, err := net.Listen("tcp", tcpPort)
	if err != nil {
		logger.Fatal("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	logger.Info("The dispatcher is listening on TCP port %d, waiting for connections", cfg.TCPPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Error accepting connection: %v", err)
			continue
		}
		go handleValidatorConnection(conn)
	}
}

func handleValidatorConnection(conn net.Conn) {
	// Wait for the validator to send its listening address
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		logger.Error("Error reading validator address: %v", err)
		conn.Close()
		return
	}

	data := string(buffer[:n])
	var validatorAddr string

	// Parse the ADDR: message
	if len(data) > 5 && data[:5] == "ADDR:" {
		validatorAddr = data[5:]
		logger.Info("Validator registered with address: %s", validatorAddr)
	} else {
		logger.Error("Invalid address format from validator: %s", data)
		conn.Close()
		return
	}

	validators.Lock()
	validators.connections[conn] = true
	// Store the validator's listening address
	if validators.addresses == nil {
		validators.addresses = make(map[net.Conn]string)
	}
	validators.addresses[conn] = validatorAddr
	validators.Unlock()

	logger.Info("New validator connected. Total validators: %d", len(validators.connections))

	// Handle connection until it's closed
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logger.Info("Validator disconnected: %v", err)
			validators.Lock()
			delete(validators.connections, conn)
			delete(validators.addresses, conn)
			conn.Close()
			validators.Unlock()
			logger.Info("Validator removed. Remaining validators: %d", len(validators.connections))
			return
		}

		data := string(buffer[:n])
		if data != "ping" {
			logger.Debug("Received data from validator: %s", data)
		}

		// If it's a heartbeat, respond with pong
		if data == "ping" {
			_, err = conn.Write([]byte("pong"))
			if err != nil {
				logger.Error("Failed to send heartbeat response: %v", err)
				validators.Lock()
				delete(validators.connections, conn)
				delete(validators.addresses, conn)
				conn.Close()
				validators.Unlock()
				return
			}
		}
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	logger.Info("Received client request: %s", string(body))

	// Store the request in database if enabled
	var request *store.SignRequest
	if enableDB {
		var err error
		request, err = dbClient.InsertDisSignRequest(string(body))
		if err != nil {
			logger.Error("Failed to store sign request: %v", err)
			// Continue processing even if DB storage fails
		}
	}

	// Create separate connections for each validator
	type ValidatorClient struct {
		Conn net.Conn
		Addr string
	}

	var validatorClients []ValidatorClient

	// Get validator addresses
	validators.RLock()
	for conn := range validators.connections {
		addr := validators.addresses[conn]
		if addr != "" {
			validatorClients = append(validatorClients, ValidatorClient{
				Conn: conn,
				Addr: addr,
			})
		}
	}
	validatorCount := len(validatorClients)
	validators.RUnlock()

	if validatorCount == 0 {
		logger.Warn("No validators connected")
		http.Error(w, "No validators connected", http.StatusInternalServerError)
		return
	}

	// Create channels for collecting responses
	type ValidatorResponse struct {
		Response []byte
		Error    error
	}
	results := make(chan ValidatorResponse, validatorCount)

	// Create new connections for each request
	for _, vc := range validatorClients {
		go func(addr string) {
			// Create a new connection for this request
			reqConn, err := net.Dial("tcp", addr)
			if err != nil {
				results <- ValidatorResponse{Error: fmt.Errorf("connection error: %v", err)}
				return
			}
			defer reqConn.Close()

			// Set timeouts
			reqConn.SetDeadline(time.Now().Add(900 * time.Millisecond))

			// Send request
			_, err = reqConn.Write(body)
			if err != nil {
				results <- ValidatorResponse{Error: fmt.Errorf("write error: %v", err)}
				return
			}

			// Read response
			var responseData []byte
			buffer := make([]byte, 1024)

			for {
				n, err := reqConn.Read(buffer)
				if err != nil {
					if len(responseData) == 0 {
						results <- ValidatorResponse{Error: fmt.Errorf("read error: %v", err)}
					} else {
						results <- ValidatorResponse{Response: responseData}
					}
					return
				}
				responseData = append(responseData, buffer[:n]...)

				// If we received less than buffer size, we've got the complete message
				if n < len(buffer) {
					break
				}
			}

			logger.Debug("Received response from validator: %s", string(responseData))
			results <- ValidatorResponse{Response: responseData}
		}(vc.Addr)
	}

	// Collect responses with timeout
	timeout := time.After(1 * time.Second)
	validResponses := make([][]byte, 0, validatorCount)
	errorCount := 0

	// Wait for all responses or timeout
	for i := 0; i < validatorCount; i++ {
		select {
		case result := <-results:
			if result.Error != nil {
				logger.Error("Validator error: %v", result.Error)
				errorCount++
			} else {
				validResponses = append(validResponses, result.Response)
				logger.Debug("Received valid response: %s", string(result.Response))
			}
		case <-timeout:
			logger.Warn("Timeout reached. Received %d valid responses, %d errors, %d validators did not respond",
				len(validResponses), errorCount, validatorCount-len(validResponses)-errorCount)
			goto respondToClient
		}
	}

respondToClient:
	if len(validResponses) == 0 {
		logger.Error("No valid responses received from validators")
		if enableDB && request != nil {
			_ = dbClient.UpdateDisSignRequestStatus(request.ID, "FAILED")
		}
		http.Error(w, "No valid responses received from validators", http.StatusInternalServerError)
		return
	}

	// Store validator responses if DB is enabled
	if enableDB && request != nil {
		for _, response := range validResponses {
			_, err := dbClient.InsertDisSignResponse(request.ID, "validator-id", hex.EncodeToString(response))
			if err != nil {
				logger.Error("Failed to store validator response: %v", err)
			}
		}
		err := dbClient.UpdateDisSignRequestStatus(request.ID, "COMPLETED")
		if err != nil {
			logger.Error("Failed to update request status: %v", err)
		}
	}

	// Write summary of responses
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"total_validators":    validatorCount,
		"responses_received":  len(validResponses),
		"errors":              errorCount,
		"no_response":         validatorCount - len(validResponses) - errorCount,
		"validator_responses": validResponses,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}