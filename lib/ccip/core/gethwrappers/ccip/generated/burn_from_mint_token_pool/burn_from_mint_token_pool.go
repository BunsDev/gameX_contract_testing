// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_from_mint_token_pool

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

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

type PoolLockOrBurnInV1 struct {
	Receiver            []byte
	RemoteChainSelector uint64
	OriginalSender      common.Address
	Amount              *big.Int
	LocalToken          common.Address
}

type PoolLockOrBurnOutV1 struct {
	DestPoolAddress []byte
	DestPoolData    []byte
}

type PoolReleaseOrMintInV1 struct {
	OriginalSender      []byte
	RemoteChainSelector uint64
	Receiver            common.Address
	Amount              *big.Int
	SourcePoolAddress   []byte
	SourcePoolData      []byte
	OffchainTokenData   []byte
}

type PoolReleaseOrMintOutV1 struct {
	LocalToken        common.Address
	DestinationAmount *big.Int
}

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type RateLimiterTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

type TokenPoolChainUpdate struct {
	RemoteChainSelector       uint64
	Allowed                   bool
	RemotePoolAddress         []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

var BurnFromMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBurnMintERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotARampOnRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"DisabledNonZeroRateLimit\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"InvalidRatelimitRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidSourcePoolAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"NonExistentChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitMustBeDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"previousPoolAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.ChainUpdate[]\",\"name\":\"chains\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentInboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentOutboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemotePool\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRmnProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedChains\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"internalType\":\"structPool.LockOrBurnInV1\",\"name\":\"lockOrBurnIn\",\"type\":\"tuple\"}],\"name\":\"lockOrBurn\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"destPoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destPoolData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.LockOrBurnOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainTokenData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.ReleaseOrMintInV1\",\"name\":\"releaseOrMintIn\",\"type\":\"tuple\"}],\"name\":\"releaseOrMint\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundConfig\",\"type\":\"tuple\"}],\"name\":\"setChainRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"setRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162004220380380620042208339810160408190526200003491620008c0565b8383838333806000816200008f5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c257620000c2816200018f565b5050506001600160a01b0384161580620000e357506001600160a01b038116155b80620000f657506001600160a01b038216155b1562000115576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b0384811660805282811660a052600480546001600160a01b031916918316919091179055825115801560c05262000168576040805160008152602081019091526200016890846200023a565b5062000185925050506001600160a01b0385163060001962000397565b5050505062000afc565b336001600160a01b03821603620001e95760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000086565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60c0516200025b576040516335f4a7b360e01b815260040160405180910390fd5b60005b8251811015620002e65760008382815181106200027f576200027f620009d0565b60209081029190910101519050620002996002826200047d565b15620002dc576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b506001016200025e565b5060005b8151811015620003925760008282815181106200030b576200030b620009d0565b6020026020010151905060006001600160a01b0316816001600160a01b03160362000337575062000389565b620003446002826200049d565b1562000387576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101620002ea565b505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015620003e9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200040f9190620009e6565b6200041b919062000a16565b604080516001600160a01b038616602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663095ea7b360e01b179091529192506200047791869190620004b416565b50505050565b600062000494836001600160a01b03841662000585565b90505b92915050565b600062000494836001600160a01b03841662000689565b6040805180820190915260208082527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65649082015260009062000503906001600160a01b038516908490620006db565b80519091501562000392578080602001905181019062000524919062000a2c565b620003925760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000086565b600081815260018301602052604081205480156200067e576000620005ac60018362000a57565b8554909150600090620005c29060019062000a57565b90508181146200062e576000866000018281548110620005e657620005e6620009d0565b90600052602060002001549050808760000184815481106200060c576200060c620009d0565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062000642576200064262000a6d565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000497565b600091505062000497565b6000818152600183016020526040812054620006d25750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000497565b50600062000497565b6060620006ec8484600085620006f4565b949350505050565b606082471015620007575760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000086565b600080866001600160a01b0316858760405162000775919062000aa9565b60006040518083038185875af1925050503d8060008114620007b4576040519150601f19603f3d011682016040523d82523d6000602084013e620007b9565b606091505b509092509050620007cd87838387620007d8565b979650505050505050565b606083156200084c57825160000362000844576001600160a01b0385163b620008445760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000086565b5081620006ec565b620006ec8383815115620008635781518083602001fd5b8060405162461bcd60e51b815260040162000086919062000ac7565b6001600160a01b03811681146200089557600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b8051620008bb816200087f565b919050565b60008060008060808587031215620008d757600080fd5b8451620008e4816200087f565b602086810151919550906001600160401b03808211156200090457600080fd5b818801915088601f8301126200091957600080fd5b8151818111156200092e576200092e62000898565b8060051b604051601f19603f8301168101818110858211171562000956576200095662000898565b60405291825284820192508381018501918b8311156200097557600080fd5b938501935b828510156200099e576200098e85620008ae565b845293850193928501926200097a565b809850505050505050620009b560408601620008ae565b9150620009c560608601620008ae565b905092959194509250565b634e487b7160e01b600052603260045260246000fd5b600060208284031215620009f957600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b8082018082111562000497576200049762000a00565b60006020828403121562000a3f57600080fd5b8151801515811462000a5057600080fd5b9392505050565b8181038181111562000497576200049762000a00565b634e487b7160e01b600052603160045260246000fd5b60005b8381101562000aa057818101518382015260200162000a86565b50506000910152565b6000825162000abd81846020870162000a83565b9190910192915050565b602081526000825180602084015262000ae881604085016020870162000a83565b601f01601f19169190910160400192915050565b60805160a05160c0516136a762000b7960003960008181610487015281816115880152611f6e015260008181610461015281816113d0015261183c0152600081816102180152818161026d015281816106a2015281816107fc015281816117590152818161195401528181611f04015261215e01526136a76000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c80639a4575b9116100e3578063c75eea9c1161008c578063e0351e1311610066578063e0351e1314610485578063f2fde38b146104ab578063f6e2145e146104be57600080fd5b8063c75eea9c14610439578063cf7401f31461044c578063dc0bd9711461045f57600080fd5b8063b0f479a1116100bd578063b0f479a1146103f3578063c0d7865514610411578063c4bffe2b1461042457600080fd5b80639a4575b91461034f578063a7cd63b71461036f578063af58d59f1461038457600080fd5b80634059f55b1161014557806379ba50971161011f57806379ba5097146103165780638926f54f1461031e5780638da5cb5b1461033157600080fd5b80634059f55b146102aa57806354c8a4f3146102ee57806378a010b21461030357600080fd5b8063181f5a7711610176578063181f5a77146101da57806321df0da714610216578063240028e81461025d57600080fd5b806301ffc9a7146101925780630a2fd493146101ba575b600080fd5b6101a56101a0366004612901565b6104d1565b60405190151581526020015b60405180910390f35b6101cd6101c8366004612960565b6105b6565b6040516101b191906129df565b6101cd6040518060400160405280601f81526020017f4275726e46726f6d4d696e74546f6b656e506f6f6c20312e352e302d6465760081525081565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b1565b6101a561026b366004612a1f565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff90811691161490565b6102bd6102b8366004612a3c565b610666565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260209283015192810192909252016101b1565b6103016102fc366004612ac3565b61082c565b005b610301610311366004612b2f565b6108a7565b610301610a1b565b6101a561032c366004612960565b610b18565b60005473ffffffffffffffffffffffffffffffffffffffff16610238565b61036261035d366004612bb2565b610b2f565b6040516101b19190612bed565b610377610bd6565b6040516101b19190612c4d565b610397610392366004612960565b610be7565b6040516101b1919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60045473ffffffffffffffffffffffffffffffffffffffff16610238565b61030161041f366004612a1f565b610cbc565b61042c610d97565b6040516101b19190612ca7565b610397610447366004612960565b610e4f565b61030161045a366004612e03565b610f21565b7f0000000000000000000000000000000000000000000000000000000000000000610238565b7f00000000000000000000000000000000000000000000000000000000000000006101a5565b6103016104b9366004612a1f565b610f39565b6103016104cc366004612e48565b610f4d565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167faff2afbf00000000000000000000000000000000000000000000000000000000148061056457507fffffffff0000000000000000000000000000000000000000000000000000000082167f773a5d4500000000000000000000000000000000000000000000000000000000145b806105b057507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b67ffffffffffffffff811660009081526007602052604090206004018054606091906105e190612e8a565b80601f016020809104026020016040519081016040528092919081815260200182805461060d90612e8a565b801561065a5780601f1061062f5761010080835404028352916020019161065a565b820191906000526020600020905b81548152906001019060200180831161063d57829003601f168201915b50505050509050919050565b604080518082019091526000808252602082015261068b61068683612f88565b611392565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166340c10f196106d76060850160408601612a1f565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff909116600482015260608501356024820152604401600060405180830381600087803b15801561074757600080fd5b505af115801561075b573d6000803e3d6000fd5b50610770925050506060830160408401612a1f565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f084606001356040516107d291815260200190565b60405180910390a3506040805180820190915273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152606090910135602082015290565b610834611503565b6108a18484808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505060408051602080880282810182019093528782529093508792508691829185019084908082843760009201919091525061158692505050565b50505050565b6108af611503565b6108b883610b18565b6108ff576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b67ffffffffffffffff83166000908152600760205260408120600401805461092690612e8a565b80601f016020809104026020016040519081016040528092919081815260200182805461095290612e8a565b801561099f5780601f106109745761010080835404028352916020019161099f565b820191906000526020600020905b81548152906001019060200180831161098257829003601f168201915b5050505067ffffffffffffffff86166000908152600760205260409020919250506004016109ce8385836130bb565b508367ffffffffffffffff167fdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf828585604051610a0d939291906131d5565b60405180910390a250505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a9c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016108f6565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60006105b0600567ffffffffffffffff841661173c565b6040805180820190915260608082526020820152610b54610b4f83613239565b611757565b610b61826060013561191f565b6040516060830135815233907f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df79060200160405180910390a26040518060400160405280610bbb8460200160208101906101c89190612960565b81526040805160208181019092526000815291015292915050565b6060610be260026119c8565b905090565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845260028201546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff1615159482019490945260039091015480841660608301529190910490911660808201526105b0906119d5565b610cc4611503565b73ffffffffffffffffffffffffffffffffffffffff8116610d11576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684910160405180910390a15050565b60606000610da560056119c8565b90506000815167ffffffffffffffff811115610dc357610dc3612ce9565b604051908082528060200260200182016040528015610dec578160200160208202803683370190505b50905060005b8251811015610e4857828181518110610e0d57610e0d6132c0565b6020026020010151828281518110610e2757610e276132c0565b67ffffffffffffffff90921660209283029190910190910152600101610df2565b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff1615159482019490945260019091015480841660608301529190910490911660808201526105b0906119d5565b610f29611503565b610f34838383611a87565b505050565b610f41611503565b610f4a81611b71565b50565b610f55611503565b60005b81811015610f34576000838383818110610f7457610f746132c0565b9050602002810190610f8691906132ef565b610f8f9061332d565b9050610fa48160600151826020015115611c66565b610fb78160800151826020015115611c66565b806020015115611280578051610fd99060059067ffffffffffffffff16611da3565b61101e5780516040517f1d5ad3c500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016108f6565b80604001515160000361105d576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805161010081018252606083810180516020908101516fffffffffffffffffffffffffffffffff9081168486019081524263ffffffff9081166080808901829052865151151560a0808b01919091528751870151861660c08b015296518a0151851660e08a015292885288519586018952828a01805186015185168752868601919091528051511515868a015280518501518416868801525188015183168583015283870194855288880151878901908152895167ffffffffffffffff1660009081526007865289902088518051825482890151838e01519289167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316177001000000000000000000000000000000009188168202177fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff908116740100000000000000000000000000000000000000009415158502178655848d015194890151948a16948a168202949094176001860155995180516002860180549b8301519f830151918b169b9093169a909a179d9096168a029c909c1790911696151502959095179098559485015194015193811693169091029190911760038201559151909190600482019061123390826133b1565b50508151606083015160808401516040517f0f135cbb9afa12a8bf3bbd071c117bcca4ddeca6160ef7f33d012a81b9c0c4719450611273939291906134cb565b60405180910390a1611389565b80516112989060059067ffffffffffffffff16611daf565b6112dd5780516040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016108f6565b805167ffffffffffffffff16600090815260076020526040812080547fffffffffffffffffffffff0000000000000000000000000000000000000000009081168255600182018390556002820180549091169055600381018290559061134660048301826128b3565b5050805160405167ffffffffffffffff90911681527f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599169060200160405180910390a15b50600101610f58565b60208101516040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa15801561142c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611450919061354e565b15611487576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114948160200151611dbb565b6114a181602001516105b6565b80519060200120816080015180519060200120146114f15780608001516040517f24eb47e50000000000000000000000000000000000000000000000000000000081526004016108f691906129df565b610f4a81602001518260600151611ee1565b60005473ffffffffffffffffffffffffffffffffffffffff163314611584576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016108f6565b565b7f00000000000000000000000000000000000000000000000000000000000000006115dd576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b82518110156116735760008382815181106115fd576115fd6132c0565b6020026020010151905061161b816002611f2890919063ffffffff16565b1561166a5760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b506001016115e0565b5060005b8151811015610f34576000828281518110611694576116946132c0565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036116d85750611734565b6116e3600282611f4a565b156117325760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101611677565b600081815260018301602052604081205415155b9392505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff16146117fe5760808101516040517f961c9a4f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016108f6565b60208101516040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa158015611898573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118bc919061354e565b156118f3576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119008160400151611f6c565b61190d8160200151611ff0565b610f4a8160200151826060015161213e565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152306004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b1580156119ad57600080fd5b505af11580156119c1573d6000803e3d6000fd5b5050505050565b6060600061175083612182565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152611a6382606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff1642611a47919061359a565b85608001516fffffffffffffffffffffffffffffffff166121dd565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b611a9083610b18565b611ad2576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016108f6565b611add826000611c66565b67ffffffffffffffff83166000908152600760205260409020611b009083612207565b611b0b816000611c66565b67ffffffffffffffff83166000908152600760205260409020611b319060020182612207565b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b838383604051611b64939291906134cb565b60405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff821603611bf0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016108f6565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b815115611d315781602001516fffffffffffffffffffffffffffffffff1682604001516fffffffffffffffffffffffffffffffff16101580611cbc575060408201516fffffffffffffffffffffffffffffffff16155b15611cf557816040517f70505e560000000000000000000000000000000000000000000000000000000081526004016108f691906135ad565b8015611d2d576040517f433fc33d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60408201516fffffffffffffffffffffffffffffffff16151580611d6a575060208201516fffffffffffffffffffffffffffffffff1615155b15611d2d57816040517fd68af9cc0000000000000000000000000000000000000000000000000000000081526004016108f691906135ad565b600061175083836123a9565b600061175083836123f8565b611dc481610b18565b611e06576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016108f6565b600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925233602483015273ffffffffffffffffffffffffffffffffffffffff16906383826b2b90604401602060405180830381865afa158015611e85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea9919061354e565b610f4a576040517f728fe07b0000000000000000000000000000000000000000000000000000000081523360048201526024016108f6565b67ffffffffffffffff82166000908152600760205260409020611d2d90600201827f00000000000000000000000000000000000000000000000000000000000000006124eb565b60006117508373ffffffffffffffffffffffffffffffffffffffff84166123f8565b60006117508373ffffffffffffffffffffffffffffffffffffffff84166123a9565b7f00000000000000000000000000000000000000000000000000000000000000008015611fa15750611f9f60028261286e565b155b15610f4a576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016108f6565b611ff981610b18565b61203b576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016108f6565b600480546040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925273ffffffffffffffffffffffffffffffffffffffff169063a8d87a3b90602401602060405180830381865afa1580156120b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120d891906135e9565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f4a576040517f728fe07b0000000000000000000000000000000000000000000000000000000081523360048201526024016108f6565b67ffffffffffffffff82166000908152600760205260409020611d2d90827f00000000000000000000000000000000000000000000000000000000000000006124eb565b60608160000180548060200260200160405190810160405280929190818152602001828054801561065a57602002820191906000526020600020905b8154815260200190600101908083116121be5750505050509050919050565b60006121fc856121ed8486613606565b6121f7908761361d565b61289d565b90505b949350505050565b815460009061223090700100000000000000000000000000000000900463ffffffff164261359a565b905080156122d25760018301548354612278916fffffffffffffffffffffffffffffffff808216928116918591700100000000000000000000000000000000909104166121dd565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b602082015183546122f8916fffffffffffffffffffffffffffffffff908116911661289d565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990611b649084906135ad565b60008181526001830160205260408120546123f0575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556105b0565b5060006105b0565b600081815260018301602052604081205480156124e157600061241c60018361359a565b85549091506000906124309060019061359a565b9050818114612495576000866000018281548110612450576124506132c0565b9060005260206000200154905080876000018481548110612473576124736132c0565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806124a6576124a6613630565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506105b0565b60009150506105b0565b825474010000000000000000000000000000000000000000900460ff161580612512575081155b1561251c57505050565b825460018401546fffffffffffffffffffffffffffffffff8083169291169060009061256290700100000000000000000000000000000000900463ffffffff164261359a565b9050801561262257818311156125a4576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018601546125de9083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff166121dd565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b848210156126d95773ffffffffffffffffffffffffffffffffffffffff8416612681576040517ff94ebcd100000000000000000000000000000000000000000000000000000000815260048101839052602481018690526044016108f6565b6040517f1a76572a000000000000000000000000000000000000000000000000000000008152600481018390526024810186905273ffffffffffffffffffffffffffffffffffffffff851660448201526064016108f6565b848310156127ec5760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1690600090829061271d908261359a565b612727878a61359a565b612731919061361d565b61273b919061365f565b905073ffffffffffffffffffffffffffffffffffffffff8616612794576040517f15279c0800000000000000000000000000000000000000000000000000000000815260048101829052602481018690526044016108f6565b6040517fd0c8d23a000000000000000000000000000000000000000000000000000000008152600481018290526024810186905273ffffffffffffffffffffffffffffffffffffffff871660448201526064016108f6565b6127f6858461359a565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515611750565b60008183106128ac5781611750565b5090919050565b5080546128bf90612e8a565b6000825580601f106128cf575050565b601f016020900490600052602060002090810190610f4a91905b808211156128fd57600081556001016128e9565b5090565b60006020828403121561291357600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461175057600080fd5b803567ffffffffffffffff8116811461295b57600080fd5b919050565b60006020828403121561297257600080fd5b61175082612943565b6000815180845260005b818110156129a157602081850181015186830182015201612985565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000611750602083018461297b565b73ffffffffffffffffffffffffffffffffffffffff81168114610f4a57600080fd5b803561295b816129f2565b600060208284031215612a3157600080fd5b8135611750816129f2565b600060208284031215612a4e57600080fd5b813567ffffffffffffffff811115612a6557600080fd5b820160e0818503121561175057600080fd5b60008083601f840112612a8957600080fd5b50813567ffffffffffffffff811115612aa157600080fd5b6020830191508360208260051b8501011115612abc57600080fd5b9250929050565b60008060008060408587031215612ad957600080fd5b843567ffffffffffffffff80821115612af157600080fd5b612afd88838901612a77565b90965094506020870135915080821115612b1657600080fd5b50612b2387828801612a77565b95989497509550505050565b600080600060408486031215612b4457600080fd5b612b4d84612943565b9250602084013567ffffffffffffffff80821115612b6a57600080fd5b818601915086601f830112612b7e57600080fd5b813581811115612b8d57600080fd5b876020828501011115612b9f57600080fd5b6020830194508093505050509250925092565b600060208284031215612bc457600080fd5b813567ffffffffffffffff811115612bdb57600080fd5b820160a0818503121561175057600080fd5b602081526000825160406020840152612c09606084018261297b565b905060208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152612c44828261297b565b95945050505050565b6020808252825182820181905260009190848201906040850190845b81811015612c9b57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612c69565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015612c9b57835167ffffffffffffffff1683529284019291840191600101612cc3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715612d3b57612d3b612ce9565b60405290565b60405160a0810167ffffffffffffffff81118282101715612d3b57612d3b612ce9565b8015158114610f4a57600080fd5b80356fffffffffffffffffffffffffffffffff8116811461295b57600080fd5b600060608284031215612da457600080fd5b6040516060810181811067ffffffffffffffff82111715612dc757612dc7612ce9565b6040529050808235612dd881612d64565b8152612de660208401612d72565b6020820152612df760408401612d72565b60408201525092915050565b600080600060e08486031215612e1857600080fd5b612e2184612943565b9250612e308560208601612d92565b9150612e3f8560808601612d92565b90509250925092565b60008060208385031215612e5b57600080fd5b823567ffffffffffffffff811115612e7257600080fd5b612e7e85828601612a77565b90969095509350505050565b600181811c90821680612e9e57607f821691505b602082108103612ed7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600082601f830112612eee57600080fd5b813567ffffffffffffffff80821115612f0957612f09612ce9565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715612f4f57612f4f612ce9565b81604052838152866020858801011115612f6857600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060e08236031215612f9a57600080fd5b612fa2612d18565b823567ffffffffffffffff80821115612fba57600080fd5b612fc636838701612edd565b8352612fd460208601612943565b6020840152612fe560408601612a14565b604084015260608501356060840152608085013591508082111561300857600080fd5b61301436838701612edd565b608084015260a085013591508082111561302d57600080fd5b61303936838701612edd565b60a084015260c085013591508082111561305257600080fd5b5061305f36828601612edd565b60c08301525092915050565b601f821115610f34576000816000526020600020601f850160051c810160208610156130945750805b601f850160051c820191505b818110156130b3578281556001016130a0565b505050505050565b67ffffffffffffffff8311156130d3576130d3612ce9565b6130e7836130e18354612e8a565b8361306b565b6000601f84116001811461313957600085156131035750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556119c1565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b828110156131885786850135825560209485019460019092019101613168565b50868210156131c3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b6040815260006131e8604083018661297b565b82810360208401528381528385602083013760006020858301015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116820101915050949350505050565b600060a0823603121561324b57600080fd5b613253612d41565b823567ffffffffffffffff81111561326a57600080fd5b61327636828601612edd565b82525061328560208401612943565b60208201526040830135613298816129f2565b60408201526060838101359082015260808301356132b5816129f2565b608082015292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee183360301811261332357600080fd5b9190910192915050565b6000610120823603121561334057600080fd5b613348612d41565b61335183612943565b8152602083013561336181612d64565b6020820152604083013567ffffffffffffffff81111561338057600080fd5b61338c36828601612edd565b60408301525061339f3660608501612d92565b60608201526132b53660c08501612d92565b815167ffffffffffffffff8111156133cb576133cb612ce9565b6133df816133d98454612e8a565b8461306b565b602080601f83116001811461343257600084156133fc5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556130b3565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561347f57888601518255948401946001909101908401613460565b50858210156134bb57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b67ffffffffffffffff8416815260e0810161351760208301858051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b82511515608083015260208301516fffffffffffffffffffffffffffffffff90811660a084015260408401511660c08301526121ff565b60006020828403121561356057600080fd5b815161175081612d64565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156105b0576105b061356b565b606081016105b082848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b6000602082840312156135fb57600080fd5b8151611750816129f2565b80820281158282048414176105b0576105b061356b565b808201808211156105b0576105b061356b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600082613695577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c6343000818000a",
}

