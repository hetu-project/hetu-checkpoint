package config

// DispatcherConfig holds configuration for the dispatcher service
type DispatcherConfig struct {
	DBHost                  string `json:"db_host"`
	DBPort                  int    `json:"db_port"`
	DBUser                  string `json:"db_user"`
	DBPassword              string `json:"db_password"`
	DBName                  string `json:"db_name"`
	HTTPPort                int    `json:"http_port"`
	TCPPort                 int    `json:"tcp_port"`
	EnableReport            bool   `json:"enable_report"`
	ChainGRpcURL            string `json:"chain_grpc_url"`
	CometBFTSvr             string `json:"comet_bft_svr"`
	ChainID                 string `json:"chain_id"`
	EthRpcURL               string `json:"eth_rpc_url"`
	EthChainID              int64  `json:"eth_chain_id"`
	ValidatorStakingAddress string `json:"validator_staking_address"`
	RewardDistrInterval     int    `json:"reward_distr_interval"` // Number of epochs behind the current epoch to start reward distributions
}

// ValidatorConfig holds configuration for the validator service
type ValidatorConfig struct {
	DBHost                  string `json:"db_host"`
	DBPort                  int    `json:"db_port"`
	DBUser                  string `json:"db_user"`
	DBPassword              string `json:"db_password"`
	DBName                  string `json:"db_name"`
	DispatcherTcp           string `json:"dispatcher_tcp"`
	Port                    int    `json:"port"`
	ListenAddr              string `json:"listen_addr"`
	ChainRpcURL             string `json:"chain_rpc_url"`
	StakingTokenAddress     string `json:"staking_token_address"`
	ValidatorStakingAddress string `json:"validator_staking_address"`
	DispatcherURL           string `json:"dispatcher_url"`
	ChainGRpcURL            string `json:"chain_grpc_url"`
	CometBFTSvr             string `json:"comet_bft_svr"`
	ChainID                 string `json:"chain_id"`
}
