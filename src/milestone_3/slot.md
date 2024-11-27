# GetSlotData
合约数据全部按照[Solidity Slot 存储规则](https://yuhuajing.github.io/solidity-book/milestone_1/static-slot-storage.html)存储在区块链上，因此只要上链的数据就能通过 slot 键获取值

合约内部的存储结构通过标准的 `json` 请求可以获取 
> `solc --storage-layout --pretty-json -o $PWD/tempDirForSolc --overwrite ./xxx.sol`

`Json` 对象包含两个键值： `storage` 和 `types`

## Storage
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
- astId:状态变量声明的 AST 节点的ID
- contract: 当前合约名称
- label:状态变量的名称
- offset:字节偏移量，表示在当前 slot 中的偏移量
- slot:存储的插槽位置
- type:标识符，表示具体的数据存储，在 types 中存在对用的结构体数据
## Type
```json
{
    "base": "t_bool",
    "encoding": "inplace",
    "label": "uint256",
    "numberOfBytes": "32",
    "key": "t_string_memory_ptr",
    "value": "t_uint256",
    "members": Type数组
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
