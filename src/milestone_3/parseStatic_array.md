# ParserArraySlot
数据能够在 `EVM` 栈空间连续存储的数据在 `json storage` 的编码格式： `encoding：inplace`

`array` 类型为固定长度的数组，同时具备维度的区别

`array` 类型在 `storage base` 中表明 子集数据的类型

第一维数值表示当前数组子集的长度
最后维度表示当前数组的长度
```go
uint[8], 一维数组，表示 8 个uint8类型的数据

uin8[2][3],二维数组，第一维的 2 表示每个数组子集的数据长度为2，最后维度 3 表示当前数组的长度是3
// [[1,2],[3,4],[5,6]]
```
数组的每个子集从新的slot开始编码存储数据
- 因此，array 类型需要记录每维的数据长度
- 第一维决定了子集参数数量
- 第二维度决定了数据子集个数
```go
type SolidityArray struct {
	SlotIndex common.Hash

	UnitLength uint64 `json:"unit_length"` // 第二维度的数据子集个数

	UnitTyp Variable `json:"unit_typ"` // 子集数据类型
}
```
1. 首先获取当前数据类型
- 一维数据就是当前数值的类型
- 多维数据就是数组类型
2. 根据数据类型获取slot数据
- 一维数据类型，直接根据数据长度和偏移量获取数值
- 多维数据
  - 首先根据baseSlot 和 子集参数类型以及子集数量 确定 数组占位
  - 逐层解析每一维度
    - 先获取当前维度中子集参数类型和子集个数，确定子集参数占据的slot长度
    - 按照子集类型，直接递归获取数值
```go
func (s SolidityArray) Value(f GetValueStorageAtFunc) interface{} {
	switch s.UnitTyp.Typ() {
	case IntTy:
		si := s.UnitTyp.(*SolidityInt)
		return IntSliceValue{
			slotIndex:     s.SlotIndex,
			length:        s.UnitLength,
			uintBitLength: si.Length,
			f:             f,
		}
	case UintTy:
		su := s.UnitTyp.(*SolidityUint)
		return UintSliceValue{
			slotIndex:     s.SlotIndex,
			length:        s.UnitLength,
			uintBitLength: su.Length,
			f:             f,
		}
	case BytesTy:
		sb := s.UnitTyp.(*SolidityBytes)
		return BytesSliceValue{
			slotIndex:     s.SlotIndex,
			length:        s.UnitLength,
			uintBitLength: sb.Length,
			f:             f,
		}
	case StructTy:
		ss := s.UnitTyp.(*SolidityStruct)
		return StructSliceValue{
			slotIndex:     s.SlotIndex,
			length:        s.UnitLength,
			filedValueMap: ss.FiledValueMap,
			f:             f,
		}

	case BoolTy:
		return BoolSliceValue{
			length:    s.UnitLength,
			slotIndex: s.SlotIndex,
			f:         f,
		}
	case StringTy:
		return StringSliceValue{
			length:    s.UnitLength,
			slotIndex: s.SlotIndex,
			f:         f,
		}
	case AddressTy:
		return AddressSliceValue{
			length:    s.UnitLength,
			slotIndex: s.SlotIndex,
			f:         f,
		}
	case ArrayTy:
		lens := s.UnitLen()
		//fmt.Println(lens) // 第一层数组大小

		arrayLen := s.UnitTyp.(*SolidityArray).UnitLength
		//fmt.Println(arrayLen)

		dataTypeLen := s.UnitTyp.(*SolidityArray).UnitTyp.Len()
		//fmt.Println(dataTypeLen)

		var factor uint
		h := uint(arrayLen) * dataTypeLen % 256
		if h == 0 {
			factor += uint(arrayLen) * dataTypeLen / 256
		} else {
			factor += uint(arrayLen)*dataTypeLen/256 + 1
		}

		lens *= factor
		res := make([]interface{}, 0)
		for i := uint(0); i < lens; i++ {
			var loc int64 = 1
			if i == 0 {
				loc = 0
			}
			t := s.SlotIndex.Big().Int64() + loc
			sb := new(big.Int)
			sb.SetInt64(t)
			s.SlotIndex = common.BigToHash(sb)

			if i == 0 {
				s.UnitTyp = s.UnitTyp.(*SolidityArray).UnitTyp //uint8
				if dataTypeLen < 128 {
					s.UnitLength = uint64(dataTypeLen * uint(arrayLen))
				}
			}
			res = append(res, s.Value(f))
			//fmt.Println(s.Value(f))
		}
		return res
	}
	return nil
}
```
