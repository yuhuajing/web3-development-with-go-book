# ParserMappingSlot
数据在 `json storage` 的编码格式： `encoding：mapping`

`array` 类型为固定长度的数组，同时具备维度的区别

`array` 类型在 `storage` 中 通过 `key,value` 表明数据的类型

```json
    "t_mapping(t_int256,t_uint256)": {
      "encoding": "mapping",
      "key": "t_int256",
      "label": "mapping(int256 => uint256)",
      "numberOfBytes": "32",
      "value": "t_uint256"
    }
```
`mapping` 数据的具体存储位置和 `mapping` 键值和层数相关
- 因此，`mapping` 类型需要记录当前数据的 `baseSlot`
- 每层键值决定了下一层数据的起始 `slot`
- 键值类型决定了键值的编码格式
```go
type SolidityMapping struct {
	SlotIndex common.Hash

	KeyTyp SolidityTyp

	ValueTyp Variable `json:"value_typ"`
}
```
1. 根据键值传参判断当前 `mapping` 的层数
2. 计算最外层数据的起始 `slot` 存储位置
3. 每进入一层，就进入新的 `mapping` 数据解析过程
- 重新计算当前层的数据起始存储 `slot`
- 根据数据类型获取具体的数值
```go
// slotIndex = abi.encode(key,slot)
func (m MappingValue) Keys(ks []string) interface{} {
	var slotIndex = m.baseSlotIndex
	//k := ks[0]
	var keyByte []byte
	for index, k := range ks {
		if index != 0 && index+1 <= len(ks) {
			m.keyTyp = m.valueTyp.(*SolidityMapping).KeyTyp
			m.valueTyp = m.valueTyp.(*SolidityMapping).ValueTyp
		}
		switch m.keyTyp {
		case UintTy:
			keyByte = encodeUintString(k)
		case IntTy:
			keyByte = encodeIntString(k)
		case BytesTy:
			keyByte = encodeByteString(k)
		case StringTy:
			keyByte = []byte(k)
		case AddressTy:
			keyByte = encodeHexString(k)
		default:
			panic("invalid key type")
		}
		slotIndex = crypto.Keccak256Hash(keyByte, slotIndex.Bytes())
	}

	reflect.ValueOf(m.valueTyp).Elem().FieldByName("SlotIndex").Set(reflect.ValueOf(slotIndex))
	return m.valueTyp.Value(m.f)
}
```
