# Estimate
模拟当前交易执行需要的 `gas` 花销，该交易不会发送到链上执行，只会在 `rpc` 节点本地模拟执行
- `From`: 交易的构造地址
- `To`: 交易的接收方。`EOA` 或者合约地址
- `Gas`:`0` 表示在模拟执行过程中，`gas` 不限量，用来模拟出最终的花销
- `Value`:交易附加的 `Value`,转账需要花费额外的 `gas`，所以需要表明
- `Data`:交易附加的 `data`,`data` 数据上链或者发送到合约处理，都需要额外的 `gas`

```go
// CallMsg contains parameters for contract calls.
type CallMsg struct {
	From      common.Address  // the sender of the 'transaction'
	To        *common.Address // the destination contract (nil for contract creation)
	Gas       uint64          // if 0, the call executes with near-infinite gas
	GasPrice  *big.Int        // wei <-> gas exchange ratio
	GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
	GasTipCap *big.Int        // EIP-1559 tip per gas.
	Value     *big.Int        // amount of wei sent along with the call
	Data      []byte          // input data, usually an ABI-encoded contract method invocation

	AccessList types.AccessList // EIP-2930 access list.

	// For BlobTxType
	BlobGasFeeCap *big.Int
	BlobHashes    []common.Hash
}
```
`RPC` 节点在本地基于当前 `Pending` 的链状态模拟执行该交易，但是不能保证和真实链上 `gas` 消耗完全一致，仅能作为参考

因为，在模拟执行和真实上链的空隙，有可能存在其他交易更新了和当前交易相关地址的数据，造成 `gas` 消耗的增加或减少
## 代码
```go

func EstimateGas(from, to string, data []interface{}, value uint64) uint64 {
	var ctx = context.Background()
	var err error
	var (
		fromAddr  = common.HexToAddress(from)     // Convert the from address from hex to an Ethereum address.
		toAddr    = common.HexToAddress(to)       // Convert the to address from hex to an Ethereum address.
		amount    = new(big.Int).SetUint64(value) // Convert the value from uint64 to *big.Int.
		bytesData []byte
	)

	// Encode the data if it's not already hex-encoded.
	bytesData = encodeData(data)

	// Create a message which contains information about the transaction.
	msg := ethereum.CallMsg{
		From:  fromAddr,
		To:    &toAddr,
		Gas:   0x00,
		Value: amount,
		Data:  bytesData,
	}

	// Estimate the gas required for the transaction.
	gas, err := client.EstimateGas(ctx, msg)
	if err != nil {
		log.Fatalln(err)
	}

	return gas
}
```
