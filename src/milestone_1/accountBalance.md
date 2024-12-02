# 账户信息
## 读取账户余额
`eth_getBalance` 支持读取传参地址在特定区块高度、特定区块hash、最新区块中、Pending池的余额
### 读取最新区块高度的账户余额
```go
	balance, err := client.BalanceAt(ctx, account, nil) //nil is the latest block
```
### 读取特定区块中账户余额-BlockHeight
```go
	blockNum := big.NewInt(99999)
	balance, err := client.BalanceAt(ctx, account, blockNum)
```
区块 `num` 的相关异常情况：
```go
func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
// Sign returns:
//   - -1 if x < 0;
//   - 0 if x == 0;
//   - +1 if x > 0.
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	// It's negative.
	// IsInt64 reports whether x can be represented as an int64.
	if number.IsInt64() {
		return rpc.BlockNumber(number.Int64()).String()
	}
	// It's negative and large, which is invalid.
	return fmt.Sprintf("<invalid %d>", number)
}
```
- `blockNum == nil`， 表示基于最新区块高度的合约状态读取 `slot` 数值
- `blockNum` 为负数:
  - 数值 `-1 ~-4` 都有具体类型的对应
  - 数值 `<-4`,报错 `invalid argument 1: hex string without 0x prefix`
```go
func (bn BlockNumber) String() string {
	switch bn {
	case EarliestBlockNumber: //0
		return "earliest"
	case LatestBlockNumber://-2
		return "latest"
	case PendingBlockNumber://-1
		return "pending"
	case FinalizedBlockNumber://-3
		return "finalized"
	case SafeBlockNumber://-4
		return "safe"
	default:
		if bn < 0 {
			return fmt.Sprintf("<invalid %d>", bn)
		}
		return hexutil.Uint64(bn).String()
	}
}
```
- blockNum 为正数: 读取截止当前区块的合约内部存储的 slot 数值
- 提供的区块高度 `> latestBlockHeight`，报错 `ethereum.NotFound`

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

## 完整代码
```go
package milestone1

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
)

func BalanceFromLatestBlock(account common.Address, ctx context.Context) (*big.Int, error) {
	balance, err := client.BalanceAt(ctx, account, nil) //nil is the latest block
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account balance err = %v", err)))
	}
	return balance, nil
}

func BalanceFromBlock(account common.Address, number *big.Int, ctx context.Context) (*big.Int, error) {
	balance, err := client.BalanceAt(ctx, account, number)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account balance err = %v", err)))
	}
	return balance, nil
}

func BalanceFromBlockHash(account common.Address, hash common.Hash, ctx context.Context) (*big.Int, error) {
	balance, err := client.BalanceAtHash(ctx, account, hash)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account balance err = %v", err)))
	}
	return balance, nil
}

func BalanceFromPendingPool(account common.Address, ctx context.Context) (*big.Int, error) {
	balance, err := client.PendingBalanceAt(ctx, account)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get account balance err = %v", err)))
	}
	return balance, nil
}

func calcuBalanceToEth(bal *big.Int, decimal int) *big.Float {
	balance := new(big.Float)
	balance.SetString(bal.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(decimal)))
	return balance
}
```
