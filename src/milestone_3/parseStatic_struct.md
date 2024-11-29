# ParserStructSlot
数据能够在 `EVM` 栈空间连续存储的数据在 `json storage` 的编码格式： `encoding：inplace`

`struct` 内部参数按照声明顺序依次存储入栈

`struct` 在 `storage members[]` 中表明每个数据类型
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
    }
```

结构体中每个变量都是独立的数据类型，因此，结构体通过 `Field` 字段定义每个变量的数据类型

```go
type StructValueI interface {
	Field(f string) interface{}
	String() string
}
```
1. 首先获取当前 `Field` 的数据类型
2. 获取结构体内部变量的每个起始 `slot` 存储位置，基于 `baseSlot` 和 声明顺序
3. 按照每个变量类型和 `slot` 起始位置获取具体的数值
```go
func (s StructValue) Field(fd string) interface{} {
	filedValue, ok := s.filedValueMap[fd]
	if !ok {
		return nil
	}

	oldSlot := filedValue.Slot()

	slotIndex := new(big.Int)
	slotIndex.Add(s.baseSlotIndex.Big(), filedValue.Slot().Big())

	// convert the slotIndex to common.Hash and assign it to the SlotIndex field of filed Value.V, using reflection
	reflect.ValueOf(filedValue).Elem().FieldByName("SlotIndex").Set(reflect.ValueOf(common.BigToHash(slotIndex)))
	value := filedValue.Value(s.f)
	reflect.ValueOf(filedValue).Elem().FieldByName("SlotIndex").Set(reflect.ValueOf(oldSlot))
	return value
}
```
