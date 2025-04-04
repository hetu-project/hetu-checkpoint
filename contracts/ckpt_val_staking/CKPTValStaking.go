// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CKPTValStaking

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CKPTValStakingMetaData contains all meta data concerning the CKPTValStaking contract.
var CKPTValStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumStake\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"epochNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"powerSum\",\"type\":\"uint64\"}],\"name\":\"CheckpointSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"name\":\"UnstakeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dispatcherURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"blsPublicKey\",\"type\":\"string\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dispatcherURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"blsPublicKey\",\"type\":\"string\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPATCHER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epochNum\",\"type\":\"uint64\"}],\"name\":\"distributeCheckpointRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"distributedEpochs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"epochToCheckpoint\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"epochNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bitmap\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsMultiSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsAggrPk\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"powerSum\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getTopValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"dispatcherURLs\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"blsPublicKeys\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dispatcherURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsPublicKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activationTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantDispatcherRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantDistributerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantValidatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initiateUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_dispatcherURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_blsPublicKey\",\"type\":\"string\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeDispatcherRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeDistributerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeValidatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockPeriod\",\"type\":\"uint256\"}],\"name\":\"setLockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumStake\",\"type\":\"uint256\"}],\"name\":\"setMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardRate\",\"type\":\"uint256\"}],\"name\":\"setRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeLockPeriod\",\"type\":\"uint256\"}],\"name\":\"setStakeLockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeActivationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epochNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_bitmap\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsMultiSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsAggrPk\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_powerSum\",\"type\":\"uint64\"}],\"name\":\"submitCheckpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"updateValidatorCursor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_dispatcherURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_blsPublicKey\",\"type\":\"string\"}],\"name\":\"updateValidatorInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dispatcherURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsPublicKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405262093a806004556000600655620151806007553480156200002457600080fd5b5060405162005e9438038062005e9483398181016040528101906200004a9190620004b8565b33600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000c05760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000b7919062000525565b60405180910390fd5b620000d181620001d960201b60201c565b5082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160038190555080600581905550620001366000801b336200029d60201b60201c565b50620001697f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c98926336200029d60201b60201c565b506200019c7ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c336200029d60201b60201c565b50620001cf7f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e336200029d60201b60201c565b5050505062000542565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000620002b18383620003a060201b60201c565b6200039557600180600085815260200190815260200160002060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550620003316200040b60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600190506200039a565b600090505b92915050565b60006001600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600033905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620004458262000418565b9050919050565b620004578162000438565b81146200046357600080fd5b50565b60008151905062000477816200044c565b92915050565b6000819050919050565b62000492816200047d565b81146200049e57600080fd5b50565b600081519050620004b28162000487565b92915050565b600080600060608486031215620004d457620004d362000413565b5b6000620004e48682870162000466565b9350506020620004f786828701620004a1565b92505060406200050a86828701620004a1565b9150509250925092565b6200051f8162000438565b82525050565b60006020820190506200053c600083018462000514565b92915050565b61594280620005526000396000f3fe608060405234801561001057600080fd5b506004361061028a5760003560e01c80637a7664601161015c578063a217fddf116100ce578063d547741f11610087578063d547741f146107a2578063e18d4f8b146107be578063ec5ffac2146107da578063f10e00d6146107f8578063f2fde38b14610816578063fa52c7d8146108325761028a565b8063a217fddf146106dd578063a5f68a18146106fb578063a694fc3a14610730578063a8667ce11461074c578063ae5ac92114610768578063c49baebe146107845761028a565b806393a5b1b61161012057806393a5b1b6146105f457806399745318146106275780639dc820a5146106575780639e447fc6146106735780639ef5aee91461068f578063a17c89f9146106ad5761028a565b80637a7664601461053c5780637b0a47ee1461056c5780638ae647df1461058a5780638da5cb5b146105a657806391d14854146105c45761028a565b80633fd8b02f116102005780636d5beac7116101b95780636d5beac7146104a25780637071688a146104be578063715018a6146104dc57806372f702f3146104e657806374a28d6114610504578063779972da146105205761028a565b80633fd8b02f14610408578063499ab37e146104265780634dda27b41461044457806351057fad1461046057806363803b231461047c5780636b11b1eb146104865761028a565b8063295142041161025257806329514204146103725780632e3a96901461038e5780632f2ff15d146103aa57806336568abe146103c6578063372500ab146103e25780633d81380d146103ec5761028a565b806301ffc9a71461028f5780631904bb2e146102bf578063233e9903146102f6578063248a9ca3146103125780632786879f14610342575b600080fd5b6102a960048036038101906102a49190613c52565b610869565b6040516102b69190613c9a565b60405180910390f35b6102d960048036038101906102d49190613d13565b6108e3565b6040516102ed989796959493929190613de9565b60405180910390f35b610310600480360381019061030b9190613ea1565b610ba2565b005b61032c60048036038101906103279190613f04565b610bb4565b6040516103399190613f40565b60405180910390f35b61035c60048036038101906103579190613d13565b610bd4565b6040516103699190613f5b565b60405180910390f35b61038c60048036038101906103879190613d13565b610bec565b005b6103a860048036038101906103a39190613ea1565b610c21565b005b6103c460048036038101906103bf9190613f76565b610c33565b005b6103e060048036038101906103db9190613f76565b610c55565b005b6103ea610cd0565b005b6104066004803603810190610401919061401b565b610e9d565b005b6104106111d5565b60405161041d9190613f5b565b60405180910390f35b61042e6111db565b60405161043b9190613f5b565b60405180910390f35b61045e60048036038101906104599190614132565b6111e1565b005b61047a60048036038101906104759190613ea1565b611505565b005b61048461158f565b005b6104a0600480360381019061049b9190613d13565b61181a565b005b6104bc60048036038101906104b79190613d13565b61184f565b005b6104c6611884565b6040516104d39190613f5b565b60405180910390f35b6104e4611891565b005b6104ee6118a5565b6040516104fb9190614280565b60405180910390f35b61051e6004803603810190610519919061429b565b6118cb565b005b61053a60048036038101906105359190613ea1565b611c4e565b005b61055660048036038101906105519190613d13565b611c60565b6040516105639190613f5b565b60405180910390f35b610574611cac565b6040516105819190613f5b565b60405180910390f35b6105a4600480360381019061059f9190613d13565b611cb2565b005b6105ae611ce7565b6040516105bb91906142d7565b60405180910390f35b6105de60048036038101906105d99190613f76565b611d10565b6040516105eb9190613c9a565b60405180910390f35b61060e60048036038101906106099190613ea1565b611d7b565b60405161061e949392919061457a565b60405180910390f35b610641600480360381019061063c9190613ea1565b61235c565b60405161064e91906142d7565b60405180910390f35b610671600480360381019061066c919061401b565b61239b565b005b61068d60048036038101906106889190613ea1565b612587565b005b610697612599565b6040516106a49190613f40565b60405180910390f35b6106c760048036038101906106c2919061429b565b6125bd565b6040516106d49190613c9a565b60405180910390f35b6106e56125dd565b6040516106f29190613f40565b60405180910390f35b6107156004803603810190610710919061429b565b6125e4565b6040516107279695949392919061463f565b60405180910390f35b61074a60048036038101906107459190613ea1565b6127e0565b005b61076660048036038101906107619190613d13565b612cd9565b005b610782600480360381019061077d9190613ea1565b612d0e565b005b61078c613087565b6040516107999190613f40565b60405180910390f35b6107bc60048036038101906107b79190613f76565b6130ab565b005b6107d860048036038101906107d39190613d13565b6130cd565b005b6107e2613102565b6040516107ef9190613f5b565b60405180910390f35b610800613108565b60405161080d9190613f40565b60405180910390f35b610830600480360381019061082b9190613d13565b61312c565b005b61084c60048036038101906108479190613d13565b6131b2565b6040516108609897969594939291906146b5565b60405180910390f35b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806108dc57506108db82613317565b5b9050919050565b6000806000606080600080600080600a60008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816002015490508160060160009054906101000a900460ff16801561095d575060008260000154115b80156109a85750600860008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020544210155b15610a005760008260010154426109bf9190614770565b9050670de0b6b3a76400008360000154600354836109dd91906147a4565b6109e791906147a4565b6109f19190614815565b826109fc9190614846565b9150505b600860008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549250816000015481836003015484600401856005018660060160009054906101000a900460ff16876007015489848054610a79906148a9565b80601f0160208091040260200160405190810160405280929190818152602001828054610aa5906148a9565b8015610af25780601f10610ac757610100808354040283529160200191610af2565b820191906000526020600020905b815481529060010190602001808311610ad557829003601f168201915b50505050509450838054610b05906148a9565b80601f0160208091040260200160405190810160405280929190818152602001828054610b31906148a9565b8015610b7e5780601f10610b5357610100808354040283529160200191610b7e565b820191906000526020600020905b815481529060010190602001808311610b6157829003601f168201915b50505050509350995099509950995099509950995099505050919395975091939597565b610baa613381565b8060058190555050565b600060016000838152602001908152602001600020600101549050919050565b60086020528060005260406000206000915090505481565b610bf4613381565b610c1e7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c98926826130ab565b50565b610c29613381565b8060078190555050565b610c3c82610bb4565b610c4581613408565b610c4f838361341c565b50505050565b610c5d61350d565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610cc1576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ccb8282613515565b505050565b610cd933613608565b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020154905060008111610d63576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5a90614926565b60405180910390fd5b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020181905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610e08929190614946565b6020604051808303816000875af1158015610e27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4b919061499b565b503373ffffffffffffffffffffffffffffffffffffffff167ffc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe82604051610e929190613f5b565b60405180910390a250565b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015414610f22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1990614a14565b60405180910390fd5b610f4c7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892633611d10565b610f8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8290614aa6565b60405180910390fd5b60405180610100016040528060008152602001428152602001600081526020016000815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152602001600015158152602001600b80549050815250600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040190816110d49190614c97565b5060a08201518160050190816110ea9190614c97565b5060c08201518160060160006101000a81548160ff02191690831515021790555060e08201518160070155905050600b339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f5196ca52b5246aa712ef709c67d65576c4f5ca8aba7d2df5d782cca62cae21d5858585856040516111c79493929190614da5565b60405180910390a250505050565b60045481565b60075481565b61120b7ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c33611d10565b61124a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124190614e52565b60405180910390fd5b6000600960008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16146112d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112cc90614ee4565b60405180910390fd5b6040518060c001604052808a67ffffffffffffffff16815260200189815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018267ffffffffffffffff16815250600960008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151816001015560408201518160020190816114579190614f5f565b50606082015181600301908161146d9190614f5f565b5060808201518160040190816114839190614f5f565b5060a08201518160050160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050508867ffffffffffffffff167f609e08589f9d05d683838b8bb92f105d0796c16700397a452a95b4c936dac34689836040516114f2929190615031565b60405180910390a2505050505050505050565b6000600b80549050905060008111611552576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611549906150a6565b60405180910390fd5b600060065490506000828410611568578261156a565b835b90508281836115799190614846565b61158391906150c6565b60068190555050505050565b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015411611614576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161160b90615143565b60405180910390fd5b600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030154421015611699576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611690906151af565b60405180910390fd5b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015490506000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030181905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401611785929190614946565b6020604051808303816000875af11580156117a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117c8919061499b565b503373ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f758260405161180f9190613f5b565b60405180910390a250565b611822613381565b61184c7ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c82610c33565b50565b611857613381565b6118817f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e82610c33565b50565b6000600b80549050905090565b611899613381565b6118a36000613737565b565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6118f57f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e33611d10565b611934576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161192b90615241565b60405180910390fd5b6000600960008367ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020905060008160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16036119c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119bb906152d3565b60405180910390fd5b600c60008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611a39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a3090615365565b60405180910390fd5b6000611a436137fb565b905060008151905060008111611a8e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a85906153f7565b60405180910390fd5b6000818460050160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16611abe9190614815565b905060005b82811015611c06576000848281518110611ae057611adf615417565b5b60200260200101519050600082905061020081108015611b925750611b91876002018054611b0d906148a9565b80601f0160208091040260200160405190810160405280929190818152602001828054611b39906148a9565b8015611b865780601f10611b5b57610100808354040283529160200191611b86565b820191906000526020600020905b815481529060010190602001808311611b6957829003601f168201915b505050505082613af7565b5b15611bf15783600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002016000828254611be99190614846565b925050819055505b50508080611bfe90615446565b915050611ac3565b506001600c60008767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050505050565b611c56613381565b8060048190555050565b6000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050919050565b60035481565b611cba613381565b611ce47f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892682610c33565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006001600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6060806060806000600b80549050905060008103611ed357600067ffffffffffffffff811115611dae57611dad614ac6565b5b604051908082528060200260200182016040528015611ddc5781602001602082028036833780820191505090505b50600067ffffffffffffffff811115611df857611df7614ac6565b5b604051908082528060200260200182016040528015611e265781602001602082028036833780820191505090505b50600067ffffffffffffffff811115611e4257611e41614ac6565b5b604051908082528060200260200182016040528015611e7557816020015b6060815260200190600190039081611e605790505b50600067ffffffffffffffff811115611e9157611e90614ac6565b5b604051908082528060200260200182016040528015611ec457816020015b6060815260200190600190039081611eaf5790505b50945094509450945050612355565b6000818710611ee25781611ee4565b865b90508067ffffffffffffffff811115611f0057611eff614ac6565b5b604051908082528060200260200182016040528015611f2e5781602001602082028036833780820191505090505b5095508067ffffffffffffffff811115611f4b57611f4a614ac6565b5b604051908082528060200260200182016040528015611f795781602001602082028036833780820191505090505b5094508067ffffffffffffffff811115611f9657611f95614ac6565b5b604051908082528060200260200182016040528015611fc957816020015b6060815260200190600190039081611fb45790505b5093508067ffffffffffffffff811115611fe657611fe5614ac6565b5b60405190808252806020026020018201604052801561201957816020015b60608152602001906001900390816120045790505b509250600080600654905060005b848110801561203557508383105b1561233a5760008582846120499190614846565b61205391906150c6565b90506000600b828154811061206b5761206a615417565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060060160009054906101000a900460ff1680156120fd575060008160000154115b80156121485750600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020544210155b1561232457818c878151811061216157612160615417565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600001548b87815181106121b3576121b2615417565b5b6020026020010181815250508060040180546121ce906148a9565b80601f01602080910402602001604051908101604052809291908181526020018280546121fa906148a9565b80156122475780601f1061221c57610100808354040283529160200191612247565b820191906000526020600020905b81548152906001019060200180831161222a57829003601f168201915b50505050508a878151811061225f5761225e615417565b5b6020026020010181905250806005018054612279906148a9565b80601f01602080910402602001604051908101604052809291908181526020018280546122a5906148a9565b80156122f25780601f106122c7576101008083540402835291602001916122f2565b820191906000526020600020905b8154815290600101906020018083116122d557829003601f168201915b505050505089878151811061230a57612309615417565b5b6020026020010181905250858061232090615446565b9650505b505050808061233290615446565b915050612027565b5082821015612350578188528187528186528185525b505050505b9193509193565b600b818154811061236c57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015411612420576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612417906154da565b60405180910390fd5b61244a7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892633611d10565b612489576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161248090614aa6565b60405180910390fd5b8383600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040191826124da929190615505565b508181600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600501918261252c929190615505565b503373ffffffffffffffffffffffffffffffffffffffff167f7f35b3dd337a65b9cbaf229ba2c1c6e4aa76460db271df1248aa1a605d44ddf2858585856040516125799493929190614da5565b60405180910390a250505050565b61258f613381565b8060038190555050565b7ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c81565b600c6020528060005260406000206000915054906101000a900460ff1681565b6000801b81565b60096020528060005260406000206000915090508060000160009054906101000a900467ffffffffffffffff1690806001015490806002018054612627906148a9565b80601f0160208091040260200160405190810160405280929190818152602001828054612653906148a9565b80156126a05780601f10612675576101008083540402835291602001916126a0565b820191906000526020600020905b81548152906001019060200180831161268357829003601f168201915b5050505050908060030180546126b5906148a9565b80601f01602080910402602001604051908101604052809291908181526020018280546126e1906148a9565b801561272e5780601f106127035761010080835404028352916020019161272e565b820191906000526020600020905b81548152906001019060200180831161271157829003601f168201915b505050505090806004018054612743906148a9565b80601f016020809104026020016040519081016040528092919081815260200182805461276f906148a9565b80156127bc5780601f10612791576101008083540402835291602001916127bc565b820191906000526020600020905b81548152906001019060200180831161279f57829003601f168201915b5050505050908060050160009054906101000a900467ffffffffffffffff16905086565b61280a7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892633611d10565b612849576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161284090614aa6565b60405180910390fd5b60055481101561288e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161288590615647565b60405180910390fd5b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015414612913576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161290a906156b3565b60405180910390fd5b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015414612998576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161298f9061571f565b60405180910390fd5b6129a133613608565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401612a009392919061573f565b6020604051808303816000875af1158015612a1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a43919061499b565b5080600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154612a929190614846565b600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160009054906101000a900460ff16612c37576001600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160006101000a81548160ff021916908315150217905550600b80549050600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060070181905550600b339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60075442612c459190614846565b600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d82604051612cce9190613f5b565b60405180910390a250565b612ce1613381565b612d0b7f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e826130ab565b50565b60008111612d51576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d48906157c2565b60405180910390fd5b80600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541015612dd6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dcd9061582e565b60405180910390fd5b6000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015414612e5b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e529061589a565b60405180910390fd5b612e6433613608565b60045442612e729190614846565b600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003018190555080600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154612f069190614770565b600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015403612ff2576000600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160006101000a81548160ff0219169083151502179055505b3373ffffffffffffffffffffffffffffffffffffffff167f9cee5eacd317086e5050733a5dad16ef5aaec185de792371a9625bfee2a3213b82600a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015460405161307c9291906158ba565b60405180910390a250565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892681565b6130b482610bb4565b6130bd81613408565b6130c78383613515565b50505050565b6130d5613381565b6130ff7ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c826130ab565b50565b60055481565b7f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e81565b613134613381565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036131a65760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161319d91906142d7565b60405180910390fd5b6131af81613737565b50565b600a6020528060005260406000206000915090508060000154908060010154908060020154908060030154908060040180546131ed906148a9565b80601f0160208091040260200160405190810160405280929190818152602001828054613219906148a9565b80156132665780601f1061323b57610100808354040283529160200191613266565b820191906000526020600020905b81548152906001019060200180831161324957829003601f168201915b50505050509080600501805461327b906148a9565b80601f01602080910402602001604051908101604052809291908181526020018280546132a7906148a9565b80156132f45780601f106132c9576101008083540402835291602001916132f4565b820191906000526020600020905b8154815290600101906020018083116132d757829003601f168201915b5050505050908060060160009054906101000a900460ff16908060070154905088565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b61338961350d565b73ffffffffffffffffffffffffffffffffffffffff166133a7611ce7565b73ffffffffffffffffffffffffffffffffffffffff1614613406576133ca61350d565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016133fd91906142d7565b60405180910390fd5b565b6134198161341461350d565b613b66565b50565b60006134288383611d10565b61350257600180600085815260200190815260200160002060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061349f61350d565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050613507565b600090505b92915050565b600033905090565b60006135218383611d10565b156135fd5760006001600085815260200190815260200160002060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061359a61350d565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a460019050613602565b600090505b92915050565b6000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060060160009054906101000a900460ff16801561366d575060008160000154115b80156136b85750600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020544210155b156137335760008160010154426136cf9190614770565b90506000811115613731576000670de0b6b3a76400008360000154600354846136f891906147a4565b61370291906147a4565b61370c9190614815565b905080836002015461371e9190614846565b8360020181905550428360010181905550505b505b5050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60606000600b80549050905060006102008211613818578161381c565b6102005b905060008111613861576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613858906150a6565b60405180910390fd5b60008167ffffffffffffffff81111561387d5761387c614ac6565b5b6040519080825280602002602001820160405280156138ab5781602001602082028036833780820191505090505b50905060005b8281101561395857600b81815481106138cd576138cc615417565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682828151811061390b5761390a615417565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808061395090615446565b9150506138b1565b5060005b6001836139699190614770565b811015613aed5760005b600182856139819190614770565b61398b9190614770565b811015613ad9576139dc8382815181106139a8576139a7615417565b5b6020026020010151846001846139be9190614846565b815181106139cf576139ce615417565b5b6020026020010151613bb7565b15613ac657826001826139ef9190614846565b81518110613a00576139ff615417565b5b6020026020010151838281518110613a1b57613a1a615417565b5b6020026020010151848381518110613a3657613a35615417565b5b6020026020010185600185613a4b9190614846565b81518110613a5c57613a5b615417565b5b602002602001018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152508273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050505b8080613ad190615446565b915050613973565b508080613ae590615446565b91505061395c565b5080935050505090565b600080600883613b079190614815565b90506000600884613b1891906150c6565b905084518210613b2d57600092505050613b60565b6000816001901b868481518110613b4757613b46615417565b5b602001015160f81c60f81b60f81c60ff16161415925050505b92915050565b613b708282611d10565b613bb35780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401613baa9291906158e3565b60405180910390fd5b5050565b60008173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1611905092915050565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b613c2f81613bfa565b8114613c3a57600080fd5b50565b600081359050613c4c81613c26565b92915050565b600060208284031215613c6857613c67613bf0565b5b6000613c7684828501613c3d565b91505092915050565b60008115159050919050565b613c9481613c7f565b82525050565b6000602082019050613caf6000830184613c8b565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613ce082613cb5565b9050919050565b613cf081613cd5565b8114613cfb57600080fd5b50565b600081359050613d0d81613ce7565b92915050565b600060208284031215613d2957613d28613bf0565b5b6000613d3784828501613cfe565b91505092915050565b6000819050919050565b613d5381613d40565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613d93578082015181840152602081019050613d78565b60008484015250505050565b6000601f19601f8301169050919050565b6000613dbb82613d59565b613dc58185613d64565b9350613dd5818560208601613d75565b613dde81613d9f565b840191505092915050565b600061010082019050613dff600083018b613d4a565b613e0c602083018a613d4a565b613e196040830189613d4a565b8181036060830152613e2b8188613db0565b90508181036080830152613e3f8187613db0565b9050613e4e60a0830186613c8b565b613e5b60c0830185613d4a565b613e6860e0830184613d4a565b9998505050505050505050565b613e7e81613d40565b8114613e8957600080fd5b50565b600081359050613e9b81613e75565b92915050565b600060208284031215613eb757613eb6613bf0565b5b6000613ec584828501613e8c565b91505092915050565b6000819050919050565b613ee181613ece565b8114613eec57600080fd5b50565b600081359050613efe81613ed8565b92915050565b600060208284031215613f1a57613f19613bf0565b5b6000613f2884828501613eef565b91505092915050565b613f3a81613ece565b82525050565b6000602082019050613f556000830184613f31565b92915050565b6000602082019050613f706000830184613d4a565b92915050565b60008060408385031215613f8d57613f8c613bf0565b5b6000613f9b85828601613eef565b9250506020613fac85828601613cfe565b9150509250929050565b600080fd5b600080fd5b600080fd5b60008083601f840112613fdb57613fda613fb6565b5b8235905067ffffffffffffffff811115613ff857613ff7613fbb565b5b60208301915083600182028301111561401457614013613fc0565b5b9250929050565b6000806000806040858703121561403557614034613bf0565b5b600085013567ffffffffffffffff81111561405357614052613bf5565b5b61405f87828801613fc5565b9450945050602085013567ffffffffffffffff81111561408257614081613bf5565b5b61408e87828801613fc5565b925092505092959194509250565b600067ffffffffffffffff82169050919050565b6140b98161409c565b81146140c457600080fd5b50565b6000813590506140d6816140b0565b92915050565b60008083601f8401126140f2576140f1613fb6565b5b8235905067ffffffffffffffff81111561410f5761410e613fbb565b5b60208301915083600182028301111561412b5761412a613fc0565b5b9250929050565b600080600080600080600080600060c08a8c03121561415457614153613bf0565b5b60006141628c828d016140c7565b99505060206141738c828d01613eef565b98505060408a013567ffffffffffffffff81111561419457614193613bf5565b5b6141a08c828d016140dc565b975097505060608a013567ffffffffffffffff8111156141c3576141c2613bf5565b5b6141cf8c828d016140dc565b955095505060808a013567ffffffffffffffff8111156141f2576141f1613bf5565b5b6141fe8c828d016140dc565b935093505060a06142118c828d016140c7565b9150509295985092959850929598565b6000819050919050565b600061424661424161423c84613cb5565b614221565b613cb5565b9050919050565b60006142588261422b565b9050919050565b600061426a8261424d565b9050919050565b61427a8161425f565b82525050565b60006020820190506142956000830184614271565b92915050565b6000602082840312156142b1576142b0613bf0565b5b60006142bf848285016140c7565b91505092915050565b6142d181613cd5565b82525050565b60006020820190506142ec60008301846142c8565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61432781613cd5565b82525050565b6000614339838361431e565b60208301905092915050565b6000602082019050919050565b600061435d826142f2565b61436781856142fd565b93506143728361430e565b8060005b838110156143a357815161438a888261432d565b975061439583614345565b925050600181019050614376565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6143e581613d40565b82525050565b60006143f783836143dc565b60208301905092915050565b6000602082019050919050565b600061441b826143b0565b61442581856143bb565b9350614430836143cc565b8060005b8381101561446157815161444888826143eb565b975061445383614403565b925050600181019050614434565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b60006144b682613d59565b6144c0818561449a565b93506144d0818560208601613d75565b6144d981613d9f565b840191505092915050565b60006144f083836144ab565b905092915050565b6000602082019050919050565b60006145108261446e565b61451a8185614479565b93508360208202850161452c8561448a565b8060005b85811015614568578484038952815161454985826144e4565b9450614554836144f8565b925060208a01995050600181019050614530565b50829750879550505050505092915050565b600060808201905081810360008301526145948187614352565b905081810360208301526145a88186614410565b905081810360408301526145bc8185614505565b905081810360608301526145d08184614505565b905095945050505050565b6145e48161409c565b82525050565b600081519050919050565b600082825260208201905092915050565b6000614611826145ea565b61461b81856145f5565b935061462b818560208601613d75565b61463481613d9f565b840191505092915050565b600060c08201905061465460008301896145db565b6146616020830188613f31565b81810360408301526146738187614606565b905081810360608301526146878186614606565b9050818103608083015261469b8185614606565b90506146aa60a08301846145db565b979650505050505050565b6000610100820190506146cb600083018b613d4a565b6146d8602083018a613d4a565b6146e56040830189613d4a565b6146f26060830188613d4a565b81810360808301526147048187613db0565b905081810360a08301526147188186613db0565b905061472760c0830185613c8b565b61473460e0830184613d4a565b9998505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061477b82613d40565b915061478683613d40565b925082820390508181111561479e5761479d614741565b5b92915050565b60006147af82613d40565b91506147ba83613d40565b92508282026147c881613d40565b915082820484148315176147df576147de614741565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061482082613d40565b915061482b83613d40565b92508261483b5761483a6147e6565b5b828204905092915050565b600061485182613d40565b915061485c83613d40565b925082820190508082111561487457614873614741565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806148c157607f821691505b6020821081036148d4576148d361487a565b5b50919050565b7f4e6f207265776172647320746f20636c61696d00000000000000000000000000600082015250565b6000614910601383613d64565b915061491b826148da565b602082019050919050565b6000602082019050818103600083015261493f81614903565b9050919050565b600060408201905061495b60008301856142c8565b6149686020830184613d4a565b9392505050565b61497881613c7f565b811461498357600080fd5b50565b6000815190506149958161496f565b92915050565b6000602082840312156149b1576149b0613bf0565b5b60006149bf84828501614986565b91505092915050565b7f416c726561647920726567697374657265640000000000000000000000000000600082015250565b60006149fe601283613d64565b9150614a09826149c8565b602082019050919050565b60006020820190508181036000830152614a2d816149f1565b9050919050565b7f434b505456616c5374616b696e673a206d75737420686176652076616c69646160008201527f746f7220726f6c6520746f207374616b65000000000000000000000000000000602082015250565b6000614a90603183613d64565b9150614a9b82614a34565b604082019050919050565b60006020820190508181036000830152614abf81614a83565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302614b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614b1a565b614b618683614b1a565b95508019841693508086168417925050509392505050565b6000614b94614b8f614b8a84613d40565b614221565b613d40565b9050919050565b6000819050919050565b614bae83614b79565b614bc2614bba82614b9b565b848454614b27565b825550505050565b600090565b614bd7614bca565b614be2818484614ba5565b505050565b5b81811015614c0657614bfb600082614bcf565b600181019050614be8565b5050565b601f821115614c4b57614c1c81614af5565b614c2584614b0a565b81016020851015614c34578190505b614c48614c4085614b0a565b830182614be7565b50505b505050565b600082821c905092915050565b6000614c6e60001984600802614c50565b1980831691505092915050565b6000614c878383614c5d565b9150826002028217905092915050565b614ca082613d59565b67ffffffffffffffff811115614cb957614cb8614ac6565b5b614cc382546148a9565b614cce828285614c0a565b600060209050601f831160018114614d015760008415614cef578287015190505b614cf98582614c7b565b865550614d61565b601f198416614d0f86614af5565b60005b82811015614d3757848901518255600182019150602085019450602081019050614d12565b86831015614d545784890151614d50601f891682614c5d565b8355505b6001600288020188555050505b505050505050565b82818337600083830152505050565b6000614d848385613d64565b9350614d91838584614d69565b614d9a83613d9f565b840190509392505050565b60006040820190508181036000830152614dc0818688614d78565b90508181036020830152614dd5818486614d78565b905095945050505050565b7f434b505456616c5374616b696e673a206d75737420686176652064697370617460008201527f6368657220726f6c6520746f207375626d697420636865636b706f696e740000602082015250565b6000614e3c603e83613d64565b9150614e4782614de0565b604082019050919050565b60006020820190508181036000830152614e6b81614e2f565b9050919050565b7f436865636b706f696e7420616c72656164792065786973747320666f7220746860008201527f69732065706f6368000000000000000000000000000000000000000000000000602082015250565b6000614ece602883613d64565b9150614ed982614e72565b604082019050919050565b60006020820190508181036000830152614efd81614ec1565b9050919050565b60008190508160005260206000209050919050565b601f821115614f5a57614f2b81614f04565b614f3484614b0a565b81016020851015614f43578190505b614f57614f4f85614b0a565b830182614be7565b50505b505050565b614f68826145ea565b67ffffffffffffffff811115614f8157614f80614ac6565b5b614f8b82546148a9565b614f96828285614f19565b600060209050601f831160018114614fc95760008415614fb7578287015190505b614fc18582614c7b565b865550615029565b601f198416614fd786614f04565b60005b82811015614fff57848901518255600182019150602085019450602081019050614fda565b8683101561501c5784890151615018601f891682614c5d565b8355505b6001600288020188555050505b505050505050565b60006040820190506150466000830185613f31565b61505360208301846145db565b9392505050565b7f4e6f2076616c696461746f727320617661696c61626c65000000000000000000600082015250565b6000615090601783613d64565b915061509b8261505a565b602082019050919050565b600060208201905081810360008301526150bf81615083565b9050919050565b60006150d182613d40565b91506150dc83613d40565b9250826150ec576150eb6147e6565b5b828206905092915050565b7f4e6f20756e7374616b696e6720696e2070726f67726573730000000000000000600082015250565b600061512d601883613d64565b9150615138826150f7565b602082019050919050565b6000602082019050818103600083015261515c81615120565b9050919050565b7f5374696c6c20696e206c6f636b20706572696f64000000000000000000000000600082015250565b6000615199601483613d64565b91506151a482615163565b602082019050919050565b600060208201905081810360008301526151c88161518c565b9050919050565b7f434b505456616c5374616b696e673a206d75737420686176652064697374726960008201527f627574657220726f6c6520746f20646973747269627574652072657761726473602082015250565b600061522b604083613d64565b9150615236826151cf565b604082019050919050565b6000602082019050818103600083015261525a8161521e565b9050919050565b7f436865636b706f696e7420646f6573206e6f7420657869737420666f7220746860008201527f69732065706f6368000000000000000000000000000000000000000000000000602082015250565b60006152bd602883613d64565b91506152c882615261565b604082019050919050565b600060208201905081810360008301526152ec816152b0565b9050919050565b7f5265776172647320616c726561647920646973747269627574656420666f722060008201527f746869732065706f636800000000000000000000000000000000000000000000602082015250565b600061534f602a83613d64565b915061535a826152f3565b604082019050919050565b6000602082019050818103600083015261537e81615342565b9050919050565b7f4e6f2076616c696461746f727320617661696c61626c6520666f72207265776160008201527f7264730000000000000000000000000000000000000000000000000000000000602082015250565b60006153e1602383613d64565b91506153ec82615385565b604082019050919050565b60006020820190508181036000830152615410816153d4565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061545182613d40565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361548357615482614741565b5b600182019050919050565b7f4e6f7420612076616c696461746f720000000000000000000000000000000000600082015250565b60006154c4600f83613d64565b91506154cf8261548e565b602082019050919050565b600060208201905081810360008301526154f3816154b7565b9050919050565b600082905092915050565b61550f83836154fa565b67ffffffffffffffff81111561552857615527614ac6565b5b61553282546148a9565b61553d828285614c0a565b6000601f83116001811461556c576000841561555a578287013590505b6155648582614c7b565b8655506155cc565b601f19841661557a86614af5565b60005b828110156155a25784890135825560018201915060208501945060208101905061557d565b868310156155bf57848901356155bb601f891682614c5d565b8355505b6001600288020188555050505b50505050505050565b7f4d757374207374616b6520657861637420626967676572207468616e206d696e60008201527f696d756d20616d6f756e74000000000000000000000000000000000000000000602082015250565b6000615631602b83613d64565b915061563c826155d5565b604082019050919050565b6000602082019050818103600083015261566081615624565b9050919050565b7f416c7265616479207374616b6564000000000000000000000000000000000000600082015250565b600061569d600e83613d64565b91506156a882615667565b602082019050919050565b600060208201905081810360008301526156cc81615690565b9050919050565b7f556e7374616b696e6720696e2070726f67726573730000000000000000000000600082015250565b6000615709601583613d64565b9150615714826156d3565b602082019050919050565b60006020820190508181036000830152615738816156fc565b9050919050565b600060608201905061575460008301866142c8565b61576160208301856142c8565b61576e6040830184613d4a565b949350505050565b7f43616e6e6f7420756e7374616b65203000000000000000000000000000000000600082015250565b60006157ac601083613d64565b91506157b782615776565b602082019050919050565b600060208201905081810360008301526157db8161579f565b9050919050565b7f4e6f7420656e6f756768207374616b6564000000000000000000000000000000600082015250565b6000615818601183613d64565b9150615823826157e2565b602082019050919050565b600060208201905081810360008301526158478161580b565b9050919050565b7f556e7374616b696e6720616c726561647920696e2070726f6772657373000000600082015250565b6000615884601d83613d64565b915061588f8261584e565b602082019050919050565b600060208201905081810360008301526158b381615877565b9050919050565b60006040820190506158cf6000830185613d4a565b6158dc6020830184613d4a565b9392505050565b60006040820190506158f860008301856142c8565b6159056020830184613f31565b939250505056fea264697066735822122003c5fd732e5bddd78de8d1dc41c513bd067e785920684674cf556c96e7e2d03264736f6c63430008140033",
}

// CKPTValStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use CKPTValStakingMetaData.ABI instead.
var CKPTValStakingABI = CKPTValStakingMetaData.ABI

// CKPTValStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CKPTValStakingMetaData.Bin instead.
var CKPTValStakingBin = CKPTValStakingMetaData.Bin

// DeployCKPTValStaking deploys a new Ethereum contract, binding an instance of CKPTValStaking to it.
func DeployCKPTValStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _stakingToken common.Address, _rewardRate *big.Int, _minimumStake *big.Int) (common.Address, *types.Transaction, *CKPTValStaking, error) {
	parsed, err := CKPTValStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CKPTValStakingBin), backend, _stakingToken, _rewardRate, _minimumStake)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CKPTValStaking{CKPTValStakingCaller: CKPTValStakingCaller{contract: contract}, CKPTValStakingTransactor: CKPTValStakingTransactor{contract: contract}, CKPTValStakingFilterer: CKPTValStakingFilterer{contract: contract}}, nil
}

// CKPTValStaking is an auto generated Go binding around an Ethereum contract.
type CKPTValStaking struct {
	CKPTValStakingCaller     // Read-only binding to the contract
	CKPTValStakingTransactor // Write-only binding to the contract
	CKPTValStakingFilterer   // Log filterer for contract events
}

// CKPTValStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CKPTValStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKPTValStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CKPTValStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKPTValStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CKPTValStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKPTValStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CKPTValStakingSession struct {
	Contract     *CKPTValStaking   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CKPTValStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CKPTValStakingCallerSession struct {
	Contract *CKPTValStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CKPTValStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CKPTValStakingTransactorSession struct {
	Contract     *CKPTValStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CKPTValStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CKPTValStakingRaw struct {
	Contract *CKPTValStaking // Generic contract binding to access the raw methods on
}

// CKPTValStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CKPTValStakingCallerRaw struct {
	Contract *CKPTValStakingCaller // Generic read-only contract binding to access the raw methods on
}

// CKPTValStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CKPTValStakingTransactorRaw struct {
	Contract *CKPTValStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCKPTValStaking creates a new instance of CKPTValStaking, bound to a specific deployed contract.
func NewCKPTValStaking(address common.Address, backend bind.ContractBackend) (*CKPTValStaking, error) {
	contract, err := bindCKPTValStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CKPTValStaking{CKPTValStakingCaller: CKPTValStakingCaller{contract: contract}, CKPTValStakingTransactor: CKPTValStakingTransactor{contract: contract}, CKPTValStakingFilterer: CKPTValStakingFilterer{contract: contract}}, nil
}

// NewCKPTValStakingCaller creates a new read-only instance of CKPTValStaking, bound to a specific deployed contract.
func NewCKPTValStakingCaller(address common.Address, caller bind.ContractCaller) (*CKPTValStakingCaller, error) {
	contract, err := bindCKPTValStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingCaller{contract: contract}, nil
}

// NewCKPTValStakingTransactor creates a new write-only instance of CKPTValStaking, bound to a specific deployed contract.
func NewCKPTValStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*CKPTValStakingTransactor, error) {
	contract, err := bindCKPTValStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingTransactor{contract: contract}, nil
}

