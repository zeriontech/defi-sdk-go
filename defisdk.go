package defisdk

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	registry "github.com/zeriontech/defi-sdk-go/adapter-registry"
)

type DefiSDK struct {
	registryAddress common.Address
	client          *ethclient.Client
	registry        *registry.AdapterRegistry
}

const registryAddressString string = "0x06fe76b2f432fdfecaef1a7d4f6c3d41b5861672"

func New(nodeURL string) DefiSDK {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}
	registryAddress := common.HexToAddress(registryAddressString)
	instance, err := registry.NewAdapterRegistry(registryAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	sdk := DefiSDK{registryAddress, client, instance}
	return sdk
}

func (sdk DefiSDK) GetSupportedProtocols() []string {
	supportedProtocols, err := sdk.registry.GetProtocolNames(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return supportedProtocols
}

func (sdk DefiSDK) GetTokenAdapterNames() []string {
	tokenAdapters, err := sdk.registry.GetTokenAdapterNames(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return tokenAdapters
}

func (sdk DefiSDK) GetBalances(userAddress string) []registry.ProtocolBalance {
	address := common.HexToAddress(userAddress)

	protocolBalances, err := sdk.registry.GetBalances(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	return protocolBalances
}

func (sdk DefiSDK) GetProtocolBalance(userAddress string, protocolName string) registry.ProtocolBalance {
	address := common.HexToAddress(userAddress)
	protocolNames := []string{protocolName}

	protocolBalances, err := sdk.registry.GetProtocolBalances(&bind.CallOpts{}, address, protocolNames)
	if err != nil {
		log.Fatal(err)
	}
	return protocolBalances[0]
}

func (sdk DefiSDK) GetTokenComponents(tokenType string, tokenAddress string) []registry.TokenBalance {
	tokenBalance, err := sdk.registry.GetFullTokenBalance(&bind.CallOpts{}, tokenType, common.HexToAddress(tokenAddress))
	if err != nil {
		log.Fatal(err)
	}
	return tokenBalance.Underlying
}
