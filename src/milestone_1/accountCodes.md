# 账户信息
## 读取账户数据
合约地址和EOA地址的区别在于合约地址在账户中存储合约代码，通过判断地址数据长度可以判断当前地址是否是合约地址

`eth_getCode` 支持读取传参地址在特定区块高度、特定区块hash、最新区块中、Pending池中的账户 contract codes

### 读取最新区块高度的账户contract codes
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

### 基于区块高度 99999 读取账户codes
```go
	blockNum := big.NewInt(99999)
	
	bytecode, err := client.CodeAt(context.Background(), addr, blockNum) //nil is the latest block
```
### 基于区块hash 读取账户codes
```go
	blockHash := common.HexToHash("0x0fa8fe23357be11db6273d5744a091b7f5baa70d7824addd680c8ed1fd2fbf0b")
    
	bytecode, err := client.CodeAtHash(ctx, account, blockHash)
}
```
### 账户提交的待处理交易进入 pending 池等待校验打包，获取账户在 pending 池中的最新codes
```go
    bytecode, err := client.PendingCodeAt(ctx, account)
```