// NewCKPTValStakingFilterer creates a new log filterer instance of CKPTValStaking, bound to a specific deployed contract.
func NewCKPTValStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*CKPTValStakingFilterer, error) {
	contract, err := bindCKPTValStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingFilterer{contract: contract}, nil
}

// bindCKPTValStaking binds a generic wrapper to an already deployed contract.
func bindCKPTValStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CKPTValStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CKPTValStaking *CKPTValStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CKPTValStaking.Contract.CKPTValStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CKPTValStaking *CKPTValStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.CKPTValStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CKPTValStaking *CKPTValStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.CKPTValStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CKPTValStaking *CKPTValStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CKPTValStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CKPTValStaking *CKPTValStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CKPTValStaking *CKPTValStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DEFAULTADMINROLE(&_CKPTValStaking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DEFAULTADMINROLE(&_CKPTValStaking.CallOpts)
}

// DISPATCHERROLE is a free data retrieval call binding the contract method 0x9ef5aee9.
//
// Solidity: function DISPATCHER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) DISPATCHERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "DISPATCHER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISPATCHERROLE is a free data retrieval call binding the contract method 0x9ef5aee9.
//
// Solidity: function DISPATCHER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) DISPATCHERROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DISPATCHERROLE(&_CKPTValStaking.CallOpts)
}

// DISPATCHERROLE is a free data retrieval call binding the contract method 0x9ef5aee9.
//
// Solidity: function DISPATCHER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) DISPATCHERROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DISPATCHERROLE(&_CKPTValStaking.CallOpts)
}

