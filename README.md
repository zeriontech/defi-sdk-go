# defi-sdk-go
This library is a simple Go wrapper for DeFi SDK.

#  Install 
```bash
$ go get github.com/zeriontech/defi-sdk-go
```
# Usage

### Initialize DeFiSDK
DeFi SDK directly connects to the Ethereum blockchain. You are welcome to use an Ethereum node of your choice to start using DeFi SDK. 
If you don't have a node, in the example below you can use a node provided by Cloudflare and served through our domain. 

```golang
import defisdk "github.com/zeriontech/defi-sdk-go"

ethereumNodeUrl := "https://eth-mainnet.zerion.io/"
defisdk := defisdk.New(ethereumNodeUrl)
```
### Get supported protocols
```go
defisdk.GetSupportedProtocols()
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
