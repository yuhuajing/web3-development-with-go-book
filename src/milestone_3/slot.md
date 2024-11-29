# GetSlotData
合约数据全部按照[Solidity Slot 存储规则](https://yuhuajing.github.io/solidity-book/milestone_1/static-slot-storage.html)存储在区块链上，因此只要上链的数据就能通过 slot 键获取值
## GetSlocByKey
合约数据按照声明顺序和编码规则存储在链上空间
### 读取区块内slot值
```go
func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	// It's negative.
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
  - 数值 `<-4`,报错
```go
const (
SafeBlockNumber      = BlockNumber(-4)
FinalizedBlockNumber = BlockNumber(-3)
LatestBlockNumber    = BlockNumber(-2)
PendingBlockNumber   = BlockNumber(-1)
EarliestBlockNumber  = BlockNumber(0)
)
```
- blockNum 为正数: 读取截止当前区块的合约内部存储的 slot 数值
- blockNum > 最新区快： 报错 ` error = header not found`

```go
func GetStorageAtBlock(ctx context.Context, address common.Address, slot common.Hash, blockNum *big.Int) (*big.Int, error) {
	//t := common.BigToHash(big.NewInt(int64(slot)))
	int256 := new(big.Int)
	res, err := client.StorageAt(ctx, address, slot, blockNum) // nil is the latest blockNum
	if err != nil {
		return int256, err
	}
	int256.SetBytes(res)
	return int256, nil
}
```
### 读取区块slot值
基于区块 hash 锁定区块，读取截止区块高度的合约数据的 slot 数值
- hash 不存在：报错 `error = header for hash not found`
```go
func GetStorageAtHash(ctx context.Context, address common.Address, slot common.Hash, hash common.Hash) (*big.Int, error) {
	int256 := new(big.Int)
	res, err := client.StorageAtHash(ctx, address, slot, hash)
	if err != nil {
		return int256, err
	}
	int256.SetBytes(res)

	return int256, nil
}
```
### PendingStorage
```go

func GetPendingStorage(ctx context.Context, address common.Address, slot common.Hash) (*big.Int, error) {
	int256 := new(big.Int)
	res, err := client.PendingStorageAt(ctx, address, slot)
	if err != nil {
		return int256, err
	}
	int256.SetBytes(res)

	return int256, nil
}
```
## ContractSlotParser
合约内部的存储结构通过标准的 `json` 请求可以获取 
> `solc --storage-layout --pretty-json -o $PWD/tempDirForSolc --overwrite ./xxx.sol`

`Json` 对象包含两个键值： `storage` 和 `types`
### Storage
```json
{
    "astId": 2,
    "contract": "fileA:A",
    "label": "x",
    "offset": 0,
    "slot": "0",
    "type": "t_uint256"
}
```
- `astId`:状态变量声明的 `AST` 节点的 `ID`
- `contract`: 当前合约名称
- `label`:状态变量的名称
- `offset`:字节偏移量，表示在当前 `slot` 中的偏移量
- `slot`:存储的插槽位置
- `type`:标识符，表示具体的数据存储，在 `types` 中存在对用的结构体数据
### Type
```json
{
    "base": "t_bool",
    "encoding": "inplace",
    "label": "uint256",
    "numberOfBytes": "32",
    "key": "t_string_memory_ptr",
    "value": "t_uint256",
    "members": `Type` 数组
}
```
基础数据结构：
- `encoding`：数据编码方式
  - `inplace`:数据能够在插槽中连续存储的数据
  - `mapping`:基于 `Keccak-256` 寻址
  - `dynamic_array`:基于 `Keccak-256` 寻址
  - `bytes`:单槽或基于 `Keccak-256` 哈希值，取决于数据大小
- `label`:类型名称
- `numberOfBytes`: 数据存储占据的字节数，如果 大于 `32`，表示使用一个以上的插槽存储数据

其中， `mapping` 类型额外包含
- `key`: 键值类型
- `value`:值类型
```json
    "t_mapping(t_uint256,t_mapping(t_address,t_uint256))": {
      "encoding": "mapping",
      "key": "t_uint256",
      "label": "mapping(uint256 => mapping(address => uint256))",
      "numberOfBytes": "32",
      "value": "t_mapping(t_address,t_uint256)"
    },
```

数组包含：
- `base`: 数组成员的数据类型
```json
    "t_array(t_bool)5_storage": {
      "base": "t_bool",
      "encoding": "inplace",
      "label": "bool[5]",
      "numberOfBytes": "32"
    },
```

结构体包含:
- `members`: 数组类型，表示结构体内部每个值的类型
```json
    "t_struct(Entity)62_storage": {
      "encoding": "inplace",
      "label": "struct StorageScan.Entity",
      "members": [
        {
          "astId": 57,
          "contract": "StorageScan.sol:StorageScan",
          "label": "age",
          "offset": 0,
          "slot": "0",
          "type": "t_uint64"
        },
        {
          "astId": 59,
          "contract": "StorageScan.sol:StorageScan",
          "label": "id",
          "offset": 8,
          "slot": "0",
          "type": "t_uint128"
        },
        {
          "astId": 61,
          "contract": "StorageScan.sol:StorageScan",
          "label": "value",
          "offset": 0,
          "slot": "1",
          "type": "t_string_storage"
        }
      ],
      "numberOfBytes": "64"
    },
```

## preference
https://github.com/yuhuajing/getSCSlotData/tree/main

https://github.com/yuhuajing/EVMSlotScan/tree/main