// DISTRIBUTERROLE is a free data retrieval call binding the contract method 0xf10e00d6.
//
// Solidity: function DISTRIBUTER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) DISTRIBUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "DISTRIBUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISTRIBUTERROLE is a free data retrieval call binding the contract method 0xf10e00d6.
//
// Solidity: function DISTRIBUTER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) DISTRIBUTERROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DISTRIBUTERROLE(&_CKPTValStaking.CallOpts)
}

// DISTRIBUTERROLE is a free data retrieval call binding the contract method 0xf10e00d6.
//
// Solidity: function DISTRIBUTER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) DISTRIBUTERROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DISTRIBUTERROLE(&_CKPTValStaking.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) VALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) VALIDATORROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.VALIDATORROLE(&_CKPTValStaking.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) VALIDATORROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.VALIDATORROLE(&_CKPTValStaking.CallOpts)
}

// DistributedEpochs is a free data retrieval call binding the contract method 0xa17c89f9.
//
// Solidity: function distributedEpochs(uint64 ) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCaller) DistributedEpochs(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "distributedEpochs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DistributedEpochs is a free data retrieval call binding the contract method 0xa17c89f9.
//
// Solidity: function distributedEpochs(uint64 ) view returns(bool)
func (_CKPTValStaking *CKPTValStakingSession) DistributedEpochs(arg0 uint64) (bool, error) {
	return _CKPTValStaking.Contract.DistributedEpochs(&_CKPTValStaking.CallOpts, arg0)
}

// DistributedEpochs is a free data retrieval call binding the contract method 0xa17c89f9.
//
// Solidity: function distributedEpochs(uint64 ) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCallerSession) DistributedEpochs(arg0 uint64) (bool, error) {
	return _CKPTValStaking.Contract.DistributedEpochs(&_CKPTValStaking.CallOpts, arg0)
}

// EpochToCheckpoint is a free data retrieval call binding the contract method 0xa5f68a18.
//
// Solidity: function epochToCheckpoint(uint64 ) view returns(uint64 epochNum, bytes32 blockHash, bytes bitmap, bytes blsMultiSig, bytes blsAggrPk, uint64 powerSum)
func (_CKPTValStaking *CKPTValStakingCaller) EpochToCheckpoint(opts *bind.CallOpts, arg0 uint64) (struct {
	EpochNum    uint64
	BlockHash   [32]byte
	Bitmap      []byte
	BlsMultiSig []byte
	BlsAggrPk   []byte
	PowerSum    uint64
}, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "epochToCheckpoint", arg0)

	outstruct := new(struct {
		EpochNum    uint64
		BlockHash   [32]byte
		Bitmap      []byte
		BlsMultiSig []byte
		BlsAggrPk   []byte
		PowerSum    uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EpochNum = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Bitmap = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.BlsMultiSig = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.BlsAggrPk = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.PowerSum = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// EpochToCheckpoint is a free data retrieval call binding the contract method 0xa5f68a18.
//
// Solidity: function epochToCheckpoint(uint64 ) view returns(uint64 epochNum, bytes32 blockHash, bytes bitmap, bytes blsMultiSig, bytes blsAggrPk, uint64 powerSum)
func (_CKPTValStaking *CKPTValStakingSession) EpochToCheckpoint(arg0 uint64) (struct {
	EpochNum    uint64
	BlockHash   [32]byte
	Bitmap      []byte
	BlsMultiSig []byte
	BlsAggrPk   []byte
	PowerSum    uint64
}, error) {
	return _CKPTValStaking.Contract.EpochToCheckpoint(&_CKPTValStaking.CallOpts, arg0)
}

// EpochToCheckpoint is a free data retrieval call binding the contract method 0xa5f68a18.
//
// Solidity: function epochToCheckpoint(uint64 ) view returns(uint64 epochNum, bytes32 blockHash, bytes bitmap, bytes blsMultiSig, bytes blsAggrPk, uint64 powerSum)
func (_CKPTValStaking *CKPTValStakingCallerSession) EpochToCheckpoint(arg0 uint64) (struct {
	EpochNum    uint64
	BlockHash   [32]byte
	Bitmap      []byte
	BlsMultiSig []byte
	BlsAggrPk   []byte
	PowerSum    uint64
}, error) {
	return _CKPTValStaking.Contract.EpochToCheckpoint(&_CKPTValStaking.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CKPTValStaking.Contract.GetRoleAdmin(&_CKPTValStaking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CKPTValStaking.Contract.GetRoleAdmin(&_CKPTValStaking.CallOpts, role)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) GetStake(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getStake", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) GetStake(_validator common.Address) (*big.Int, error) {
	return _CKPTValStaking.Contract.GetStake(&_CKPTValStaking.CallOpts, _validator)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetStake(_validator common.Address) (*big.Int, error) {
	return _CKPTValStaking.Contract.GetStake(&_CKPTValStaking.CallOpts, _validator)
}

// GetTopValidators is a free data retrieval call binding the contract method 0x93a5b1b6.
//
// Solidity: function getTopValidators(uint256 _count) view returns(address[] addresses, uint256[] stakes, string[] dispatcherURLs, string[] blsPublicKeys)
func (_CKPTValStaking *CKPTValStakingCaller) GetTopValidators(opts *bind.CallOpts, _count *big.Int) (struct {
	Addresses      []common.Address
	Stakes         []*big.Int
	DispatcherURLs []string
	BlsPublicKeys  []string
}, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getTopValidators", _count)

	outstruct := new(struct {
		Addresses      []common.Address
		Stakes         []*big.Int
		DispatcherURLs []string
		BlsPublicKeys  []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.DispatcherURLs = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.BlsPublicKeys = *abi.ConvertType(out[3], new([]string)).(*[]string)

	return *outstruct, err

}

// GetTopValidators is a free data retrieval call binding the contract method 0x93a5b1b6.
//
// Solidity: function getTopValidators(uint256 _count) view returns(address[] addresses, uint256[] stakes, string[] dispatcherURLs, string[] blsPublicKeys)
func (_CKPTValStaking *CKPTValStakingSession) GetTopValidators(_count *big.Int) (struct {
	Addresses      []common.Address
	Stakes         []*big.Int
	DispatcherURLs []string
	BlsPublicKeys  []string
}, error) {
	return _CKPTValStaking.Contract.GetTopValidators(&_CKPTValStaking.CallOpts, _count)
}

// GetTopValidators is a free data retrieval call binding the contract method 0x93a5b1b6.
//
// Solidity: function getTopValidators(uint256 _count) view returns(address[] addresses, uint256[] stakes, string[] dispatcherURLs, string[] blsPublicKeys)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetTopValidators(_count *big.Int) (struct {
	Addresses      []common.Address
	Stakes         []*big.Int
	DispatcherURLs []string
	BlsPublicKeys  []string
}, error) {
	return _CKPTValStaking.Contract.GetTopValidators(&_CKPTValStaking.CallOpts, _count)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _validator) view returns(uint256 stakedAmount, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index, uint256 activationTime)
func (_CKPTValStaking *CKPTValStakingCaller) GetValidator(opts *bind.CallOpts, _validator common.Address) (struct {
	StakedAmount   *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
	ActivationTime *big.Int
}, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getValidator", _validator)

	outstruct := new(struct {
		StakedAmount   *big.Int
		PendingRewards *big.Int
		UnstakeTime    *big.Int
		DispatcherURL  string
		BlsPublicKey   string
		IsActive       bool
		Index          *big.Int
		ActivationTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PendingRewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnstakeTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DispatcherURL = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.BlsPublicKey = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Index = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ActivationTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _validator) view returns(uint256 stakedAmount, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index, uint256 activationTime)
func (_CKPTValStaking *CKPTValStakingSession) GetValidator(_validator common.Address) (struct {
	StakedAmount   *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
	ActivationTime *big.Int
}, error) {
	return _CKPTValStaking.Contract.GetValidator(&_CKPTValStaking.CallOpts, _validator)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _validator) view returns(uint256 stakedAmount, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index, uint256 activationTime)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetValidator(_validator common.Address) (struct {
	StakedAmount   *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
	ActivationTime *big.Int
}, error) {
	return _CKPTValStaking.Contract.GetValidator(&_CKPTValStaking.CallOpts, _validator)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) GetValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) GetValidatorCount() (*big.Int, error) {
	return _CKPTValStaking.Contract.GetValidatorCount(&_CKPTValStaking.CallOpts)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetValidatorCount() (*big.Int, error) {
	return _CKPTValStaking.Contract.GetValidatorCount(&_CKPTValStaking.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CKPTValStaking *CKPTValStakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CKPTValStaking.Contract.HasRole(&_CKPTValStaking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CKPTValStaking.Contract.HasRole(&_CKPTValStaking.CallOpts, role, account)
}

// LockPeriod is a free data retrieval call binding the contract method 0x3fd8b02f.
//
// Solidity: function lockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) LockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "lockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockPeriod is a free data retrieval call binding the contract method 0x3fd8b02f.
//
// Solidity: function lockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) LockPeriod() (*big.Int, error) {
	return _CKPTValStaking.Contract.LockPeriod(&_CKPTValStaking.CallOpts)
}

// LockPeriod is a free data retrieval call binding the contract method 0x3fd8b02f.
//
// Solidity: function lockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) LockPeriod() (*big.Int, error) {
	return _CKPTValStaking.Contract.LockPeriod(&_CKPTValStaking.CallOpts)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) MinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "minimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) MinimumStake() (*big.Int, error) {
	return _CKPTValStaking.Contract.MinimumStake(&_CKPTValStaking.CallOpts)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) MinimumStake() (*big.Int, error) {
	return _CKPTValStaking.Contract.MinimumStake(&_CKPTValStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CKPTValStaking *CKPTValStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CKPTValStaking *CKPTValStakingSession) Owner() (common.Address, error) {
	return _CKPTValStaking.Contract.Owner(&_CKPTValStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CKPTValStaking *CKPTValStakingCallerSession) Owner() (common.Address, error) {
	return _CKPTValStaking.Contract.Owner(&_CKPTValStaking.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) RewardRate() (*big.Int, error) {
	return _CKPTValStaking.Contract.RewardRate(&_CKPTValStaking.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) RewardRate() (*big.Int, error) {
	return _CKPTValStaking.Contract.RewardRate(&_CKPTValStaking.CallOpts)
}

// StakeActivationTime is a free data retrieval call binding the contract method 0x2786879f.
//
// Solidity: function stakeActivationTime(address ) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) StakeActivationTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "stakeActivationTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeActivationTime is a free data retrieval call binding the contract method 0x2786879f.
//
// Solidity: function stakeActivationTime(address ) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) StakeActivationTime(arg0 common.Address) (*big.Int, error) {
	return _CKPTValStaking.Contract.StakeActivationTime(&_CKPTValStaking.CallOpts, arg0)
}

// StakeActivationTime is a free data retrieval call binding the contract method 0x2786879f.
//
// Solidity: function stakeActivationTime(address ) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) StakeActivationTime(arg0 common.Address) (*big.Int, error) {
	return _CKPTValStaking.Contract.StakeActivationTime(&_CKPTValStaking.CallOpts, arg0)
}

// StakeLockPeriod is a free data retrieval call binding the contract method 0x499ab37e.
//
// Solidity: function stakeLockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) StakeLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "stakeLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLockPeriod is a free data retrieval call binding the contract method 0x499ab37e.
//
// Solidity: function stakeLockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) StakeLockPeriod() (*big.Int, error) {
	return _CKPTValStaking.Contract.StakeLockPeriod(&_CKPTValStaking.CallOpts)
}

// StakeLockPeriod is a free data retrieval call binding the contract method 0x499ab37e.
//
// Solidity: function stakeLockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) StakeLockPeriod() (*big.Int, error) {
	return _CKPTValStaking.Contract.StakeLockPeriod(&_CKPTValStaking.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CKPTValStaking *CKPTValStakingCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CKPTValStaking *CKPTValStakingSession) StakingToken() (common.Address, error) {
	return _CKPTValStaking.Contract.StakingToken(&_CKPTValStaking.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CKPTValStaking *CKPTValStakingCallerSession) StakingToken() (common.Address, error) {
	return _CKPTValStaking.Contract.StakingToken(&_CKPTValStaking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CKPTValStaking *CKPTValStakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CKPTValStaking.Contract.SupportsInterface(&_CKPTValStaking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CKPTValStaking.Contract.SupportsInterface(&_CKPTValStaking.CallOpts, interfaceId)
}

// ValidatorAddresses is a free data retrieval call binding the contract method 0x99745318.
//
// Solidity: function validatorAddresses(uint256 ) view returns(address)
func (_CKPTValStaking *CKPTValStakingCaller) ValidatorAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "validatorAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorAddresses is a free data retrieval call binding the contract method 0x99745318.
//
// Solidity: function validatorAddresses(uint256 ) view returns(address)
func (_CKPTValStaking *CKPTValStakingSession) ValidatorAddresses(arg0 *big.Int) (common.Address, error) {
	return _CKPTValStaking.Contract.ValidatorAddresses(&_CKPTValStaking.CallOpts, arg0)
}

// ValidatorAddresses is a free data retrieval call binding the contract method 0x99745318.
//
// Solidity: function validatorAddresses(uint256 ) view returns(address)
func (_CKPTValStaking *CKPTValStakingCallerSession) ValidatorAddresses(arg0 *big.Int) (common.Address, error) {
	return _CKPTValStaking.Contract.ValidatorAddresses(&_CKPTValStaking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stakedAmount, uint256 lastRewardTime, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index)
func (_CKPTValStaking *CKPTValStakingCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	StakedAmount   *big.Int
	LastRewardTime *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
}, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		StakedAmount   *big.Int
		LastRewardTime *big.Int
		PendingRewards *big.Int
		UnstakeTime    *big.Int
		DispatcherURL  string
		BlsPublicKey   string
		IsActive       bool
		Index          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PendingRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DispatcherURL = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.BlsPublicKey = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.IsActive = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Index = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stakedAmount, uint256 lastRewardTime, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index)
func (_CKPTValStaking *CKPTValStakingSession) Validators(arg0 common.Address) (struct {
	StakedAmount   *big.Int
	LastRewardTime *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
}, error) {
	return _CKPTValStaking.Contract.Validators(&_CKPTValStaking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stakedAmount, uint256 lastRewardTime, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index)
func (_CKPTValStaking *CKPTValStakingCallerSession) Validators(arg0 common.Address) (struct {
	StakedAmount   *big.Int
	LastRewardTime *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
}, error) {
	return _CKPTValStaking.Contract.Validators(&_CKPTValStaking.CallOpts, arg0)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_CKPTValStaking *CKPTValStakingTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_CKPTValStaking *CKPTValStakingSession) ClaimRewards() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.ClaimRewards(&_CKPTValStaking.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.ClaimRewards(&_CKPTValStaking.TransactOpts)
}

// CompleteUnstake is a paid mutator transaction binding the contract method 0x63803b23.
//
// Solidity: function completeUnstake() returns()
func (_CKPTValStaking *CKPTValStakingTransactor) CompleteUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "completeUnstake")
}

// CompleteUnstake is a paid mutator transaction binding the contract method 0x63803b23.
//
// Solidity: function completeUnstake() returns()
func (_CKPTValStaking *CKPTValStakingSession) CompleteUnstake() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.CompleteUnstake(&_CKPTValStaking.TransactOpts)
}

// CompleteUnstake is a paid mutator transaction binding the contract method 0x63803b23.
//
// Solidity: function completeUnstake() returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) CompleteUnstake() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.CompleteUnstake(&_CKPTValStaking.TransactOpts)
}

// DistributeCheckpointRewards is a paid mutator transaction binding the contract method 0x74a28d61.
//
// Solidity: function distributeCheckpointRewards(uint64 _epochNum) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) DistributeCheckpointRewards(opts *bind.TransactOpts, _epochNum uint64) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "distributeCheckpointRewards", _epochNum)
}

// DistributeCheckpointRewards is a paid mutator transaction binding the contract method 0x74a28d61.
//
// Solidity: function distributeCheckpointRewards(uint64 _epochNum) returns()
func (_CKPTValStaking *CKPTValStakingSession) DistributeCheckpointRewards(_epochNum uint64) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.DistributeCheckpointRewards(&_CKPTValStaking.TransactOpts, _epochNum)
}

// DistributeCheckpointRewards is a paid mutator transaction binding the contract method 0x74a28d61.
//
// Solidity: function distributeCheckpointRewards(uint64 _epochNum) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) DistributeCheckpointRewards(_epochNum uint64) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.DistributeCheckpointRewards(&_CKPTValStaking.TransactOpts, _epochNum)
}

// GrantDispatcherRole is a paid mutator transaction binding the contract method 0x6b11b1eb.
//
// Solidity: function grantDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) GrantDispatcherRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "grantDispatcherRole", account)
}

// GrantDispatcherRole is a paid mutator transaction binding the contract method 0x6b11b1eb.
//
// Solidity: function grantDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) GrantDispatcherRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantDispatcherRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantDispatcherRole is a paid mutator transaction binding the contract method 0x6b11b1eb.
//
// Solidity: function grantDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) GrantDispatcherRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantDispatcherRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantDistributerRole is a paid mutator transaction binding the contract method 0x6d5beac7.
//
// Solidity: function grantDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) GrantDistributerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "grantDistributerRole", account)
}

// GrantDistributerRole is a paid mutator transaction binding the contract method 0x6d5beac7.
//
// Solidity: function grantDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) GrantDistributerRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantDistributerRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantDistributerRole is a paid mutator transaction binding the contract method 0x6d5beac7.
//
// Solidity: function grantDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) GrantDistributerRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantDistributerRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantRole(&_CKPTValStaking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantRole(&_CKPTValStaking.TransactOpts, role, account)
}

// GrantValidatorRole is a paid mutator transaction binding the contract method 0x8ae647df.
//
// Solidity: function grantValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) GrantValidatorRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "grantValidatorRole", account)
}

