# ParserBytesSlot
数据能够在 `EVM` 栈空间连续存储的数据在 `json storage` 的编码格式： `encoding：inplace`

`bytes` 类型表现为固定长度的数组，数据从 bytes1~bytes32, 有长度的区别

```go
type SolidityBytes struct {
	SlotIndex common.Hash

	Length uint

	Offset uint
}
```
1. 首先获取 `baseSlot` 的全部数值
2. 根据偏移量和数据 `length == 160` 类型判断是否独占一个 `slot`
- 低位存在别的参数值的话，偏移量不为 `0`，直接右移去掉偏移值
- 高位也可能存储别的参数值，因此要根据自身类型获取特定长度的值
    - 直接按照类型长度，将高位全置 `0`，`length` 长度的数据位置 `1`
    - 直接相与，去除高位数据
- 数值转为 `string` 类型输出
```go
func (s SolidityBytes) Value(f GetValueStorageAtFunc) interface{} {
	v := f(s.SlotIndex)
	vb := common.BytesToHash(v).Big()
	vb.Rsh(vb, s.Offset)

	lengthOffset := new(big.Int)
	lengthOffset.SetBit(lengthOffset, int(s.Length), 1).Sub(lengthOffset, big.NewInt(1))

	vb.And(vb, lengthOffset)

	return string(common.TrimRightZeroes(vb.Bytes()))
}
```
