package defisdk

import "testing"

func TestMainFunctions(t *testing.T) {

	sdk := New("https://eth-mainnet.zerion.io/")
	sdk.GetTokenAdapterNames()

	user := "0xa218a8346454c982912cf6d14c714663c2d510d8"
	sdk.GetBalances(user)

	protocol := "Aave"
	sdk.GetProtocolBalance(user, protocol)

	uniswapDaiPool := "0x2a1530c4c41db0b0b2bb646cb5eb1a67b7158667"
	sdk.GetTokenComponents("Uniswap V1 pool token", uniswapDaiPool)
}