// GrantValidatorRole is a paid mutator transaction binding the contract method 0x8ae647df.
//
// Solidity: function grantValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) GrantValidatorRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantValidatorRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantValidatorRole is a paid mutator transaction binding the contract method 0x8ae647df.
//
// Solidity: function grantValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) GrantValidatorRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantValidatorRole(&_CKPTValStaking.TransactOpts, account)
}

// InitiateUnstake is a paid mutator transaction binding the contract method 0xae5ac921.
//
// Solidity: function initiateUnstake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) InitiateUnstake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "initiateUnstake", _amount)
}

// InitiateUnstake is a paid mutator transaction binding the contract method 0xae5ac921.
//
// Solidity: function initiateUnstake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingSession) InitiateUnstake(_amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.InitiateUnstake(&_CKPTValStaking.TransactOpts, _amount)
}

// InitiateUnstake is a paid mutator transaction binding the contract method 0xae5ac921.
//
// Solidity: function initiateUnstake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) InitiateUnstake(_amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.InitiateUnstake(&_CKPTValStaking.TransactOpts, _amount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x3d81380d.
//
// Solidity: function registerValidator(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RegisterValidator(opts *bind.TransactOpts, _dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "registerValidator", _dispatcherURL, _blsPublicKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x3d81380d.
//
// Solidity: function registerValidator(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingSession) RegisterValidator(_dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RegisterValidator(&_CKPTValStaking.TransactOpts, _dispatcherURL, _blsPublicKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x3d81380d.
//
// Solidity: function registerValidator(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RegisterValidator(_dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RegisterValidator(&_CKPTValStaking.TransactOpts, _dispatcherURL, _blsPublicKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CKPTValStaking *CKPTValStakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RenounceOwnership(&_CKPTValStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RenounceOwnership(&_CKPTValStaking.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CKPTValStaking *CKPTValStakingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RenounceRole(&_CKPTValStaking.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RenounceRole(&_CKPTValStaking.TransactOpts, role, callerConfirmation)
}

// RevokeDispatcherRole is a paid mutator transaction binding the contract method 0xe18d4f8b.
//
// Solidity: function revokeDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RevokeDispatcherRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "revokeDispatcherRole", account)
}

