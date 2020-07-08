# defi-sdk-go
This library is a simple Go wrapper for DeFi SDK.

Visit [docs.zerion.io](https://docs.zerion.io) for full documentation. 

#  Install 
```bash
$ go get github.com/zeriontech/defi-sdk-go
```
# Usage

## Initialize DeFiSDK
DeFi SDK directly connects to the Ethereum blockchain. You are welcome to use an Ethereum node of your choice to start using DeFi SDK. 
If you don't have a node, in the example below you can use a node provided by Cloudflare and served through our domain. 

```golang
package main

import defisdk "github.com/zeriontech/defi-sdk-go"

func main() {
	ethereumNodeUrl := "https://eth-mainnet.zerion.io/"
	sdk := defisdk.New(ethereumNodeUrl)
}
```
## Examples
### Get supported protocols
```go
sdk.GetSupportedProtocols()
// [PieDAO Multi-Collateral Dai Bancor DeFi Money Market TokenSets 0x Staking Uniswap V1 Synthetix PoolTogether Dai Savings Rate Chai iearn.finance (v3) iearn.finance (v2) Idle dYdX Curve Compound Balancer Aave]
```
### Get account balance locked in a protocol
```go
userAddress := "0xa218a8346454c982912cf6d14c714663c2d510d8"
protocol := "Aave"

sdk.GetProtocolBalance(userAddress, protocol)
// [{
// Metadata:{
//     Name:Aave 
//     Description:Decentralized lending & borrowing protocol 
//     WebsiteURL:aave.com 
//     IconURL:protocol-icons.s3.amazonaws.com/aave.png 
//     Version:+0
// } 
// AdapterBalances:[{
//     Metadata:{
//         AdapterAddress:[139 98 192 32 145 254 6 174 52 84 211 193 41 33 179 38 17 186 85 1] 
//         AdapterType:Asset
//     } 
//     Balances:[{
//         Base:{
//             Metadata:{
//                 Token:[252 30 105 15 97 239 217 97 41 75 62 28 227 49 63 189 138 164 248 93] 
//                 Name:Aave Interest bearing DAI 
//                 Symbol:aDAI 
//                 Decimals:18
//             } 
//             Amount:+1008383439792242970
//         } 
//         Underlying:[{
//             Metadata:{
//                 Token:[107 23 84 116 232 144 148 196 77 169 139 149 78 237 234 196 149 39 29 15] 
//                 Name:Dai Stablecoin 
//                 Symbol:DAI 
//                 Decimals:18
//             } 
//             Amount:+1008383439792242970
//         }]
//     }]
// }]
// }]
```

### Get supported token types
```go
sdk.GetTokenAdapterNames()
// [PieDAO Pie Token SmartToken MToken SetToken Uniswap V1 pool token PoolTogether pool Chai token YToken IdleToken Curve pool token CToken Balancer pool token AToken ERC20]
```

### Get derivative token underlying components
```go
uniswapDaiPool := "0x2a1530c4c41db0b0b2bb646cb5eb1a67b7158667"
sdk.GetTokenComponents("Uniswap V1 pool token", uniswapDaiPool)

// [{
// Metadata:{
//     Token:[238 238 238 238 238 238 238 238 238 238 238 238 238 238 238 238 238 238 238 238] 
//     Name:Ether 
//     Symbol:ETH 
//     Decimals:18
// } 
// Amount:+1108093230485279057
// } 
// {
// Metadata:{
//     Token:[107 23 84 116 232 144 148 196 77 169 139 149 78 237 234 196 149 39 29 15] 
//     Name:Dai Stablecoin 
//     Symbol:DAI 
//     Decimals:18
// } 
// Amount:+221366932637588225901
// }]
```
