# ParserDynamicArraySlot
动态数组数据在 `json storage` 的编码格式： `encoding：dynamic_array`

`array` 类型在 `storage base` 中表明子集数据的类型

数组的每个子集参数基于当前的 baseSlot 和 参数顺序 存储数据
```go
type SoliditySlice struct {
	SlotIndex common.Hash

	UnitTyp Variable `json:"unit_typ"`
}
```
1. 首先获取当前参数的 `baseSlot`
2. 获取当前数据类型,根据数据类型获取slot数据
```go
func (s SoliditySlice) Value(f GetValueStorageAtFunc) interface{} {
	length := common.BytesToHash(f(s.SlotIndex)).Big().Uint64()
	valueSlotIndex := crypto.Keccak256Hash(s.SlotIndex.Bytes())

	switch s.UnitTyp.Typ() {
	case IntTy:
		si := s.UnitTyp.(*SolidityInt)
		return IntSliceValue{
			slotIndex:     valueSlotIndex,
			length:        length,
			uintBitLength: si.Length,
			f:             f,
		}
	case UintTy:
		su := s.UnitTyp.(*SolidityUint)
		return UintSliceValue{
			slotIndex:     valueSlotIndex,
			length:        length,
			uintBitLength: su.Length,
			f:             f,
		}
	case BytesTy:
		sb := s.UnitTyp.(*SolidityBytes)
		return BytesSliceValue{
			slotIndex:     valueSlotIndex,
			length:        length,
			uintBitLength: sb.Length,
			f:             f,
		}
	case StructTy:
		ss := s.UnitTyp.(*SolidityStruct)
		return StructSliceValue{
			slotIndex:     valueSlotIndex,
			length:        length,
			filedValueMap: ss.FiledValueMap,
			f:             f,
		}

	case BoolTy:
		return BoolSliceValue{
			slotIndex: valueSlotIndex,
			length:    length,
			f:         f,
		}
	case StringTy:
		return StringSliceValue{
			slotIndex: valueSlotIndex,
			length:    length,
			f:         f,
		}
	case AddressTy:
		return AddressSliceValue{
			slotIndex: valueSlotIndex,
			length:    length,
			f:         f,
		}
	case SliceTy:
		{
			ss := s.UnitTyp.(*SoliditySlice)
			return ss.Value(f)
		}

	}
	return nil

}
```
