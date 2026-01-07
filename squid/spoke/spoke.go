// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package spoke

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

// ISignatureTransferPermitTransferFrom is an auto generated low-level Go binding around an user-defined struct.
type ISignatureTransferPermitTransferFrom struct {
	Permitted ISignatureTransferTokenPermissions
	Nonce     *big.Int
	Deadline  *big.Int
}

// ISignatureTransferTokenPermissions is an auto generated low-level Go binding around an user-defined struct.
type ISignatureTransferTokenPermissions struct {
	Token  common.Address
	Amount *big.Int
}

// ISpokeOrder is an auto generated low-level Go binding around an user-defined struct.
type ISpokeOrder struct {
	FromAddress  common.Address
	ToAddress    common.Address
	Filler       common.Address
	FromToken    common.Address
	ToToken      common.Address
	Expiry       *big.Int
	FromAmount   *big.Int
	FillAmount   *big.Int
	FeeRate      *big.Int
	FromChain    *big.Int
	ToChain      *big.Int
	PostHookHash [32]byte
}

// ISquidMulticallCall is an auto generated low-level Go binding around an user-defined struct.
type ISquidMulticallCall struct {
	CallType uint8
	Target   common.Address
	Value    *big.Int
	CallData []byte
	Payload  []byte
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// SpokeMetaData contains all meta data concerning the Spoke contract.
var SpokeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasRequired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFillType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"optionType\",\"type\":\"uint16\"}],\"name\":\"InvalidOptionType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPostHookProvided\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSettlementFiller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSettlementSourceToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTotalNativeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUserSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokensNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedByGateway\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyFeeCollector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyFillerCanSettle\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderAlreadySettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderStateNotCreated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedOriginAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedOriginEndpoint\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedNativeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UntrustedChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroStringLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postHookHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structISpoke.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postHookHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structISpoke.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"SettlementForwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIAxelarGateway\",\"name\":\"gateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIAxelarGasService\",\"name\":\"gasService\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIPermit2\",\"name\":\"permit2\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractISquidMulticall\",\"name\":\"squidMulticall\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hubChainName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hubAddress\",\"type\":\"string\"}],\"name\":\"SpokeInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"TokensReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"}],\"name\":\"TrustedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"address_\",\"type\":\"string\"}],\"name\":\"TrustedAddressSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addressToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postHookHash\",\"type\":\"bytes32\"}],\"internalType\":\"structISpoke.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumISquidMulticall.CallType\",\"name\":\"callType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structISquidMulticall.Call[][]\",\"name\":\"calls\",\"type\":\"tuple[][]\"}],\"name\":\"batchFillOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postHookHash\",\"type\":\"bytes32\"}],\"internalType\":\"structISpoke.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"fromTokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"}],\"name\":\"batchMultiTokenSingleChainSettlements\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"collectFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postHookHash\",\"type\":\"bytes32\"}],\"internalType\":\"structISpoke.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commandId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commandId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeWithToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postHookHash\",\"type\":\"bytes32\"}],\"internalType\":\"structISpoke.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumISquidMulticall.CallType\",\"name\":\"callType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structISquidMulticall.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"fillOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"lzFee\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"gasLimit\",\"type\":\"uint128\"},{\"internalType\":\"enumICoral.Provider\",\"name\":\"provider\",\"type\":\"uint8\"}],\"name\":\"forwardSettlements\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasService\",\"outputs\":[{\"internalType\":\"contractIAxelarGasService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIAxelarGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hubAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hubAddressBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hubChainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hubEndpoint\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAxelarGateway\",\"name\":\"_axelarGateway\",\"type\":\"address\"},{\"internalType\":\"contractIAxelarGasService\",\"name\":\"_axelarGasService\",\"type\":\"address\"},{\"internalType\":\"contractIPermit2\",\"name\":\"_permit2\",\"type\":\"address\"},{\"internalType\":\"contractISquidMulticall\",\"name\":\"_squidMulticall\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_hubChainName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hubAddress\",\"type\":\"string\"},{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"_endpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hub\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_hubEndpoint\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"address_\",\"type\":\"string\"}],\"name\":\"isTrustedAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderHashToStatus\",\"outputs\":[{\"internalType\":\"enumISpoke.OrderStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2\",\"outputs\":[{\"internalType\":\"contractIPermit2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint128\",\"name\":\"gasLimit\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"_payInLzToken\",\"type\":\"bool\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postHookHash\",\"type\":\"bytes32\"}],\"internalType\":\"structISpoke.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"refundOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"settlementToStatus\",\"outputs\":[{\"internalType\":\"enumISpoke.SettlementStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postHookHash\",\"type\":\"bytes32\"}],\"internalType\":\"structISpoke.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"sponsorOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"postHookHash\",\"type\":\"bytes32\"}],\"internalType\":\"structISpoke.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureTransfer.TokenPermissions\",\"name\":\"permitted\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureTransfer.PermitTransferFrom\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"sponsorOrderUsingPermit2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"squidMulticall\",\"outputs\":[{\"internalType\":\"contractISquidMulticall\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToCollectedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"}],\"name\":\"trustedAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"trustedAddress_\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chain\",\"type\":\"string\"}],\"name\":\"trustedAddressHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"trustedAddressHash_\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SpokeABI is the input ABI used to generate the binding from.
// Deprecated: Use SpokeMetaData.ABI instead.
var SpokeABI = SpokeMetaData.ABI

// Spoke is an auto generated Go binding around an Ethereum contract.
type Spoke struct {
	SpokeCaller     // Read-only binding to the contract
	SpokeTransactor // Write-only binding to the contract
	SpokeFilterer   // Log filterer for contract events
}

// SpokeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpokeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpokeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpokeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpokeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpokeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpokeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpokeSession struct {
	Contract     *Spoke            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpokeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpokeCallerSession struct {
	Contract *SpokeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SpokeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpokeTransactorSession struct {
	Contract     *SpokeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpokeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpokeRaw struct {
	Contract *Spoke // Generic contract binding to access the raw methods on
}

// SpokeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpokeCallerRaw struct {
	Contract *SpokeCaller // Generic read-only contract binding to access the raw methods on
}

// SpokeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpokeTransactorRaw struct {
	Contract *SpokeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpoke creates a new instance of Spoke, bound to a specific deployed contract.
func NewSpoke(address common.Address, backend bind.ContractBackend) (*Spoke, error) {
	contract, err := bindSpoke(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spoke{SpokeCaller: SpokeCaller{contract: contract}, SpokeTransactor: SpokeTransactor{contract: contract}, SpokeFilterer: SpokeFilterer{contract: contract}}, nil
}

// NewSpokeCaller creates a new read-only instance of Spoke, bound to a specific deployed contract.
func NewSpokeCaller(address common.Address, caller bind.ContractCaller) (*SpokeCaller, error) {
	contract, err := bindSpoke(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpokeCaller{contract: contract}, nil
}

// NewSpokeTransactor creates a new write-only instance of Spoke, bound to a specific deployed contract.
func NewSpokeTransactor(address common.Address, transactor bind.ContractTransactor) (*SpokeTransactor, error) {
	contract, err := bindSpoke(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpokeTransactor{contract: contract}, nil
}

// NewSpokeFilterer creates a new log filterer instance of Spoke, bound to a specific deployed contract.
func NewSpokeFilterer(address common.Address, filterer bind.ContractFilterer) (*SpokeFilterer, error) {
	contract, err := bindSpoke(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpokeFilterer{contract: contract}, nil
}

// bindSpoke binds a generic wrapper to an already deployed contract.
func bindSpoke(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SpokeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spoke *SpokeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spoke.Contract.SpokeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spoke *SpokeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spoke.Contract.SpokeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spoke *SpokeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spoke.Contract.SpokeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spoke *SpokeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spoke.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spoke *SpokeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spoke.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spoke *SpokeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spoke.Contract.contract.Transact(opts, method, params...)
}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address _addr) pure returns(bytes32)
func (_Spoke *SpokeCaller) AddressToBytes32(opts *bind.CallOpts, _addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "addressToBytes32", _addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address _addr) pure returns(bytes32)
func (_Spoke *SpokeSession) AddressToBytes32(_addr common.Address) ([32]byte, error) {
	return _Spoke.Contract.AddressToBytes32(&_Spoke.CallOpts, _addr)
}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address _addr) pure returns(bytes32)
func (_Spoke *SpokeCallerSession) AddressToBytes32(_addr common.Address) ([32]byte, error) {
	return _Spoke.Contract.AddressToBytes32(&_Spoke.CallOpts, _addr)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Spoke *SpokeCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Spoke *SpokeSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Spoke.Contract.AllowInitializePath(&_Spoke.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Spoke *SpokeCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Spoke.Contract.AllowInitializePath(&_Spoke.CallOpts, origin)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string chainName_)
func (_Spoke *SpokeCaller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string chainName_)
func (_Spoke *SpokeSession) ChainName() (string, error) {
	return _Spoke.Contract.ChainName(&_Spoke.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string chainName_)
func (_Spoke *SpokeCallerSession) ChainName() (string, error) {
	return _Spoke.Contract.ChainName(&_Spoke.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Spoke *SpokeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Spoke *SpokeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Spoke.Contract.Eip712Domain(&_Spoke.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Spoke *SpokeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Spoke.Contract.Eip712Domain(&_Spoke.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Spoke *SpokeCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Spoke *SpokeSession) Endpoint() (common.Address, error) {
	return _Spoke.Contract.Endpoint(&_Spoke.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Spoke *SpokeCallerSession) Endpoint() (common.Address, error) {
	return _Spoke.Contract.Endpoint(&_Spoke.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Spoke *SpokeCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Spoke *SpokeSession) FeeCollector() (common.Address, error) {
	return _Spoke.Contract.FeeCollector(&_Spoke.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Spoke *SpokeCallerSession) FeeCollector() (common.Address, error) {
	return _Spoke.Contract.FeeCollector(&_Spoke.CallOpts)
}

// GasService is a free data retrieval call binding the contract method 0x6a22d8cc.
//
// Solidity: function gasService() view returns(address)
func (_Spoke *SpokeCaller) GasService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "gasService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasService is a free data retrieval call binding the contract method 0x6a22d8cc.
//
// Solidity: function gasService() view returns(address)
func (_Spoke *SpokeSession) GasService() (common.Address, error) {
	return _Spoke.Contract.GasService(&_Spoke.CallOpts)
}

// GasService is a free data retrieval call binding the contract method 0x6a22d8cc.
//
// Solidity: function gasService() view returns(address)
func (_Spoke *SpokeCallerSession) GasService() (common.Address, error) {
	return _Spoke.Contract.GasService(&_Spoke.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Spoke *SpokeCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Spoke *SpokeSession) Gateway() (common.Address, error) {
	return _Spoke.Contract.Gateway(&_Spoke.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Spoke *SpokeCallerSession) Gateway() (common.Address, error) {
	return _Spoke.Contract.Gateway(&_Spoke.CallOpts)
}

// HubAddress is a free data retrieval call binding the contract method 0xf69b1b29.
//
// Solidity: function hubAddress() view returns(string)
func (_Spoke *SpokeCaller) HubAddress(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "hubAddress")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HubAddress is a free data retrieval call binding the contract method 0xf69b1b29.
//
// Solidity: function hubAddress() view returns(string)
func (_Spoke *SpokeSession) HubAddress() (string, error) {
	return _Spoke.Contract.HubAddress(&_Spoke.CallOpts)
}

// HubAddress is a free data retrieval call binding the contract method 0xf69b1b29.
//
// Solidity: function hubAddress() view returns(string)
func (_Spoke *SpokeCallerSession) HubAddress() (string, error) {
	return _Spoke.Contract.HubAddress(&_Spoke.CallOpts)
}

// HubAddressBytes32 is a free data retrieval call binding the contract method 0x951f4016.
//
// Solidity: function hubAddressBytes32() view returns(bytes32)
func (_Spoke *SpokeCaller) HubAddressBytes32(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "hubAddressBytes32")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HubAddressBytes32 is a free data retrieval call binding the contract method 0x951f4016.
//
// Solidity: function hubAddressBytes32() view returns(bytes32)
func (_Spoke *SpokeSession) HubAddressBytes32() ([32]byte, error) {
	return _Spoke.Contract.HubAddressBytes32(&_Spoke.CallOpts)
}

// HubAddressBytes32 is a free data retrieval call binding the contract method 0x951f4016.
//
// Solidity: function hubAddressBytes32() view returns(bytes32)
func (_Spoke *SpokeCallerSession) HubAddressBytes32() ([32]byte, error) {
	return _Spoke.Contract.HubAddressBytes32(&_Spoke.CallOpts)
}

// HubChainName is a free data retrieval call binding the contract method 0x82e5f9e4.
//
// Solidity: function hubChainName() view returns(string)
func (_Spoke *SpokeCaller) HubChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "hubChainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HubChainName is a free data retrieval call binding the contract method 0x82e5f9e4.
//
// Solidity: function hubChainName() view returns(string)
func (_Spoke *SpokeSession) HubChainName() (string, error) {
	return _Spoke.Contract.HubChainName(&_Spoke.CallOpts)
}

// HubChainName is a free data retrieval call binding the contract method 0x82e5f9e4.
//
// Solidity: function hubChainName() view returns(string)
func (_Spoke *SpokeCallerSession) HubChainName() (string, error) {
	return _Spoke.Contract.HubChainName(&_Spoke.CallOpts)
}

// HubEndpoint is a free data retrieval call binding the contract method 0x2b8cb1f6.
//
// Solidity: function hubEndpoint() view returns(uint32)
func (_Spoke *SpokeCaller) HubEndpoint(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "hubEndpoint")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// HubEndpoint is a free data retrieval call binding the contract method 0x2b8cb1f6.
//
// Solidity: function hubEndpoint() view returns(uint32)
func (_Spoke *SpokeSession) HubEndpoint() (uint32, error) {
	return _Spoke.Contract.HubEndpoint(&_Spoke.CallOpts)
}

// HubEndpoint is a free data retrieval call binding the contract method 0x2b8cb1f6.
//
// Solidity: function hubEndpoint() view returns(uint32)
func (_Spoke *SpokeCallerSession) HubEndpoint() (uint32, error) {
	return _Spoke.Contract.HubEndpoint(&_Spoke.CallOpts)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Spoke *SpokeCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Spoke *SpokeSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Spoke.Contract.IsComposeMsgSender(&_Spoke.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Spoke *SpokeCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Spoke.Contract.IsComposeMsgSender(&_Spoke.CallOpts, arg0, arg1, _sender)
}

// IsTrustedAddress is a free data retrieval call binding the contract method 0xc506bff4.
//
// Solidity: function isTrustedAddress(string chain, string address_) view returns(bool)
func (_Spoke *SpokeCaller) IsTrustedAddress(opts *bind.CallOpts, chain string, address_ string) (bool, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "isTrustedAddress", chain, address_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedAddress is a free data retrieval call binding the contract method 0xc506bff4.
//
// Solidity: function isTrustedAddress(string chain, string address_) view returns(bool)
func (_Spoke *SpokeSession) IsTrustedAddress(chain string, address_ string) (bool, error) {
	return _Spoke.Contract.IsTrustedAddress(&_Spoke.CallOpts, chain, address_)
}

// IsTrustedAddress is a free data retrieval call binding the contract method 0xc506bff4.
//
// Solidity: function isTrustedAddress(string chain, string address_) view returns(bool)
func (_Spoke *SpokeCallerSession) IsTrustedAddress(chain string, address_ string) (bool, error) {
	return _Spoke.Contract.IsTrustedAddress(&_Spoke.CallOpts, chain, address_)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Spoke *SpokeCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Spoke *SpokeSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Spoke.Contract.NextNonce(&_Spoke.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Spoke *SpokeCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Spoke.Contract.NextNonce(&_Spoke.CallOpts, arg0, arg1)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Spoke *SpokeCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Spoke *SpokeSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Spoke.Contract.OAppVersion(&_Spoke.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Spoke *SpokeCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Spoke.Contract.OAppVersion(&_Spoke.CallOpts)
}

// OrderHashToStatus is a free data retrieval call binding the contract method 0x076e9f6c.
//
// Solidity: function orderHashToStatus(bytes32 ) view returns(uint8)
func (_Spoke *SpokeCaller) OrderHashToStatus(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "orderHashToStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OrderHashToStatus is a free data retrieval call binding the contract method 0x076e9f6c.
//
// Solidity: function orderHashToStatus(bytes32 ) view returns(uint8)
func (_Spoke *SpokeSession) OrderHashToStatus(arg0 [32]byte) (uint8, error) {
	return _Spoke.Contract.OrderHashToStatus(&_Spoke.CallOpts, arg0)
}

// OrderHashToStatus is a free data retrieval call binding the contract method 0x076e9f6c.
//
// Solidity: function orderHashToStatus(bytes32 ) view returns(uint8)
func (_Spoke *SpokeCallerSession) OrderHashToStatus(arg0 [32]byte) (uint8, error) {
	return _Spoke.Contract.OrderHashToStatus(&_Spoke.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Spoke *SpokeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Spoke *SpokeSession) Owner() (common.Address, error) {
	return _Spoke.Contract.Owner(&_Spoke.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Spoke *SpokeCallerSession) Owner() (common.Address, error) {
	return _Spoke.Contract.Owner(&_Spoke.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Spoke *SpokeCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Spoke *SpokeSession) Peers(eid uint32) ([32]byte, error) {
	return _Spoke.Contract.Peers(&_Spoke.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Spoke *SpokeCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _Spoke.Contract.Peers(&_Spoke.CallOpts, eid)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Spoke *SpokeCaller) Permit2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "permit2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Spoke *SpokeSession) Permit2() (common.Address, error) {
	return _Spoke.Contract.Permit2(&_Spoke.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Spoke *SpokeCallerSession) Permit2() (common.Address, error) {
	return _Spoke.Contract.Permit2(&_Spoke.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0x21b4ae78.
//
// Solidity: function quote(uint32 _dstEid, bytes32[] orderHashes, uint128 gasLimit, bool _payInLzToken) view returns(uint256 nativeFee, uint256 lzTokenFee)
func (_Spoke *SpokeCaller) Quote(opts *bind.CallOpts, _dstEid uint32, orderHashes [][32]byte, gasLimit *big.Int, _payInLzToken bool) (struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "quote", _dstEid, orderHashes, gasLimit, _payInLzToken)

	outstruct := new(struct {
		NativeFee  *big.Int
		LzTokenFee *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LzTokenFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Quote is a free data retrieval call binding the contract method 0x21b4ae78.
//
// Solidity: function quote(uint32 _dstEid, bytes32[] orderHashes, uint128 gasLimit, bool _payInLzToken) view returns(uint256 nativeFee, uint256 lzTokenFee)
func (_Spoke *SpokeSession) Quote(_dstEid uint32, orderHashes [][32]byte, gasLimit *big.Int, _payInLzToken bool) (struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}, error) {
	return _Spoke.Contract.Quote(&_Spoke.CallOpts, _dstEid, orderHashes, gasLimit, _payInLzToken)
}

// Quote is a free data retrieval call binding the contract method 0x21b4ae78.
//
// Solidity: function quote(uint32 _dstEid, bytes32[] orderHashes, uint128 gasLimit, bool _payInLzToken) view returns(uint256 nativeFee, uint256 lzTokenFee)
func (_Spoke *SpokeCallerSession) Quote(_dstEid uint32, orderHashes [][32]byte, gasLimit *big.Int, _payInLzToken bool) (struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}, error) {
	return _Spoke.Contract.Quote(&_Spoke.CallOpts, _dstEid, orderHashes, gasLimit, _payInLzToken)
}

// SettlementToStatus is a free data retrieval call binding the contract method 0xa32b52a7.
//
// Solidity: function settlementToStatus(bytes32 ) view returns(uint8)
func (_Spoke *SpokeCaller) SettlementToStatus(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "settlementToStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SettlementToStatus is a free data retrieval call binding the contract method 0xa32b52a7.
//
// Solidity: function settlementToStatus(bytes32 ) view returns(uint8)
func (_Spoke *SpokeSession) SettlementToStatus(arg0 [32]byte) (uint8, error) {
	return _Spoke.Contract.SettlementToStatus(&_Spoke.CallOpts, arg0)
}

// SettlementToStatus is a free data retrieval call binding the contract method 0xa32b52a7.
//
// Solidity: function settlementToStatus(bytes32 ) view returns(uint8)
func (_Spoke *SpokeCallerSession) SettlementToStatus(arg0 [32]byte) (uint8, error) {
	return _Spoke.Contract.SettlementToStatus(&_Spoke.CallOpts, arg0)
}

// SquidMulticall is a free data retrieval call binding the contract method 0x59ce62e9.
//
// Solidity: function squidMulticall() view returns(address)
func (_Spoke *SpokeCaller) SquidMulticall(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "squidMulticall")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SquidMulticall is a free data retrieval call binding the contract method 0x59ce62e9.
//
// Solidity: function squidMulticall() view returns(address)
func (_Spoke *SpokeSession) SquidMulticall() (common.Address, error) {
	return _Spoke.Contract.SquidMulticall(&_Spoke.CallOpts)
}

// SquidMulticall is a free data retrieval call binding the contract method 0x59ce62e9.
//
// Solidity: function squidMulticall() view returns(address)
func (_Spoke *SpokeCallerSession) SquidMulticall() (common.Address, error) {
	return _Spoke.Contract.SquidMulticall(&_Spoke.CallOpts)
}

// TokenToCollectedFees is a free data retrieval call binding the contract method 0xdb715f7b.
//
// Solidity: function tokenToCollectedFees(address ) view returns(uint256)
func (_Spoke *SpokeCaller) TokenToCollectedFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "tokenToCollectedFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToCollectedFees is a free data retrieval call binding the contract method 0xdb715f7b.
//
// Solidity: function tokenToCollectedFees(address ) view returns(uint256)
func (_Spoke *SpokeSession) TokenToCollectedFees(arg0 common.Address) (*big.Int, error) {
	return _Spoke.Contract.TokenToCollectedFees(&_Spoke.CallOpts, arg0)
}

// TokenToCollectedFees is a free data retrieval call binding the contract method 0xdb715f7b.
//
// Solidity: function tokenToCollectedFees(address ) view returns(uint256)
func (_Spoke *SpokeCallerSession) TokenToCollectedFees(arg0 common.Address) (*big.Int, error) {
	return _Spoke.Contract.TokenToCollectedFees(&_Spoke.CallOpts, arg0)
}

// TrustedAddress is a free data retrieval call binding the contract method 0x477aedc7.
//
// Solidity: function trustedAddress(string chain) view returns(string trustedAddress_)
func (_Spoke *SpokeCaller) TrustedAddress(opts *bind.CallOpts, chain string) (string, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "trustedAddress", chain)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedAddress is a free data retrieval call binding the contract method 0x477aedc7.
//
// Solidity: function trustedAddress(string chain) view returns(string trustedAddress_)
func (_Spoke *SpokeSession) TrustedAddress(chain string) (string, error) {
	return _Spoke.Contract.TrustedAddress(&_Spoke.CallOpts, chain)
}

// TrustedAddress is a free data retrieval call binding the contract method 0x477aedc7.
//
// Solidity: function trustedAddress(string chain) view returns(string trustedAddress_)
func (_Spoke *SpokeCallerSession) TrustedAddress(chain string) (string, error) {
	return _Spoke.Contract.TrustedAddress(&_Spoke.CallOpts, chain)
}

// TrustedAddressHash is a free data retrieval call binding the contract method 0xffd5982a.
//
// Solidity: function trustedAddressHash(string chain) view returns(bytes32 trustedAddressHash_)
func (_Spoke *SpokeCaller) TrustedAddressHash(opts *bind.CallOpts, chain string) ([32]byte, error) {
	var out []interface{}
	err := _Spoke.contract.Call(opts, &out, "trustedAddressHash", chain)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TrustedAddressHash is a free data retrieval call binding the contract method 0xffd5982a.
//
// Solidity: function trustedAddressHash(string chain) view returns(bytes32 trustedAddressHash_)
func (_Spoke *SpokeSession) TrustedAddressHash(chain string) ([32]byte, error) {
	return _Spoke.Contract.TrustedAddressHash(&_Spoke.CallOpts, chain)
}

// TrustedAddressHash is a free data retrieval call binding the contract method 0xffd5982a.
//
// Solidity: function trustedAddressHash(string chain) view returns(bytes32 trustedAddressHash_)
func (_Spoke *SpokeCallerSession) TrustedAddressHash(chain string) ([32]byte, error) {
	return _Spoke.Contract.TrustedAddressHash(&_Spoke.CallOpts, chain)
}

// BatchFillOrder is a paid mutator transaction binding the contract method 0x36532441.
//
// Solidity: function batchFillOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32)[] orders, (uint8,address,uint256,bytes,bytes)[][] calls) payable returns()
func (_Spoke *SpokeTransactor) BatchFillOrder(opts *bind.TransactOpts, orders []ISpokeOrder, calls [][]ISquidMulticallCall) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "batchFillOrder", orders, calls)
}

// BatchFillOrder is a paid mutator transaction binding the contract method 0x36532441.
//
// Solidity: function batchFillOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32)[] orders, (uint8,address,uint256,bytes,bytes)[][] calls) payable returns()
func (_Spoke *SpokeSession) BatchFillOrder(orders []ISpokeOrder, calls [][]ISquidMulticallCall) (*types.Transaction, error) {
	return _Spoke.Contract.BatchFillOrder(&_Spoke.TransactOpts, orders, calls)
}

// BatchFillOrder is a paid mutator transaction binding the contract method 0x36532441.
//
// Solidity: function batchFillOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32)[] orders, (uint8,address,uint256,bytes,bytes)[][] calls) payable returns()
func (_Spoke *SpokeTransactorSession) BatchFillOrder(orders []ISpokeOrder, calls [][]ISquidMulticallCall) (*types.Transaction, error) {
	return _Spoke.Contract.BatchFillOrder(&_Spoke.TransactOpts, orders, calls)
}

// BatchMultiTokenSingleChainSettlements is a paid mutator transaction binding the contract method 0x1191c72e.
//
// Solidity: function batchMultiTokenSingleChainSettlements((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32)[] orders, address[] fromTokens, address filler) payable returns()
func (_Spoke *SpokeTransactor) BatchMultiTokenSingleChainSettlements(opts *bind.TransactOpts, orders []ISpokeOrder, fromTokens []common.Address, filler common.Address) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "batchMultiTokenSingleChainSettlements", orders, fromTokens, filler)
}

// BatchMultiTokenSingleChainSettlements is a paid mutator transaction binding the contract method 0x1191c72e.
//
// Solidity: function batchMultiTokenSingleChainSettlements((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32)[] orders, address[] fromTokens, address filler) payable returns()
func (_Spoke *SpokeSession) BatchMultiTokenSingleChainSettlements(orders []ISpokeOrder, fromTokens []common.Address, filler common.Address) (*types.Transaction, error) {
	return _Spoke.Contract.BatchMultiTokenSingleChainSettlements(&_Spoke.TransactOpts, orders, fromTokens, filler)
}

// BatchMultiTokenSingleChainSettlements is a paid mutator transaction binding the contract method 0x1191c72e.
//
// Solidity: function batchMultiTokenSingleChainSettlements((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32)[] orders, address[] fromTokens, address filler) payable returns()
func (_Spoke *SpokeTransactorSession) BatchMultiTokenSingleChainSettlements(orders []ISpokeOrder, fromTokens []common.Address, filler common.Address) (*types.Transaction, error) {
	return _Spoke.Contract.BatchMultiTokenSingleChainSettlements(&_Spoke.TransactOpts, orders, fromTokens, filler)
}

// CollectFees is a paid mutator transaction binding the contract method 0x58c0f729.
//
// Solidity: function collectFees(address[] tokens) returns()
func (_Spoke *SpokeTransactor) CollectFees(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "collectFees", tokens)
}

// CollectFees is a paid mutator transaction binding the contract method 0x58c0f729.
//
// Solidity: function collectFees(address[] tokens) returns()
func (_Spoke *SpokeSession) CollectFees(tokens []common.Address) (*types.Transaction, error) {
	return _Spoke.Contract.CollectFees(&_Spoke.TransactOpts, tokens)
}

// CollectFees is a paid mutator transaction binding the contract method 0x58c0f729.
//
// Solidity: function collectFees(address[] tokens) returns()
func (_Spoke *SpokeTransactorSession) CollectFees(tokens []common.Address) (*types.Transaction, error) {
	return _Spoke.Contract.CollectFees(&_Spoke.TransactOpts, tokens)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x0d77797c.
//
// Solidity: function createOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order) payable returns()
func (_Spoke *SpokeTransactor) CreateOrder(opts *bind.TransactOpts, order ISpokeOrder) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "createOrder", order)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x0d77797c.
//
// Solidity: function createOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order) payable returns()
func (_Spoke *SpokeSession) CreateOrder(order ISpokeOrder) (*types.Transaction, error) {
	return _Spoke.Contract.CreateOrder(&_Spoke.TransactOpts, order)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x0d77797c.
//
// Solidity: function createOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order) payable returns()
func (_Spoke *SpokeTransactorSession) CreateOrder(order ISpokeOrder) (*types.Transaction, error) {
	return _Spoke.Contract.CreateOrder(&_Spoke.TransactOpts, order)
}

// Execute is a paid mutator transaction binding the contract method 0x49160658.
//
// Solidity: function execute(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload) returns()
func (_Spoke *SpokeTransactor) Execute(opts *bind.TransactOpts, commandId [32]byte, sourceChain string, sourceAddress string, payload []byte) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "execute", commandId, sourceChain, sourceAddress, payload)
}

// Execute is a paid mutator transaction binding the contract method 0x49160658.
//
// Solidity: function execute(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload) returns()
func (_Spoke *SpokeSession) Execute(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte) (*types.Transaction, error) {
	return _Spoke.Contract.Execute(&_Spoke.TransactOpts, commandId, sourceChain, sourceAddress, payload)
}

// Execute is a paid mutator transaction binding the contract method 0x49160658.
//
// Solidity: function execute(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload) returns()
func (_Spoke *SpokeTransactorSession) Execute(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte) (*types.Transaction, error) {
	return _Spoke.Contract.Execute(&_Spoke.TransactOpts, commandId, sourceChain, sourceAddress, payload)
}

// ExecuteWithToken is a paid mutator transaction binding the contract method 0x1a98b2e0.
//
// Solidity: function executeWithToken(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload, string tokenSymbol, uint256 amount) returns()
func (_Spoke *SpokeTransactor) ExecuteWithToken(opts *bind.TransactOpts, commandId [32]byte, sourceChain string, sourceAddress string, payload []byte, tokenSymbol string, amount *big.Int) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "executeWithToken", commandId, sourceChain, sourceAddress, payload, tokenSymbol, amount)
}

// ExecuteWithToken is a paid mutator transaction binding the contract method 0x1a98b2e0.
//
// Solidity: function executeWithToken(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload, string tokenSymbol, uint256 amount) returns()
func (_Spoke *SpokeSession) ExecuteWithToken(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte, tokenSymbol string, amount *big.Int) (*types.Transaction, error) {
	return _Spoke.Contract.ExecuteWithToken(&_Spoke.TransactOpts, commandId, sourceChain, sourceAddress, payload, tokenSymbol, amount)
}

// ExecuteWithToken is a paid mutator transaction binding the contract method 0x1a98b2e0.
//
// Solidity: function executeWithToken(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload, string tokenSymbol, uint256 amount) returns()
func (_Spoke *SpokeTransactorSession) ExecuteWithToken(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte, tokenSymbol string, amount *big.Int) (*types.Transaction, error) {
	return _Spoke.Contract.ExecuteWithToken(&_Spoke.TransactOpts, commandId, sourceChain, sourceAddress, payload, tokenSymbol, amount)
}

// FillOrder is a paid mutator transaction binding the contract method 0xaab59a09.
//
// Solidity: function fillOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order, (uint8,address,uint256,bytes,bytes)[] calls) payable returns()
func (_Spoke *SpokeTransactor) FillOrder(opts *bind.TransactOpts, order ISpokeOrder, calls []ISquidMulticallCall) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "fillOrder", order, calls)
}

// FillOrder is a paid mutator transaction binding the contract method 0xaab59a09.
//
// Solidity: function fillOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order, (uint8,address,uint256,bytes,bytes)[] calls) payable returns()
func (_Spoke *SpokeSession) FillOrder(order ISpokeOrder, calls []ISquidMulticallCall) (*types.Transaction, error) {
	return _Spoke.Contract.FillOrder(&_Spoke.TransactOpts, order, calls)
}

// FillOrder is a paid mutator transaction binding the contract method 0xaab59a09.
//
// Solidity: function fillOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order, (uint8,address,uint256,bytes,bytes)[] calls) payable returns()
func (_Spoke *SpokeTransactorSession) FillOrder(order ISpokeOrder, calls []ISquidMulticallCall) (*types.Transaction, error) {
	return _Spoke.Contract.FillOrder(&_Spoke.TransactOpts, order, calls)
}

// ForwardSettlements is a paid mutator transaction binding the contract method 0x0630dea4.
//
// Solidity: function forwardSettlements(bytes32[] orderHashes, uint256 lzFee, uint128 gasLimit, uint8 provider) payable returns()
func (_Spoke *SpokeTransactor) ForwardSettlements(opts *bind.TransactOpts, orderHashes [][32]byte, lzFee *big.Int, gasLimit *big.Int, provider uint8) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "forwardSettlements", orderHashes, lzFee, gasLimit, provider)
}

// ForwardSettlements is a paid mutator transaction binding the contract method 0x0630dea4.
//
// Solidity: function forwardSettlements(bytes32[] orderHashes, uint256 lzFee, uint128 gasLimit, uint8 provider) payable returns()
func (_Spoke *SpokeSession) ForwardSettlements(orderHashes [][32]byte, lzFee *big.Int, gasLimit *big.Int, provider uint8) (*types.Transaction, error) {
	return _Spoke.Contract.ForwardSettlements(&_Spoke.TransactOpts, orderHashes, lzFee, gasLimit, provider)
}

// ForwardSettlements is a paid mutator transaction binding the contract method 0x0630dea4.
//
// Solidity: function forwardSettlements(bytes32[] orderHashes, uint256 lzFee, uint128 gasLimit, uint8 provider) payable returns()
func (_Spoke *SpokeTransactorSession) ForwardSettlements(orderHashes [][32]byte, lzFee *big.Int, gasLimit *big.Int, provider uint8) (*types.Transaction, error) {
	return _Spoke.Contract.ForwardSettlements(&_Spoke.TransactOpts, orderHashes, lzFee, gasLimit, provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xb98ef275.
//
// Solidity: function initialize(address _axelarGateway, address _axelarGasService, address _permit2, address _squidMulticall, address _feeCollector, string _hubChainName, string _hubAddress, address _endpoint, address _owner, address _hub, uint32 _hubEndpoint) returns()
func (_Spoke *SpokeTransactor) Initialize(opts *bind.TransactOpts, _axelarGateway common.Address, _axelarGasService common.Address, _permit2 common.Address, _squidMulticall common.Address, _feeCollector common.Address, _hubChainName string, _hubAddress string, _endpoint common.Address, _owner common.Address, _hub common.Address, _hubEndpoint uint32) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "initialize", _axelarGateway, _axelarGasService, _permit2, _squidMulticall, _feeCollector, _hubChainName, _hubAddress, _endpoint, _owner, _hub, _hubEndpoint)
}

// Initialize is a paid mutator transaction binding the contract method 0xb98ef275.
//
// Solidity: function initialize(address _axelarGateway, address _axelarGasService, address _permit2, address _squidMulticall, address _feeCollector, string _hubChainName, string _hubAddress, address _endpoint, address _owner, address _hub, uint32 _hubEndpoint) returns()
func (_Spoke *SpokeSession) Initialize(_axelarGateway common.Address, _axelarGasService common.Address, _permit2 common.Address, _squidMulticall common.Address, _feeCollector common.Address, _hubChainName string, _hubAddress string, _endpoint common.Address, _owner common.Address, _hub common.Address, _hubEndpoint uint32) (*types.Transaction, error) {
	return _Spoke.Contract.Initialize(&_Spoke.TransactOpts, _axelarGateway, _axelarGasService, _permit2, _squidMulticall, _feeCollector, _hubChainName, _hubAddress, _endpoint, _owner, _hub, _hubEndpoint)
}

// Initialize is a paid mutator transaction binding the contract method 0xb98ef275.
//
// Solidity: function initialize(address _axelarGateway, address _axelarGasService, address _permit2, address _squidMulticall, address _feeCollector, string _hubChainName, string _hubAddress, address _endpoint, address _owner, address _hub, uint32 _hubEndpoint) returns()
func (_Spoke *SpokeTransactorSession) Initialize(_axelarGateway common.Address, _axelarGasService common.Address, _permit2 common.Address, _squidMulticall common.Address, _feeCollector common.Address, _hubChainName string, _hubAddress string, _endpoint common.Address, _owner common.Address, _hub common.Address, _hubEndpoint uint32) (*types.Transaction, error) {
	return _Spoke.Contract.Initialize(&_Spoke.TransactOpts, _axelarGateway, _axelarGasService, _permit2, _squidMulticall, _feeCollector, _hubChainName, _hubAddress, _endpoint, _owner, _hub, _hubEndpoint)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Spoke *SpokeTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Spoke *SpokeSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Spoke.Contract.LzReceive(&_Spoke.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Spoke *SpokeTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Spoke.Contract.LzReceive(&_Spoke.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// RefundOrder is a paid mutator transaction binding the contract method 0x1e44fb97.
//
// Solidity: function refundOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order) returns()
func (_Spoke *SpokeTransactor) RefundOrder(opts *bind.TransactOpts, order ISpokeOrder) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "refundOrder", order)
}

// RefundOrder is a paid mutator transaction binding the contract method 0x1e44fb97.
//
// Solidity: function refundOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order) returns()
func (_Spoke *SpokeSession) RefundOrder(order ISpokeOrder) (*types.Transaction, error) {
	return _Spoke.Contract.RefundOrder(&_Spoke.TransactOpts, order)
}

// RefundOrder is a paid mutator transaction binding the contract method 0x1e44fb97.
//
// Solidity: function refundOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order) returns()
func (_Spoke *SpokeTransactorSession) RefundOrder(order ISpokeOrder) (*types.Transaction, error) {
	return _Spoke.Contract.RefundOrder(&_Spoke.TransactOpts, order)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Spoke *SpokeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Spoke *SpokeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Spoke.Contract.RenounceOwnership(&_Spoke.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Spoke *SpokeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Spoke.Contract.RenounceOwnership(&_Spoke.TransactOpts)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Spoke *SpokeTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Spoke *SpokeSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Spoke.Contract.SetDelegate(&_Spoke.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Spoke *SpokeTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Spoke.Contract.SetDelegate(&_Spoke.TransactOpts, _delegate)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Spoke *SpokeTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Spoke *SpokeSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Spoke.Contract.SetPeer(&_Spoke.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Spoke *SpokeTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Spoke.Contract.SetPeer(&_Spoke.TransactOpts, _eid, _peer)
}

// SponsorOrder is a paid mutator transaction binding the contract method 0x79db29c8.
//
// Solidity: function sponsorOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order, bytes signature) returns()
func (_Spoke *SpokeTransactor) SponsorOrder(opts *bind.TransactOpts, order ISpokeOrder, signature []byte) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "sponsorOrder", order, signature)
}

// SponsorOrder is a paid mutator transaction binding the contract method 0x79db29c8.
//
// Solidity: function sponsorOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order, bytes signature) returns()
func (_Spoke *SpokeSession) SponsorOrder(order ISpokeOrder, signature []byte) (*types.Transaction, error) {
	return _Spoke.Contract.SponsorOrder(&_Spoke.TransactOpts, order, signature)
}

// SponsorOrder is a paid mutator transaction binding the contract method 0x79db29c8.
//
// Solidity: function sponsorOrder((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order, bytes signature) returns()
func (_Spoke *SpokeTransactorSession) SponsorOrder(order ISpokeOrder, signature []byte) (*types.Transaction, error) {
	return _Spoke.Contract.SponsorOrder(&_Spoke.TransactOpts, order, signature)
}

// SponsorOrderUsingPermit2 is a paid mutator transaction binding the contract method 0x2e3d6cf9.
//
// Solidity: function sponsorOrderUsingPermit2((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order, ((address,uint256),uint256,uint256) permit, bytes signature) returns()
func (_Spoke *SpokeTransactor) SponsorOrderUsingPermit2(opts *bind.TransactOpts, order ISpokeOrder, permit ISignatureTransferPermitTransferFrom, signature []byte) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "sponsorOrderUsingPermit2", order, permit, signature)
}

// SponsorOrderUsingPermit2 is a paid mutator transaction binding the contract method 0x2e3d6cf9.
//
// Solidity: function sponsorOrderUsingPermit2((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order, ((address,uint256),uint256,uint256) permit, bytes signature) returns()
func (_Spoke *SpokeSession) SponsorOrderUsingPermit2(order ISpokeOrder, permit ISignatureTransferPermitTransferFrom, signature []byte) (*types.Transaction, error) {
	return _Spoke.Contract.SponsorOrderUsingPermit2(&_Spoke.TransactOpts, order, permit, signature)
}

// SponsorOrderUsingPermit2 is a paid mutator transaction binding the contract method 0x2e3d6cf9.
//
// Solidity: function sponsorOrderUsingPermit2((address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order, ((address,uint256),uint256,uint256) permit, bytes signature) returns()
func (_Spoke *SpokeTransactorSession) SponsorOrderUsingPermit2(order ISpokeOrder, permit ISignatureTransferPermitTransferFrom, signature []byte) (*types.Transaction, error) {
	return _Spoke.Contract.SponsorOrderUsingPermit2(&_Spoke.TransactOpts, order, permit, signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Spoke *SpokeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Spoke.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Spoke *SpokeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Spoke.Contract.TransferOwnership(&_Spoke.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Spoke *SpokeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Spoke.Contract.TransferOwnership(&_Spoke.TransactOpts, newOwner)
}

// SpokeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Spoke contract.
type SpokeEIP712DomainChangedIterator struct {
	Event *SpokeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *SpokeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeEIP712DomainChanged)
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
		it.Event = new(SpokeEIP712DomainChanged)
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
func (it *SpokeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeEIP712DomainChanged represents a EIP712DomainChanged event raised by the Spoke contract.
type SpokeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Spoke *SpokeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*SpokeEIP712DomainChangedIterator, error) {

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &SpokeEIP712DomainChangedIterator{contract: _Spoke.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Spoke *SpokeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *SpokeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeEIP712DomainChanged)
				if err := _Spoke.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Spoke *SpokeFilterer) ParseEIP712DomainChanged(log types.Log) (*SpokeEIP712DomainChanged, error) {
	event := new(SpokeEIP712DomainChanged)
	if err := _Spoke.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeFeesCollectedIterator is returned from FilterFeesCollected and is used to iterate over the raw logs and unpacked data for FeesCollected events raised by the Spoke contract.
type SpokeFeesCollectedIterator struct {
	Event *SpokeFeesCollected // Event containing the contract specifics and raw log

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
func (it *SpokeFeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeFeesCollected)
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
		it.Event = new(SpokeFeesCollected)
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
func (it *SpokeFeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeFeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeFeesCollected represents a FeesCollected event raised by the Spoke contract.
type SpokeFeesCollected struct {
	FeeCollector common.Address
	Token        common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeesCollected is a free log retrieval operation binding the contract event 0x9bcb6d1f38f6800906185471a11ede9a8e16200853225aa62558db6076490f2d.
//
// Solidity: event FeesCollected(address indexed feeCollector, address indexed token, uint256 indexed amount)
func (_Spoke *SpokeFilterer) FilterFeesCollected(opts *bind.FilterOpts, feeCollector []common.Address, token []common.Address, amount []*big.Int) (*SpokeFeesCollectedIterator, error) {

	var feeCollectorRule []interface{}
	for _, feeCollectorItem := range feeCollector {
		feeCollectorRule = append(feeCollectorRule, feeCollectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "FeesCollected", feeCollectorRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &SpokeFeesCollectedIterator{contract: _Spoke.contract, event: "FeesCollected", logs: logs, sub: sub}, nil
}

// WatchFeesCollected is a free log subscription operation binding the contract event 0x9bcb6d1f38f6800906185471a11ede9a8e16200853225aa62558db6076490f2d.
//
// Solidity: event FeesCollected(address indexed feeCollector, address indexed token, uint256 indexed amount)
func (_Spoke *SpokeFilterer) WatchFeesCollected(opts *bind.WatchOpts, sink chan<- *SpokeFeesCollected, feeCollector []common.Address, token []common.Address, amount []*big.Int) (event.Subscription, error) {

	var feeCollectorRule []interface{}
	for _, feeCollectorItem := range feeCollector {
		feeCollectorRule = append(feeCollectorRule, feeCollectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "FeesCollected", feeCollectorRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeFeesCollected)
				if err := _Spoke.contract.UnpackLog(event, "FeesCollected", log); err != nil {
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

// ParseFeesCollected is a log parse operation binding the contract event 0x9bcb6d1f38f6800906185471a11ede9a8e16200853225aa62558db6076490f2d.
//
// Solidity: event FeesCollected(address indexed feeCollector, address indexed token, uint256 indexed amount)
func (_Spoke *SpokeFilterer) ParseFeesCollected(log types.Log) (*SpokeFeesCollected, error) {
	event := new(SpokeFeesCollected)
	if err := _Spoke.contract.UnpackLog(event, "FeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Spoke contract.
type SpokeInitializedIterator struct {
	Event *SpokeInitialized // Event containing the contract specifics and raw log

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
func (it *SpokeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeInitialized)
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
		it.Event = new(SpokeInitialized)
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
func (it *SpokeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeInitialized represents a Initialized event raised by the Spoke contract.
type SpokeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Spoke *SpokeFilterer) FilterInitialized(opts *bind.FilterOpts) (*SpokeInitializedIterator, error) {

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SpokeInitializedIterator{contract: _Spoke.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Spoke *SpokeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SpokeInitialized) (event.Subscription, error) {

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeInitialized)
				if err := _Spoke.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Spoke *SpokeFilterer) ParseInitialized(log types.Log) (*SpokeInitialized, error) {
	event := new(SpokeInitialized)
	if err := _Spoke.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the Spoke contract.
type SpokeOrderCreatedIterator struct {
	Event *SpokeOrderCreated // Event containing the contract specifics and raw log

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
func (it *SpokeOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeOrderCreated)
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
		it.Event = new(SpokeOrderCreated)
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
func (it *SpokeOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeOrderCreated represents a OrderCreated event raised by the Spoke contract.
type SpokeOrderCreated struct {
	OrderHash [32]byte
	Order     ISpokeOrder
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x181de28643611afcf1cb4c095a1ef99c157e78437294f478c978e4a56e1ca77e.
//
// Solidity: event OrderCreated(bytes32 indexed orderHash, (address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order)
func (_Spoke *SpokeFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderHash [][32]byte) (*SpokeOrderCreatedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "OrderCreated", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &SpokeOrderCreatedIterator{contract: _Spoke.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x181de28643611afcf1cb4c095a1ef99c157e78437294f478c978e4a56e1ca77e.
//
// Solidity: event OrderCreated(bytes32 indexed orderHash, (address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order)
func (_Spoke *SpokeFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *SpokeOrderCreated, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "OrderCreated", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeOrderCreated)
				if err := _Spoke.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x181de28643611afcf1cb4c095a1ef99c157e78437294f478c978e4a56e1ca77e.
//
// Solidity: event OrderCreated(bytes32 indexed orderHash, (address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order)
func (_Spoke *SpokeFilterer) ParseOrderCreated(log types.Log) (*SpokeOrderCreated, error) {
	event := new(SpokeOrderCreated)
	if err := _Spoke.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the Spoke contract.
type SpokeOrderFilledIterator struct {
	Event *SpokeOrderFilled // Event containing the contract specifics and raw log

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
func (it *SpokeOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeOrderFilled)
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
		it.Event = new(SpokeOrderFilled)
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
func (it *SpokeOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeOrderFilled represents a OrderFilled event raised by the Spoke contract.
type SpokeOrderFilled struct {
	OrderHash [32]byte
	Order     ISpokeOrder
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0x6955fd9b2a7639a9baac024897cad7007b45ffa74cbfe9582d58401ff6b977b7.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, (address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order)
func (_Spoke *SpokeFilterer) FilterOrderFilled(opts *bind.FilterOpts, orderHash [][32]byte) (*SpokeOrderFilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "OrderFilled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &SpokeOrderFilledIterator{contract: _Spoke.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0x6955fd9b2a7639a9baac024897cad7007b45ffa74cbfe9582d58401ff6b977b7.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, (address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order)
func (_Spoke *SpokeFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *SpokeOrderFilled, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "OrderFilled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeOrderFilled)
				if err := _Spoke.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0x6955fd9b2a7639a9baac024897cad7007b45ffa74cbfe9582d58401ff6b977b7.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, (address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes32) order)
func (_Spoke *SpokeFilterer) ParseOrderFilled(log types.Log) (*SpokeOrderFilled, error) {
	event := new(SpokeOrderFilled)
	if err := _Spoke.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeOrderRefundedIterator is returned from FilterOrderRefunded and is used to iterate over the raw logs and unpacked data for OrderRefunded events raised by the Spoke contract.
type SpokeOrderRefundedIterator struct {
	Event *SpokeOrderRefunded // Event containing the contract specifics and raw log

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
func (it *SpokeOrderRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeOrderRefunded)
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
		it.Event = new(SpokeOrderRefunded)
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
func (it *SpokeOrderRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeOrderRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeOrderRefunded represents a OrderRefunded event raised by the Spoke contract.
type SpokeOrderRefunded struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderRefunded is a free log retrieval operation binding the contract event 0xa60671d8537ed193e567f86ddf28cf35dc67073b5ad80a2d41359cfa78db0a1e.
//
// Solidity: event OrderRefunded(bytes32 indexed orderHash)
func (_Spoke *SpokeFilterer) FilterOrderRefunded(opts *bind.FilterOpts, orderHash [][32]byte) (*SpokeOrderRefundedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "OrderRefunded", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &SpokeOrderRefundedIterator{contract: _Spoke.contract, event: "OrderRefunded", logs: logs, sub: sub}, nil
}

// WatchOrderRefunded is a free log subscription operation binding the contract event 0xa60671d8537ed193e567f86ddf28cf35dc67073b5ad80a2d41359cfa78db0a1e.
//
// Solidity: event OrderRefunded(bytes32 indexed orderHash)
func (_Spoke *SpokeFilterer) WatchOrderRefunded(opts *bind.WatchOpts, sink chan<- *SpokeOrderRefunded, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "OrderRefunded", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeOrderRefunded)
				if err := _Spoke.contract.UnpackLog(event, "OrderRefunded", log); err != nil {
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

// ParseOrderRefunded is a log parse operation binding the contract event 0xa60671d8537ed193e567f86ddf28cf35dc67073b5ad80a2d41359cfa78db0a1e.
//
// Solidity: event OrderRefunded(bytes32 indexed orderHash)
func (_Spoke *SpokeFilterer) ParseOrderRefunded(log types.Log) (*SpokeOrderRefunded, error) {
	event := new(SpokeOrderRefunded)
	if err := _Spoke.contract.UnpackLog(event, "OrderRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Spoke contract.
type SpokeOwnershipTransferredIterator struct {
	Event *SpokeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SpokeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeOwnershipTransferred)
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
		it.Event = new(SpokeOwnershipTransferred)
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
func (it *SpokeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeOwnershipTransferred represents a OwnershipTransferred event raised by the Spoke contract.
type SpokeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Spoke *SpokeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SpokeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SpokeOwnershipTransferredIterator{contract: _Spoke.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Spoke *SpokeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SpokeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeOwnershipTransferred)
				if err := _Spoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Spoke *SpokeFilterer) ParseOwnershipTransferred(log types.Log) (*SpokeOwnershipTransferred, error) {
	event := new(SpokeOwnershipTransferred)
	if err := _Spoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokePeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the Spoke contract.
type SpokePeerSetIterator struct {
	Event *SpokePeerSet // Event containing the contract specifics and raw log

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
func (it *SpokePeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokePeerSet)
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
		it.Event = new(SpokePeerSet)
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
func (it *SpokePeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokePeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokePeerSet represents a PeerSet event raised by the Spoke contract.
type SpokePeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Spoke *SpokeFilterer) FilterPeerSet(opts *bind.FilterOpts) (*SpokePeerSetIterator, error) {

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &SpokePeerSetIterator{contract: _Spoke.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Spoke *SpokeFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *SpokePeerSet) (event.Subscription, error) {

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokePeerSet)
				if err := _Spoke.contract.UnpackLog(event, "PeerSet", log); err != nil {
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

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Spoke *SpokeFilterer) ParsePeerSet(log types.Log) (*SpokePeerSet, error) {
	event := new(SpokePeerSet)
	if err := _Spoke.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeSettlementForwardedIterator is returned from FilterSettlementForwarded and is used to iterate over the raw logs and unpacked data for SettlementForwarded events raised by the Spoke contract.
type SpokeSettlementForwardedIterator struct {
	Event *SpokeSettlementForwarded // Event containing the contract specifics and raw log

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
func (it *SpokeSettlementForwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeSettlementForwarded)
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
		it.Event = new(SpokeSettlementForwarded)
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
func (it *SpokeSettlementForwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeSettlementForwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeSettlementForwarded represents a SettlementForwarded event raised by the Spoke contract.
type SpokeSettlementForwarded struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSettlementForwarded is a free log retrieval operation binding the contract event 0x69f975bd70ea51b973eb6aff3812f49adf595bd59d6f3d29840d5695cc19ba30.
//
// Solidity: event SettlementForwarded(bytes32 indexed orderHash)
func (_Spoke *SpokeFilterer) FilterSettlementForwarded(opts *bind.FilterOpts, orderHash [][32]byte) (*SpokeSettlementForwardedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "SettlementForwarded", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &SpokeSettlementForwardedIterator{contract: _Spoke.contract, event: "SettlementForwarded", logs: logs, sub: sub}, nil
}

// WatchSettlementForwarded is a free log subscription operation binding the contract event 0x69f975bd70ea51b973eb6aff3812f49adf595bd59d6f3d29840d5695cc19ba30.
//
// Solidity: event SettlementForwarded(bytes32 indexed orderHash)
func (_Spoke *SpokeFilterer) WatchSettlementForwarded(opts *bind.WatchOpts, sink chan<- *SpokeSettlementForwarded, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "SettlementForwarded", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeSettlementForwarded)
				if err := _Spoke.contract.UnpackLog(event, "SettlementForwarded", log); err != nil {
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

// ParseSettlementForwarded is a log parse operation binding the contract event 0x69f975bd70ea51b973eb6aff3812f49adf595bd59d6f3d29840d5695cc19ba30.
//
// Solidity: event SettlementForwarded(bytes32 indexed orderHash)
func (_Spoke *SpokeFilterer) ParseSettlementForwarded(log types.Log) (*SpokeSettlementForwarded, error) {
	event := new(SpokeSettlementForwarded)
	if err := _Spoke.contract.UnpackLog(event, "SettlementForwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeSpokeInitializedIterator is returned from FilterSpokeInitialized and is used to iterate over the raw logs and unpacked data for SpokeInitialized events raised by the Spoke contract.
type SpokeSpokeInitializedIterator struct {
	Event *SpokeSpokeInitialized // Event containing the contract specifics and raw log

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
func (it *SpokeSpokeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeSpokeInitialized)
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
		it.Event = new(SpokeSpokeInitialized)
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
func (it *SpokeSpokeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeSpokeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeSpokeInitialized represents a SpokeInitialized event raised by the Spoke contract.
type SpokeSpokeInitialized struct {
	Gateway        common.Address
	GasService     common.Address
	Permit2        common.Address
	SquidMulticall common.Address
	FeeCollector   common.Address
	HubChainName   string
	HubAddress     string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSpokeInitialized is a free log retrieval operation binding the contract event 0x66b86e9a862484192b40f4e2dc828c11d9e09ca77f9d2c226a96e743f99e4ec8.
//
// Solidity: event SpokeInitialized(address indexed gateway, address indexed gasService, address indexed permit2, address squidMulticall, address feeCollector, string hubChainName, string hubAddress)
func (_Spoke *SpokeFilterer) FilterSpokeInitialized(opts *bind.FilterOpts, gateway []common.Address, gasService []common.Address, permit2 []common.Address) (*SpokeSpokeInitializedIterator, error) {

	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}
	var gasServiceRule []interface{}
	for _, gasServiceItem := range gasService {
		gasServiceRule = append(gasServiceRule, gasServiceItem)
	}
	var permit2Rule []interface{}
	for _, permit2Item := range permit2 {
		permit2Rule = append(permit2Rule, permit2Item)
	}

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "SpokeInitialized", gatewayRule, gasServiceRule, permit2Rule)
	if err != nil {
		return nil, err
	}
	return &SpokeSpokeInitializedIterator{contract: _Spoke.contract, event: "SpokeInitialized", logs: logs, sub: sub}, nil
}

// WatchSpokeInitialized is a free log subscription operation binding the contract event 0x66b86e9a862484192b40f4e2dc828c11d9e09ca77f9d2c226a96e743f99e4ec8.
//
// Solidity: event SpokeInitialized(address indexed gateway, address indexed gasService, address indexed permit2, address squidMulticall, address feeCollector, string hubChainName, string hubAddress)
func (_Spoke *SpokeFilterer) WatchSpokeInitialized(opts *bind.WatchOpts, sink chan<- *SpokeSpokeInitialized, gateway []common.Address, gasService []common.Address, permit2 []common.Address) (event.Subscription, error) {

	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}
	var gasServiceRule []interface{}
	for _, gasServiceItem := range gasService {
		gasServiceRule = append(gasServiceRule, gasServiceItem)
	}
	var permit2Rule []interface{}
	for _, permit2Item := range permit2 {
		permit2Rule = append(permit2Rule, permit2Item)
	}

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "SpokeInitialized", gatewayRule, gasServiceRule, permit2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeSpokeInitialized)
				if err := _Spoke.contract.UnpackLog(event, "SpokeInitialized", log); err != nil {
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

// ParseSpokeInitialized is a log parse operation binding the contract event 0x66b86e9a862484192b40f4e2dc828c11d9e09ca77f9d2c226a96e743f99e4ec8.
//
// Solidity: event SpokeInitialized(address indexed gateway, address indexed gasService, address indexed permit2, address squidMulticall, address feeCollector, string hubChainName, string hubAddress)
func (_Spoke *SpokeFilterer) ParseSpokeInitialized(log types.Log) (*SpokeSpokeInitialized, error) {
	event := new(SpokeSpokeInitialized)
	if err := _Spoke.contract.UnpackLog(event, "SpokeInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeTokensReleasedIterator is returned from FilterTokensReleased and is used to iterate over the raw logs and unpacked data for TokensReleased events raised by the Spoke contract.
type SpokeTokensReleasedIterator struct {
	Event *SpokeTokensReleased // Event containing the contract specifics and raw log

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
func (it *SpokeTokensReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeTokensReleased)
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
		it.Event = new(SpokeTokensReleased)
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
func (it *SpokeTokensReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeTokensReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeTokensReleased represents a TokensReleased event raised by the Spoke contract.
type SpokeTokensReleased struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensReleased is a free log retrieval operation binding the contract event 0xd48052bf92f3eec93ecdeeec72ea80e1071c926cb4d6e5a37ee71be8a0ce9a10.
//
// Solidity: event TokensReleased(bytes32 indexed orderHash)
func (_Spoke *SpokeFilterer) FilterTokensReleased(opts *bind.FilterOpts, orderHash [][32]byte) (*SpokeTokensReleasedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "TokensReleased", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &SpokeTokensReleasedIterator{contract: _Spoke.contract, event: "TokensReleased", logs: logs, sub: sub}, nil
}

// WatchTokensReleased is a free log subscription operation binding the contract event 0xd48052bf92f3eec93ecdeeec72ea80e1071c926cb4d6e5a37ee71be8a0ce9a10.
//
// Solidity: event TokensReleased(bytes32 indexed orderHash)
func (_Spoke *SpokeFilterer) WatchTokensReleased(opts *bind.WatchOpts, sink chan<- *SpokeTokensReleased, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "TokensReleased", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeTokensReleased)
				if err := _Spoke.contract.UnpackLog(event, "TokensReleased", log); err != nil {
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

// ParseTokensReleased is a log parse operation binding the contract event 0xd48052bf92f3eec93ecdeeec72ea80e1071c926cb4d6e5a37ee71be8a0ce9a10.
//
// Solidity: event TokensReleased(bytes32 indexed orderHash)
func (_Spoke *SpokeFilterer) ParseTokensReleased(log types.Log) (*SpokeTokensReleased, error) {
	event := new(SpokeTokensReleased)
	if err := _Spoke.contract.UnpackLog(event, "TokensReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeTrustedAddressRemovedIterator is returned from FilterTrustedAddressRemoved and is used to iterate over the raw logs and unpacked data for TrustedAddressRemoved events raised by the Spoke contract.
type SpokeTrustedAddressRemovedIterator struct {
	Event *SpokeTrustedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *SpokeTrustedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeTrustedAddressRemoved)
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
		it.Event = new(SpokeTrustedAddressRemoved)
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
func (it *SpokeTrustedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeTrustedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeTrustedAddressRemoved represents a TrustedAddressRemoved event raised by the Spoke contract.
type SpokeTrustedAddressRemoved struct {
	Chain string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTrustedAddressRemoved is a free log retrieval operation binding the contract event 0xf9400637a329865492b8d0d4dba4eafc7e8d5d0fae5e27b56766816d2ae1b2ca.
//
// Solidity: event TrustedAddressRemoved(string chain)
func (_Spoke *SpokeFilterer) FilterTrustedAddressRemoved(opts *bind.FilterOpts) (*SpokeTrustedAddressRemovedIterator, error) {

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "TrustedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &SpokeTrustedAddressRemovedIterator{contract: _Spoke.contract, event: "TrustedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedAddressRemoved is a free log subscription operation binding the contract event 0xf9400637a329865492b8d0d4dba4eafc7e8d5d0fae5e27b56766816d2ae1b2ca.
//
// Solidity: event TrustedAddressRemoved(string chain)
func (_Spoke *SpokeFilterer) WatchTrustedAddressRemoved(opts *bind.WatchOpts, sink chan<- *SpokeTrustedAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "TrustedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeTrustedAddressRemoved)
				if err := _Spoke.contract.UnpackLog(event, "TrustedAddressRemoved", log); err != nil {
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

// ParseTrustedAddressRemoved is a log parse operation binding the contract event 0xf9400637a329865492b8d0d4dba4eafc7e8d5d0fae5e27b56766816d2ae1b2ca.
//
// Solidity: event TrustedAddressRemoved(string chain)
func (_Spoke *SpokeFilterer) ParseTrustedAddressRemoved(log types.Log) (*SpokeTrustedAddressRemoved, error) {
	event := new(SpokeTrustedAddressRemoved)
	if err := _Spoke.contract.UnpackLog(event, "TrustedAddressRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpokeTrustedAddressSetIterator is returned from FilterTrustedAddressSet and is used to iterate over the raw logs and unpacked data for TrustedAddressSet events raised by the Spoke contract.
type SpokeTrustedAddressSetIterator struct {
	Event *SpokeTrustedAddressSet // Event containing the contract specifics and raw log

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
func (it *SpokeTrustedAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpokeTrustedAddressSet)
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
		it.Event = new(SpokeTrustedAddressSet)
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
func (it *SpokeTrustedAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpokeTrustedAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpokeTrustedAddressSet represents a TrustedAddressSet event raised by the Spoke contract.
type SpokeTrustedAddressSet struct {
	Chain   string
	Address string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTrustedAddressSet is a free log retrieval operation binding the contract event 0xdb6b260ea45f7fe513e1d3b8c21017a29e3a41610e95aefb8862b81c69aec61c.
//
// Solidity: event TrustedAddressSet(string chain, string address_)
func (_Spoke *SpokeFilterer) FilterTrustedAddressSet(opts *bind.FilterOpts) (*SpokeTrustedAddressSetIterator, error) {

	logs, sub, err := _Spoke.contract.FilterLogs(opts, "TrustedAddressSet")
	if err != nil {
		return nil, err
	}
	return &SpokeTrustedAddressSetIterator{contract: _Spoke.contract, event: "TrustedAddressSet", logs: logs, sub: sub}, nil
}

// WatchTrustedAddressSet is a free log subscription operation binding the contract event 0xdb6b260ea45f7fe513e1d3b8c21017a29e3a41610e95aefb8862b81c69aec61c.
//
// Solidity: event TrustedAddressSet(string chain, string address_)
func (_Spoke *SpokeFilterer) WatchTrustedAddressSet(opts *bind.WatchOpts, sink chan<- *SpokeTrustedAddressSet) (event.Subscription, error) {

	logs, sub, err := _Spoke.contract.WatchLogs(opts, "TrustedAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpokeTrustedAddressSet)
				if err := _Spoke.contract.UnpackLog(event, "TrustedAddressSet", log); err != nil {
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

// ParseTrustedAddressSet is a log parse operation binding the contract event 0xdb6b260ea45f7fe513e1d3b8c21017a29e3a41610e95aefb8862b81c69aec61c.
//
// Solidity: event TrustedAddressSet(string chain, string address_)
func (_Spoke *SpokeFilterer) ParseTrustedAddressSet(log types.Log) (*SpokeTrustedAddressSet, error) {
	event := new(SpokeTrustedAddressSet)
	if err := _Spoke.contract.UnpackLog(event, "TrustedAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
