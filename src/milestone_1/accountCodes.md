# 账户信息
## 读取账户数据
- `EOA`：从私钥导出的账户地址，允许直接构建交易操作账户余额
- `Smart Contract`：`EOA` 创建的合约地址，账户下存储特定的合约逻辑
    - 合约账户不允许直接构造交易
    - [EIP4337](https://docs.stackup.sh/docs/account-abstraction)介绍如何实现合约账户的交易代付
- 合约地址和EOA地址的区别在于合约地址在账户中存储合约代码，通过判断地址数据长度可以判断当前地址是否是合约地址
- `eth_getCode` 支持读取传参地址在特定区块高度、特定区块 `hash`、最新区块中、`Pending` 池中的账户 `contract codes`
### 读取最新区块高度的账户contract codes
将区块号设置为 `nil` 将使用最新的区块高度
1. 先使用简单的正则表达式来检查以太坊地址是否有效
2. 获取地址存储的代码
    1. 如果长度为空，表示目前该地址时 EOA 地址
    2. 如果长度不为空，表明该地址是合约地址
3. [Solidity判断](https://yuhuajing.github.io/solidity-book/milestone_3/contracts-getcodes.html)

```go
// check the address whether it is a valid  address
func validAddress(addr common.Address) bool {
	// 16 hex 0-f
	re := regexp.MustCompile("0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr.Hex())
}

// check the address whether is a smart contract address
func checkContractAddress(addr common.Address) bool {
	if !validAddress(addr) {
		return false
	}
	bytecode, err := client.CodeAt(context.Background(), addr, nil) //nil is the latest block
	if err != nil {
		panic(err)
	}
	isContract := len(bytecode) > 0
	if isContract {
		//fmt.Println("SC address")
		return true
	}
	fmt.Println("This is normal address, but we want a smart contract address")
	return false
}
```

### 读取特定区块中的账户codes-BlockHash
```go
	blockNum := big.NewInt(99999)
	
	bytecode, err := client.CodeAt(context.Background(), addr, blockNum) //nil is the latest block
```
### 读取特定区块中的账户codes-BlockHeight
```go
	blockHash := common.HexToHash("0x0fa8fe23357be11db6273d5744a091b7f5baa70d7824addd680c8ed1fd2fbf0b")
    
	bytecode, err := client.CodeAtHash(ctx, account, blockHash)
}
```
### 获取账户在 pending 池中的codes
合约部署交易已经提交，但是还没有得到确定

此时，如果 RPC 收到该交易，验证执行后返沪i该地址的 codes
```go
    bytecode, err := client.PendingCodeAt(ctx, account)
```
