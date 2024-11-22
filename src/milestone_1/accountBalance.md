# 账户信息
## 读取账户余额
`eth_getBalance` 支持读取传参地址在特定区块高度、特定区块hash、最新区块中、Pending池的余额

### 读取最新区块高度的账户余额
```go
func Balance(account common.Address) (*big.Int, error) {
	var ctx = context.Background()
	balance, err := client.BalanceAt(ctx, account, nil) //nil is the latest block
	if err != nil {
		return balance, err
	}
	return balance, nil
}
```
### 基于区块高度 99999 读取账户余额
```go
	blockNum := big.NewInt(99999)
	balance, err := client.BalanceAt(ctx, account, blockNum) //nil is the latest block
```
### 基于区块hash 读取账户余额
```go
	blockHash := common.HexToHash("0x0fa8fe23357be11db6273d5744a091b7f5baa70d7824addd680c8ed1fd2fbf0b")
	balance, err := client.BalanceAtHash(ctx, account, blockHash)
}
```
### 账户提交的待处理交易进入 pending 池等待校验打包，获取账户在 pending 池中的最新余额
```go
    balance, err := client.PendingBalanceAt(ctx, account)
```
### WEI/Eth converter
```go
func calcuBalanceToEth(bal *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(bal.String())
	fbalance = fbalance.Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return fbalance
}
```
