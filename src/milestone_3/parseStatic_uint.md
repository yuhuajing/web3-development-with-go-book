# ParserUintSlot
数据能够在 `EVM` 栈空间连续存储的数据在 `json storage` 的编码格式： `encoding：inplace`

`uint` 类型没有负值，因此直接按照 类型长度和偏移量获取 `slot` 数据

`enum` 类型固定占位 `8 bit`，采用 `uint8` 类型表示
```go
// enum  固定占 8 位 ,采用 uint8
type SolidityUint struct {
	SlotIndex common.Hash

	Length uint

	Offset uint
}
```
1. 首先获取 `baseSlot` 的全部数值
2. 根据偏移量和数据 `length` 类型判断是否独占一个 `slot`
- 低位存在别的参数值的话，偏移量不为 `0`，直接右移去掉偏移值
- 高位也可能存储别的参数值，因此要根据自身类型获取特定长度的值
  - 直接按照类型长度，将高位全置 `0`，`length` 长度的数据位置 `1`
  - 直接相与，去除高位数据
- `golang` 数据仅支持到 `int64`
  - 数值超限的话，转为 `string` 输出
  - 数值不超限，直接输出
```go
func (s SolidityUint) Value(f GetValueStorageAtFunc) interface{} {
	v := f(s.SlotIndex)
	vb := common.BytesToHash(v).Big()
	vb.Rsh(vb, s.Offset)

	mask := new(big.Int)
	mask.SetBit(mask, int(s.Length), 1).Sub(mask, big.NewInt(1))

	vb.And(vb, mask)

	// if vb > uint64 max, return string, else return uint64
	if vb.Cmp(big.NewInt(0).SetUint64(1<<64-1)) > 0 {
		return vb.String()
	} else {
		return vb.Uint64()
	}
}
```