// RevokeDispatcherRole is a paid mutator transaction binding the contract method 0xe18d4f8b.
//
// Solidity: function revokeDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) RevokeDispatcherRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeDispatcherRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeDispatcherRole is a paid mutator transaction binding the contract method 0xe18d4f8b.
//
// Solidity: function revokeDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RevokeDispatcherRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeDispatcherRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeDistributerRole is a paid mutator transaction binding the contract method 0xa8667ce1.
//
// Solidity: function revokeDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RevokeDistributerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "revokeDistributerRole", account)
}

// RevokeDistributerRole is a paid mutator transaction binding the contract method 0xa8667ce1.
//
// Solidity: function revokeDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) RevokeDistributerRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeDistributerRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeDistributerRole is a paid mutator transaction binding the contract method 0xa8667ce1.
//
// Solidity: function revokeDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RevokeDistributerRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeDistributerRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeRole(&_CKPTValStaking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeRole(&_CKPTValStaking.TransactOpts, role, account)
}

// RevokeValidatorRole is a paid mutator transaction binding the contract method 0x29514204.
//
// Solidity: function revokeValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RevokeValidatorRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "revokeValidatorRole", account)
}

// RevokeValidatorRole is a paid mutator transaction binding the contract method 0x29514204.
//
// Solidity: function revokeValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) RevokeValidatorRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeValidatorRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeValidatorRole is a paid mutator transaction binding the contract method 0x29514204.
//
// Solidity: function revokeValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RevokeValidatorRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeValidatorRole(&_CKPTValStaking.TransactOpts, account)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0x779972da.
//
// Solidity: function setLockPeriod(uint256 _lockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SetLockPeriod(opts *bind.TransactOpts, _lockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "setLockPeriod", _lockPeriod)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0x779972da.
//
// Solidity: function setLockPeriod(uint256 _lockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingSession) SetLockPeriod(_lockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetLockPeriod(&_CKPTValStaking.TransactOpts, _lockPeriod)
}

// SetLockPeriod is a paid mutator transaction binding the contract method 0x779972da.
//
// Solidity: function setLockPeriod(uint256 _lockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SetLockPeriod(_lockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetLockPeriod(&_CKPTValStaking.TransactOpts, _lockPeriod)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 _minimumStake) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SetMinimumStake(opts *bind.TransactOpts, _minimumStake *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "setMinimumStake", _minimumStake)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 _minimumStake) returns()
func (_CKPTValStaking *CKPTValStakingSession) SetMinimumStake(_minimumStake *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetMinimumStake(&_CKPTValStaking.TransactOpts, _minimumStake)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 _minimumStake) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SetMinimumStake(_minimumStake *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetMinimumStake(&_CKPTValStaking.TransactOpts, _minimumStake)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rewardRate) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SetRewardRate(opts *bind.TransactOpts, _rewardRate *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "setRewardRate", _rewardRate)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rewardRate) returns()
func (_CKPTValStaking *CKPTValStakingSession) SetRewardRate(_rewardRate *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetRewardRate(&_CKPTValStaking.TransactOpts, _rewardRate)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rewardRate) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SetRewardRate(_rewardRate *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetRewardRate(&_CKPTValStaking.TransactOpts, _rewardRate)
}

// SetStakeLockPeriod is a paid mutator transaction binding the contract method 0x2e3a9690.
//
// Solidity: function setStakeLockPeriod(uint256 _stakeLockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SetStakeLockPeriod(opts *bind.TransactOpts, _stakeLockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "setStakeLockPeriod", _stakeLockPeriod)
}

// SetStakeLockPeriod is a paid mutator transaction binding the contract method 0x2e3a9690.
//
// Solidity: function setStakeLockPeriod(uint256 _stakeLockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingSession) SetStakeLockPeriod(_stakeLockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetStakeLockPeriod(&_CKPTValStaking.TransactOpts, _stakeLockPeriod)
}

// SetStakeLockPeriod is a paid mutator transaction binding the contract method 0x2e3a9690.
//
// Solidity: function setStakeLockPeriod(uint256 _stakeLockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SetStakeLockPeriod(_stakeLockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetStakeLockPeriod(&_CKPTValStaking.TransactOpts, _stakeLockPeriod)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.Stake(&_CKPTValStaking.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.Stake(&_CKPTValStaking.TransactOpts, _amount)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x4dda27b4.
//
// Solidity: function submitCheckpoint(uint64 _epochNum, bytes32 _blockHash, bytes _bitmap, bytes _blsMultiSig, bytes _blsAggrPk, uint64 _powerSum) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SubmitCheckpoint(opts *bind.TransactOpts, _epochNum uint64, _blockHash [32]byte, _bitmap []byte, _blsMultiSig []byte, _blsAggrPk []byte, _powerSum uint64) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "submitCheckpoint", _epochNum, _blockHash, _bitmap, _blsMultiSig, _blsAggrPk, _powerSum)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x4dda27b4.
//
// Solidity: function submitCheckpoint(uint64 _epochNum, bytes32 _blockHash, bytes _bitmap, bytes _blsMultiSig, bytes _blsAggrPk, uint64 _powerSum) returns()
func (_CKPTValStaking *CKPTValStakingSession) SubmitCheckpoint(_epochNum uint64, _blockHash [32]byte, _bitmap []byte, _blsMultiSig []byte, _blsAggrPk []byte, _powerSum uint64) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SubmitCheckpoint(&_CKPTValStaking.TransactOpts, _epochNum, _blockHash, _bitmap, _blsMultiSig, _blsAggrPk, _powerSum)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x4dda27b4.
//
// Solidity: function submitCheckpoint(uint64 _epochNum, bytes32 _blockHash, bytes _bitmap, bytes _blsMultiSig, bytes _blsAggrPk, uint64 _powerSum) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SubmitCheckpoint(_epochNum uint64, _blockHash [32]byte, _bitmap []byte, _blsMultiSig []byte, _blsAggrPk []byte, _powerSum uint64) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SubmitCheckpoint(&_CKPTValStaking.TransactOpts, _epochNum, _blockHash, _bitmap, _blsMultiSig, _blsAggrPk, _powerSum)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CKPTValStaking *CKPTValStakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.TransferOwnership(&_CKPTValStaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.TransferOwnership(&_CKPTValStaking.TransactOpts, newOwner)
}

// UpdateValidatorCursor is a paid mutator transaction binding the contract method 0x51057fad.
//
// Solidity: function updateValidatorCursor(uint256 _count) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) UpdateValidatorCursor(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "updateValidatorCursor", _count)
}

// UpdateValidatorCursor is a paid mutator transaction binding the contract method 0x51057fad.
//
// Solidity: function updateValidatorCursor(uint256 _count) returns()
func (_CKPTValStaking *CKPTValStakingSession) UpdateValidatorCursor(_count *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.UpdateValidatorCursor(&_CKPTValStaking.TransactOpts, _count)
}

// UpdateValidatorCursor is a paid mutator transaction binding the contract method 0x51057fad.
//
// Solidity: function updateValidatorCursor(uint256 _count) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) UpdateValidatorCursor(_count *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.UpdateValidatorCursor(&_CKPTValStaking.TransactOpts, _count)
}

// UpdateValidatorInfo is a paid mutator transaction binding the contract method 0x9dc820a5.
//
// Solidity: function updateValidatorInfo(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) UpdateValidatorInfo(opts *bind.TransactOpts, _dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "updateValidatorInfo", _dispatcherURL, _blsPublicKey)
}

// UpdateValidatorInfo is a paid mutator transaction binding the contract method 0x9dc820a5.
//
// Solidity: function updateValidatorInfo(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingSession) UpdateValidatorInfo(_dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.UpdateValidatorInfo(&_CKPTValStaking.TransactOpts, _dispatcherURL, _blsPublicKey)
}

// UpdateValidatorInfo is a paid mutator transaction binding the contract method 0x9dc820a5.
//
// Solidity: function updateValidatorInfo(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) UpdateValidatorInfo(_dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.UpdateValidatorInfo(&_CKPTValStaking.TransactOpts, _dispatcherURL, _blsPublicKey)
}

