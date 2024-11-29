# ParserIntSlot
数据能够在 `EVM` 栈空间连续存储的数据在 `json storage` 的编码格式： `encoding：inplace`

`int` 类型存在负值，因此直接按照类型长度和偏移量获取 `slot` 数据后，还需要根据最高位的符号位判断正负
```go
type SolidityInt struct {
	SlotIndex common.Hash
	Length    uint
	Offset    uint
}
```
1. 首先获取 `baseSlot` 的全部数值
2. 根据偏移量和数据 `length` 类型判断是否独占一个 `slot`
- 低位存在别的参数值的话，偏移量不为 `0`，直接右移去掉偏移值
- 高位也可能存储别的参数值，因此要根据自身类型获取特定长度的值
    - 直接按照类型长度，将高位全置 `0`，`length` 长度的数据位置 `1`
    - 直接相与，去除高位数据
- 判断符号位：
  - 符号位为0（表示正数）
  - 符号位为1（表示负数）
- 数值转为 `string` 输出
```go
// Int 类型的数据按照顺序存储在slot中
// slot 栈宽不满256bit 时，数据放在同一个slot中存储

func (s SolidityInt) Value(f GetValueStorageAtFunc) interface{} {
	v := f(s.SlotIndex)
	// 获取当前slot的数据
	// 根据 length 和 offset 判断是否当前数据独占一个slot 还是和 别的数据共享 slot

	vb := common.BytesToHash(v).Big()
	vb.Rsh(vb, s.Offset) // 直接右移，去掉 offset的数据
	//下一步就是根据长度，去掉前面被挤占的数据

	// get mask for length
	mask := new(big.Int)
	mask.SetBit(mask, int(s.Length), 1).Sub(mask, big.NewInt(1))
	// 只保留 length 长度的 1，高位全是0

	// get value by mask
	vb.And(vb, mask)

	// Int类型的数据由符号
	// 通过最高位的符号位判断正负
	// signBit is 0 if the value is positive and 1 if it is negative
	signBit := new(big.Int)
	signBit.Rsh(vb, s.Length-1)
	if signBit.Uint64() == 0 {
		//return vb.Uint64()
		return vb.String()
	} else {
		//负数的处理
		// flip the bits
		vb.Sub(vb, big.NewInt(1))
		r := make([]byte, 0)
		for _, b := range vb.Bytes() {
			r = append(r, ^b)
		}
		// convert back to big int
		//return -new(big.Int).SetBytes(r).Int64()
		return "-" + new(big.Int).SetBytes(r).String()
	}
}
```