var BurnFromMintTokenPoolABI = BurnFromMintTokenPoolMetaData.ABI

var BurnFromMintTokenPoolBin = BurnFromMintTokenPoolMetaData.Bin

func DeployBurnFromMintTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnFromMintTokenPool, error) {
	parsed, err := BurnFromMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnFromMintTokenPoolBin), backend, token, allowlist, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnFromMintTokenPool{address: address, abi: *parsed, BurnFromMintTokenPoolCaller: BurnFromMintTokenPoolCaller{contract: contract}, BurnFromMintTokenPoolTransactor: BurnFromMintTokenPoolTransactor{contract: contract}, BurnFromMintTokenPoolFilterer: BurnFromMintTokenPoolFilterer{contract: contract}}, nil
}

type BurnFromMintTokenPool struct {
	address common.Address
	abi     abi.ABI
	BurnFromMintTokenPoolCaller
	BurnFromMintTokenPoolTransactor
	BurnFromMintTokenPoolFilterer
}

type BurnFromMintTokenPoolCaller struct {
	contract *bind.BoundContract
}

type BurnFromMintTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type BurnFromMintTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type BurnFromMintTokenPoolSession struct {
	Contract     *BurnFromMintTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnFromMintTokenPoolCallerSession struct {
	Contract *BurnFromMintTokenPoolCaller
	CallOpts bind.CallOpts
}

type BurnFromMintTokenPoolTransactorSession struct {
	Contract     *BurnFromMintTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type BurnFromMintTokenPoolRaw struct {
	Contract *BurnFromMintTokenPool
}

type BurnFromMintTokenPoolCallerRaw struct {
	Contract *BurnFromMintTokenPoolCaller
}

type BurnFromMintTokenPoolTransactorRaw struct {
	Contract *BurnFromMintTokenPoolTransactor
}

func NewBurnFromMintTokenPool(address common.Address, backend bind.ContractBackend) (*BurnFromMintTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(BurnFromMintTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnFromMintTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPool{address: address, abi: abi, BurnFromMintTokenPoolCaller: BurnFromMintTokenPoolCaller{contract: contract}, BurnFromMintTokenPoolTransactor: BurnFromMintTokenPoolTransactor{contract: contract}, BurnFromMintTokenPoolFilterer: BurnFromMintTokenPoolFilterer{contract: contract}}, nil
}

func NewBurnFromMintTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*BurnFromMintTokenPoolCaller, error) {
	contract, err := bindBurnFromMintTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolCaller{contract: contract}, nil
}

func NewBurnFromMintTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnFromMintTokenPoolTransactor, error) {
	contract, err := bindBurnFromMintTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolTransactor{contract: contract}, nil
}

func NewBurnFromMintTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnFromMintTokenPoolFilterer, error) {
	contract, err := bindBurnFromMintTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolFilterer{contract: contract}, nil
}

func bindBurnFromMintTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnFromMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnFromMintTokenPool.Contract.BurnFromMintTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.BurnFromMintTokenPoolTransactor.contract.Transfer(opts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.BurnFromMintTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnFromMintTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.contract.Transfer(opts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetAllowList(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetAllowList(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _BurnFromMintTokenPool.Contract.GetAllowListEnabled(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _BurnFromMintTokenPool.Contract.GetAllowListEnabled(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getRemotePool", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _BurnFromMintTokenPool.Contract.GetRemotePool(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _BurnFromMintTokenPool.Contract.GetRemotePool(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRmnProxy(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRmnProxy(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetRouter() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRouter(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRouter(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _BurnFromMintTokenPool.Contract.GetSupportedChains(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnFromMintTokenPool.Contract.GetSupportedChains(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetToken() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetToken(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetToken(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsSupportedChain(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsSupportedChain(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsSupportedToken(&_BurnFromMintTokenPool.CallOpts, token)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsSupportedToken(&_BurnFromMintTokenPool.CallOpts, token)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) Owner() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.Owner(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) Owner() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.Owner(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnFromMintTokenPool.Contract.SupportsInterface(&_BurnFromMintTokenPool.CallOpts, interfaceId)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnFromMintTokenPool.Contract.SupportsInterface(&_BurnFromMintTokenPool.CallOpts, interfaceId)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) TypeAndVersion() (string, error) {
	return _BurnFromMintTokenPool.Contract.TypeAndVersion(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _BurnFromMintTokenPool.Contract.TypeAndVersion(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.AcceptOwnership(&_BurnFromMintTokenPool.TransactOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.AcceptOwnership(&_BurnFromMintTokenPool.TransactOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnFromMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnFromMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "applyChainUpdates", chains)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyChainUpdates(&_BurnFromMintTokenPool.TransactOpts, chains)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyChainUpdates(&_BurnFromMintTokenPool.TransactOpts, chains)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.LockOrBurn(&_BurnFromMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.LockOrBurn(&_BurnFromMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ReleaseOrMint(&_BurnFromMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ReleaseOrMint(&_BurnFromMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetRemotePool(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetRemotePool(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetRouter(&_BurnFromMintTokenPool.TransactOpts, newRouter)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetRouter(&_BurnFromMintTokenPool.TransactOpts, newRouter)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.TransferOwnership(&_BurnFromMintTokenPool.TransactOpts, to)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.TransferOwnership(&_BurnFromMintTokenPool.TransactOpts, to)
}

type BurnFromMintTokenPoolAllowListAddIterator struct {
	Event *BurnFromMintTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolAllowListAdd)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolAllowListAdd)
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

func (it *BurnFromMintTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*BurnFromMintTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolAllowListAddIterator{contract: _BurnFromMintTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolAllowListAdd)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*BurnFromMintTokenPoolAllowListAdd, error) {
	event := new(BurnFromMintTokenPoolAllowListAdd)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolAllowListRemoveIterator struct {
	Event *BurnFromMintTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolAllowListRemove)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolAllowListRemove)
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

func (it *BurnFromMintTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*BurnFromMintTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolAllowListRemoveIterator{contract: _BurnFromMintTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolAllowListRemove)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*BurnFromMintTokenPoolAllowListRemove, error) {
	event := new(BurnFromMintTokenPoolAllowListRemove)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolBurnedIterator struct {
	Event *BurnFromMintTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolBurned)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolBurned)
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

func (it *BurnFromMintTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*BurnFromMintTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolBurnedIterator{contract: _BurnFromMintTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolBurned)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseBurned(log types.Log) (*BurnFromMintTokenPoolBurned, error) {
	event := new(BurnFromMintTokenPoolBurned)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolChainAddedIterator struct {
	Event *BurnFromMintTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolChainAdded)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolChainAdded)
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

func (it *BurnFromMintTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainAddedIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolChainAddedIterator{contract: _BurnFromMintTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolChainAdded)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseChainAdded(log types.Log) (*BurnFromMintTokenPoolChainAdded, error) {
	event := new(BurnFromMintTokenPoolChainAdded)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolChainConfiguredIterator struct {
	Event *BurnFromMintTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolChainConfigured)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolChainConfigured)
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

func (it *BurnFromMintTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolChainConfiguredIterator{contract: _BurnFromMintTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolChainConfigured)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseChainConfigured(log types.Log) (*BurnFromMintTokenPoolChainConfigured, error) {
	event := new(BurnFromMintTokenPoolChainConfigured)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolChainRemovedIterator struct {
	Event *BurnFromMintTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolChainRemoved)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolChainRemoved)
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

func (it *BurnFromMintTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolChainRemovedIterator{contract: _BurnFromMintTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolChainRemoved)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseChainRemoved(log types.Log) (*BurnFromMintTokenPoolChainRemoved, error) {
	event := new(BurnFromMintTokenPoolChainRemoved)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolConfigChangedIterator struct {
	Event *BurnFromMintTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolConfigChanged)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolConfigChanged)
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

func (it *BurnFromMintTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*BurnFromMintTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolConfigChangedIterator{contract: _BurnFromMintTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolConfigChanged)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseConfigChanged(log types.Log) (*BurnFromMintTokenPoolConfigChanged, error) {
	event := new(BurnFromMintTokenPoolConfigChanged)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolLockedIterator struct {
	Event *BurnFromMintTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolLocked)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolLocked)
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

func (it *BurnFromMintTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BurnFromMintTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolLockedIterator{contract: _BurnFromMintTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolLocked)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseLocked(log types.Log) (*BurnFromMintTokenPoolLocked, error) {
	event := new(BurnFromMintTokenPoolLocked)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolMintedIterator struct {
	Event *BurnFromMintTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolMinted)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolMinted)
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

func (it *BurnFromMintTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnFromMintTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolMintedIterator{contract: _BurnFromMintTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolMinted)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseMinted(log types.Log) (*BurnFromMintTokenPoolMinted, error) {
	event := new(BurnFromMintTokenPoolMinted)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolOwnershipTransferRequestedIterator struct {
	Event *BurnFromMintTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolOwnershipTransferRequested)
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

func (it *BurnFromMintTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnFromMintTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolOwnershipTransferRequestedIterator{contract: _BurnFromMintTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolOwnershipTransferRequested)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnFromMintTokenPoolOwnershipTransferRequested, error) {
	event := new(BurnFromMintTokenPoolOwnershipTransferRequested)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolOwnershipTransferredIterator struct {
	Event *BurnFromMintTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolOwnershipTransferred)
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

func (it *BurnFromMintTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnFromMintTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolOwnershipTransferredIterator{contract: _BurnFromMintTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolOwnershipTransferred)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*BurnFromMintTokenPoolOwnershipTransferred, error) {
	event := new(BurnFromMintTokenPoolOwnershipTransferred)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolReleasedIterator struct {
	Event *BurnFromMintTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolReleased)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolReleased)
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

func (it *BurnFromMintTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnFromMintTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolReleasedIterator{contract: _BurnFromMintTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolReleased)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseReleased(log types.Log) (*BurnFromMintTokenPoolReleased, error) {
	event := new(BurnFromMintTokenPoolReleased)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolRemotePoolSetIterator struct {
	Event *BurnFromMintTokenPoolRemotePoolSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolRemotePoolSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolRemotePoolSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolRemotePoolSet)
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

func (it *BurnFromMintTokenPoolRemotePoolSetIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolRemotePoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolRemotePoolSet struct {
	RemoteChainSelector uint64
	PreviousPoolAddress []byte
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolRemotePoolSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolRemotePoolSetIterator{contract: _BurnFromMintTokenPool.contract, event: "RemotePoolSet", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolRemotePoolSet)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseRemotePoolSet(log types.Log) (*BurnFromMintTokenPoolRemotePoolSet, error) {
	event := new(BurnFromMintTokenPoolRemotePoolSet)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolRouterUpdatedIterator struct {
	Event *BurnFromMintTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolRouterUpdated)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolRouterUpdated)
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

func (it *BurnFromMintTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*BurnFromMintTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolRouterUpdatedIterator{contract: _BurnFromMintTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolRouterUpdated)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*BurnFromMintTokenPoolRouterUpdated, error) {
	event := new(BurnFromMintTokenPoolRouterUpdated)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolTokensConsumedIterator struct {
	Event *BurnFromMintTokenPoolTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolTokensConsumed)
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

	select {
	case log := <-it.logs:
		it.Event = new(BurnFromMintTokenPoolTokensConsumed)
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

func (it *BurnFromMintTokenPoolTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*BurnFromMintTokenPoolTokensConsumedIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolTokensConsumedIterator{contract: _BurnFromMintTokenPool.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolTokensConsumed)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseTokensConsumed(log types.Log) (*BurnFromMintTokenPoolTokensConsumed, error) {
	event := new(BurnFromMintTokenPoolTokensConsumed)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnFromMintTokenPool.abi.Events["AllowListAdd"].ID:
		return _BurnFromMintTokenPool.ParseAllowListAdd(log)
	case _BurnFromMintTokenPool.abi.Events["AllowListRemove"].ID:
		return _BurnFromMintTokenPool.ParseAllowListRemove(log)
	case _BurnFromMintTokenPool.abi.Events["Burned"].ID:
		return _BurnFromMintTokenPool.ParseBurned(log)
	case _BurnFromMintTokenPool.abi.Events["ChainAdded"].ID:
		return _BurnFromMintTokenPool.ParseChainAdded(log)
	case _BurnFromMintTokenPool.abi.Events["ChainConfigured"].ID:
		return _BurnFromMintTokenPool.ParseChainConfigured(log)
	case _BurnFromMintTokenPool.abi.Events["ChainRemoved"].ID:
		return _BurnFromMintTokenPool.ParseChainRemoved(log)
	case _BurnFromMintTokenPool.abi.Events["ConfigChanged"].ID:
		return _BurnFromMintTokenPool.ParseConfigChanged(log)
	case _BurnFromMintTokenPool.abi.Events["Locked"].ID:
		return _BurnFromMintTokenPool.ParseLocked(log)
	case _BurnFromMintTokenPool.abi.Events["Minted"].ID:
		return _BurnFromMintTokenPool.ParseMinted(log)
	case _BurnFromMintTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _BurnFromMintTokenPool.ParseOwnershipTransferRequested(log)
	case _BurnFromMintTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _BurnFromMintTokenPool.ParseOwnershipTransferred(log)
	case _BurnFromMintTokenPool.abi.Events["Released"].ID:
		return _BurnFromMintTokenPool.ParseReleased(log)
	case _BurnFromMintTokenPool.abi.Events["RemotePoolSet"].ID:
		return _BurnFromMintTokenPool.ParseRemotePoolSet(log)
	case _BurnFromMintTokenPool.abi.Events["RouterUpdated"].ID:
		return _BurnFromMintTokenPool.ParseRouterUpdated(log)
	case _BurnFromMintTokenPool.abi.Events["TokensConsumed"].ID:
		return _BurnFromMintTokenPool.ParseTokensConsumed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnFromMintTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnFromMintTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnFromMintTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (BurnFromMintTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x0f135cbb9afa12a8bf3bbd071c117bcca4ddeca6160ef7f33d012a81b9c0c471")
}

func (BurnFromMintTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (BurnFromMintTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnFromMintTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnFromMintTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (BurnFromMintTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (BurnFromMintTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnFromMintTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnFromMintTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (BurnFromMintTokenPoolRemotePoolSet) Topic() common.Hash {
	return common.HexToHash("0xdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf")
}

func (BurnFromMintTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (BurnFromMintTokenPoolTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPool) Address() common.Address {
	return _BurnFromMintTokenPool.address
}

type BurnFromMintTokenPoolInterface interface {
	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, chains []TokenPoolChainUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnFromMintTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnFromMintTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnFromMintTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnFromMintTokenPoolAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*BurnFromMintTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*BurnFromMintTokenPoolBurned, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnFromMintTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*BurnFromMintTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnFromMintTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnFromMintTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnFromMintTokenPoolConfigChanged, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BurnFromMintTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*BurnFromMintTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnFromMintTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*BurnFromMintTokenPoolMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnFromMintTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnFromMintTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnFromMintTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnFromMintTokenPoolOwnershipTransferred, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnFromMintTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*BurnFromMintTokenPoolReleased, error)

	FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolRemotePoolSetIterator, error)

	WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolSet(log types.Log) (*BurnFromMintTokenPoolRemotePoolSet, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*BurnFromMintTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*BurnFromMintTokenPoolRouterUpdated, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*BurnFromMintTokenPoolTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*BurnFromMintTokenPoolTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