// CKPTValStakingCheckpointSubmittedIterator is returned from FilterCheckpointSubmitted and is used to iterate over the raw logs and unpacked data for CheckpointSubmitted events raised by the CKPTValStaking contract.
type CKPTValStakingCheckpointSubmittedIterator struct {
	Event *CKPTValStakingCheckpointSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingCheckpointSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingCheckpointSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingCheckpointSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingCheckpointSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingCheckpointSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingCheckpointSubmitted represents a CheckpointSubmitted event raised by the CKPTValStaking contract.
type CKPTValStakingCheckpointSubmitted struct {
	EpochNum  uint64
	BlockHash [32]byte
	PowerSum  uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCheckpointSubmitted is a free log retrieval operation binding the contract event 0x609e08589f9d05d683838b8bb92f105d0796c16700397a452a95b4c936dac346.
//
// Solidity: event CheckpointSubmitted(uint64 indexed epochNum, bytes32 blockHash, uint64 powerSum)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterCheckpointSubmitted(opts *bind.FilterOpts, epochNum []uint64) (*CKPTValStakingCheckpointSubmittedIterator, error) {

	var epochNumRule []interface{}
	for _, epochNumItem := range epochNum {
		epochNumRule = append(epochNumRule, epochNumItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "CheckpointSubmitted", epochNumRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingCheckpointSubmittedIterator{contract: _CKPTValStaking.contract, event: "CheckpointSubmitted", logs: logs, sub: sub}, nil
}

// WatchCheckpointSubmitted is a free log subscription operation binding the contract event 0x609e08589f9d05d683838b8bb92f105d0796c16700397a452a95b4c936dac346.
//
// Solidity: event CheckpointSubmitted(uint64 indexed epochNum, bytes32 blockHash, uint64 powerSum)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchCheckpointSubmitted(opts *bind.WatchOpts, sink chan<- *CKPTValStakingCheckpointSubmitted, epochNum []uint64) (event.Subscription, error) {

	var epochNumRule []interface{}
	for _, epochNumItem := range epochNum {
		epochNumRule = append(epochNumRule, epochNumItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "CheckpointSubmitted", epochNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingCheckpointSubmitted)
				if err := _CKPTValStaking.contract.UnpackLog(event, "CheckpointSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointSubmitted is a log parse operation binding the contract event 0x609e08589f9d05d683838b8bb92f105d0796c16700397a452a95b4c936dac346.
//
// Solidity: event CheckpointSubmitted(uint64 indexed epochNum, bytes32 blockHash, uint64 powerSum)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseCheckpointSubmitted(log types.Log) (*CKPTValStakingCheckpointSubmitted, error) {
	event := new(CKPTValStakingCheckpointSubmitted)
	if err := _CKPTValStaking.contract.UnpackLog(event, "CheckpointSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CKPTValStaking contract.
type CKPTValStakingOwnershipTransferredIterator struct {
	Event *CKPTValStakingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingOwnershipTransferred represents a OwnershipTransferred event raised by the CKPTValStaking contract.
type CKPTValStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CKPTValStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingOwnershipTransferredIterator{contract: _CKPTValStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CKPTValStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingOwnershipTransferred)
				if err := _CKPTValStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseOwnershipTransferred(log types.Log) (*CKPTValStakingOwnershipTransferred, error) {
	event := new(CKPTValStakingOwnershipTransferred)
	if err := _CKPTValStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the CKPTValStaking contract.
type CKPTValStakingRewardsClaimedIterator struct {
	Event *CKPTValStakingRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingRewardsClaimed represents a RewardsClaimed event raised by the CKPTValStaking contract.
type CKPTValStakingRewardsClaimed struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingRewardsClaimedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "RewardsClaimed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingRewardsClaimedIterator{contract: _CKPTValStaking.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *CKPTValStakingRewardsClaimed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "RewardsClaimed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingRewardsClaimed)
				if err := _CKPTValStaking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsClaimed is a log parse operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseRewardsClaimed(log types.Log) (*CKPTValStakingRewardsClaimed, error) {
	event := new(CKPTValStakingRewardsClaimed)
	if err := _CKPTValStaking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CKPTValStaking contract.
type CKPTValStakingRoleAdminChangedIterator struct {
	Event *CKPTValStakingRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingRoleAdminChanged represents a RoleAdminChanged event raised by the CKPTValStaking contract.
type CKPTValStakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CKPTValStakingRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingRoleAdminChangedIterator{contract: _CKPTValStaking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CKPTValStakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingRoleAdminChanged)
				if err := _CKPTValStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseRoleAdminChanged(log types.Log) (*CKPTValStakingRoleAdminChanged, error) {
	event := new(CKPTValStakingRoleAdminChanged)
	if err := _CKPTValStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CKPTValStaking contract.
type CKPTValStakingRoleGrantedIterator struct {
	Event *CKPTValStakingRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingRoleGranted represents a RoleGranted event raised by the CKPTValStaking contract.
type CKPTValStakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CKPTValStakingRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingRoleGrantedIterator{contract: _CKPTValStaking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CKPTValStakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingRoleGranted)
				if err := _CKPTValStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseRoleGranted(log types.Log) (*CKPTValStakingRoleGranted, error) {
	event := new(CKPTValStakingRoleGranted)
	if err := _CKPTValStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CKPTValStaking contract.
type CKPTValStakingRoleRevokedIterator struct {
	Event *CKPTValStakingRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingRoleRevoked represents a RoleRevoked event raised by the CKPTValStaking contract.
type CKPTValStakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CKPTValStakingRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingRoleRevokedIterator{contract: _CKPTValStaking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CKPTValStakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingRoleRevoked)
				if err := _CKPTValStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseRoleRevoked(log types.Log) (*CKPTValStakingRoleRevoked, error) {
	event := new(CKPTValStakingRoleRevoked)
	if err := _CKPTValStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the CKPTValStaking contract.
type CKPTValStakingStakedIterator struct {
	Event *CKPTValStakingStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingStaked represents a Staked event raised by the CKPTValStaking contract.
type CKPTValStakingStaked struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterStaked(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingStakedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "Staked", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingStakedIterator{contract: _CKPTValStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *CKPTValStakingStaked, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "Staked", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingStaked)
				if err := _CKPTValStaking.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseStaked(log types.Log) (*CKPTValStakingStaked, error) {
	event := new(CKPTValStakingStaked)
	if err := _CKPTValStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingUnstakeInitiatedIterator is returned from FilterUnstakeInitiated and is used to iterate over the raw logs and unpacked data for UnstakeInitiated events raised by the CKPTValStaking contract.
type CKPTValStakingUnstakeInitiatedIterator struct {
	Event *CKPTValStakingUnstakeInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingUnstakeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingUnstakeInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingUnstakeInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingUnstakeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingUnstakeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingUnstakeInitiated represents a UnstakeInitiated event raised by the CKPTValStaking contract.
type CKPTValStakingUnstakeInitiated struct {
	Validator  common.Address
	Amount     *big.Int
	UnlockTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnstakeInitiated is a free log retrieval operation binding the contract event 0x9cee5eacd317086e5050733a5dad16ef5aaec185de792371a9625bfee2a3213b.
//
// Solidity: event UnstakeInitiated(address indexed validator, uint256 amount, uint256 unlockTime)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterUnstakeInitiated(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingUnstakeInitiatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "UnstakeInitiated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingUnstakeInitiatedIterator{contract: _CKPTValStaking.contract, event: "UnstakeInitiated", logs: logs, sub: sub}, nil
}

// WatchUnstakeInitiated is a free log subscription operation binding the contract event 0x9cee5eacd317086e5050733a5dad16ef5aaec185de792371a9625bfee2a3213b.
//
// Solidity: event UnstakeInitiated(address indexed validator, uint256 amount, uint256 unlockTime)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchUnstakeInitiated(opts *bind.WatchOpts, sink chan<- *CKPTValStakingUnstakeInitiated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "UnstakeInitiated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingUnstakeInitiated)
				if err := _CKPTValStaking.contract.UnpackLog(event, "UnstakeInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstakeInitiated is a log parse operation binding the contract event 0x9cee5eacd317086e5050733a5dad16ef5aaec185de792371a9625bfee2a3213b.
//
// Solidity: event UnstakeInitiated(address indexed validator, uint256 amount, uint256 unlockTime)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseUnstakeInitiated(log types.Log) (*CKPTValStakingUnstakeInitiated, error) {
	event := new(CKPTValStakingUnstakeInitiated)
	if err := _CKPTValStaking.contract.UnpackLog(event, "UnstakeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the CKPTValStaking contract.
type CKPTValStakingUnstakedIterator struct {
	Event *CKPTValStakingUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingUnstaked represents a Unstaked event raised by the CKPTValStaking contract.
type CKPTValStakingUnstaked struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterUnstaked(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingUnstakedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "Unstaked", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingUnstakedIterator{contract: _CKPTValStaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *CKPTValStakingUnstaked, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "Unstaked", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingUnstaked)
				if err := _CKPTValStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseUnstaked(log types.Log) (*CKPTValStakingUnstaked, error) {
	event := new(CKPTValStakingUnstaked)
	if err := _CKPTValStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the CKPTValStaking contract.
type CKPTValStakingValidatorRegisteredIterator struct {
	Event *CKPTValStakingValidatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingValidatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingValidatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingValidatorRegistered represents a ValidatorRegistered event raised by the CKPTValStaking contract.
type CKPTValStakingValidatorRegistered struct {
	Validator     common.Address
	DispatcherURL string
	BlsPublicKey  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x5196ca52b5246aa712ef709c67d65576c4f5ca8aba7d2df5d782cca62cae21d5.
//
// Solidity: event ValidatorRegistered(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingValidatorRegisteredIterator{contract: _CKPTValStaking.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x5196ca52b5246aa712ef709c67d65576c4f5ca8aba7d2df5d782cca62cae21d5.
//
// Solidity: event ValidatorRegistered(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *CKPTValStakingValidatorRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingValidatorRegistered)
				if err := _CKPTValStaking.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistered is a log parse operation binding the contract event 0x5196ca52b5246aa712ef709c67d65576c4f5ca8aba7d2df5d782cca62cae21d5.
//
// Solidity: event ValidatorRegistered(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseValidatorRegistered(log types.Log) (*CKPTValStakingValidatorRegistered, error) {
	event := new(CKPTValStakingValidatorRegistered)
	if err := _CKPTValStaking.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingValidatorUpdatedIterator is returned from FilterValidatorUpdated and is used to iterate over the raw logs and unpacked data for ValidatorUpdated events raised by the CKPTValStaking contract.
type CKPTValStakingValidatorUpdatedIterator struct {
	Event *CKPTValStakingValidatorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingValidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingValidatorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingValidatorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingValidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingValidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingValidatorUpdated represents a ValidatorUpdated event raised by the CKPTValStaking contract.
type CKPTValStakingValidatorUpdated struct {
	Validator     common.Address
	DispatcherURL string
	BlsPublicKey  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorUpdated is a free log retrieval operation binding the contract event 0x7f35b3dd337a65b9cbaf229ba2c1c6e4aa76460db271df1248aa1a605d44ddf2.
//
// Solidity: event ValidatorUpdated(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterValidatorUpdated(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingValidatorUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "ValidatorUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingValidatorUpdatedIterator{contract: _CKPTValStaking.contract, event: "ValidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorUpdated is a free log subscription operation binding the contract event 0x7f35b3dd337a65b9cbaf229ba2c1c6e4aa76460db271df1248aa1a605d44ddf2.
//
// Solidity: event ValidatorUpdated(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchValidatorUpdated(opts *bind.WatchOpts, sink chan<- *CKPTValStakingValidatorUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "ValidatorUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingValidatorUpdated)
				if err := _CKPTValStaking.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorUpdated is a log parse operation binding the contract event 0x7f35b3dd337a65b9cbaf229ba2c1c6e4aa76460db271df1248aa1a605d44ddf2.
//
// Solidity: event ValidatorUpdated(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseValidatorUpdated(log types.Log) (*CKPTValStakingValidatorUpdated, error) {
	event := new(CKPTValStakingValidatorUpdated)
	if err := _CKPTValStaking.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
