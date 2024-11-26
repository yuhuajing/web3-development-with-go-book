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
### 读取特定区块中账户余额-BlockHeight
```go
	blockNum := big.NewInt(99999)
	balance, err := client.BalanceAt(ctx, account, blockNum) //nil is the latest block
```
### 读取特定区块中账户余额-BlockHash
```go
	blockHash := common.HexToHash("0x0fa8fe23357be11db6273d5744a091b7f5baa70d7824addd680c8ed1fd2fbf0b")
	balance, err := client.BalanceAtHash(ctx, account, blockHash)
}
```
### 读取 Pending 池中账户余额
账户提交的待处理交易进入 `pending` 池等待校验打包
- 账户的余额随着构建交易会付出 `gas` 或者转账
- 在 `pending` 此种获取账户最新余额，有可能比直接读取余额的值要小
  - 取决于 `RPC` 的速度，如果 `RPC` 在网络中已经接收到了该交易，则返回执行该交易后的剩余余额，否则返回正常余额
```go
    balance, err := client.PendingBalanceAt(ctx, account)
```
以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在 `ETH` 中它是 `wei`。

要读取 `ETH` 值，您必须做计算 `wei/10^18`
### WEI/Eth converter
```go
func calcuBalanceToEth(bal *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(bal.String())
	fbalance = fbalance.Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return fbalance
}
```
