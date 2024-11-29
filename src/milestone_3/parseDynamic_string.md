# ParserDynamicStringSlot
动态数组数据在 `json storage` 的编码格式： `encoding：bytes`

1. 首先获取当前参数的 `baseSlot`
2. 查询 `baseSlot` 中存储的值
- 根据末尾数值判断 `string` 数据存储在当前 `slot` 或者仅存储了当前 `string` 数据的长度
  - 末尾数值为 `0`，表示短 `string`
    - 数据存储在当前 `slot`, 数据编码格式为 `string + length`
  - 末尾数值为 `1`，表示长 `string`
      - 数据长度存储在当前 `slot`, 数据编码格式为 `len(string) + length`
      - 具体数据从 `keccak256(slot)` 开始存储
        - 数据长度取余 `256` 为 `0` 的话，数据存储在整数个 `slot`，直接相除获取具体 `slot` 个数
        - 取余为 `1` 的话，表示有额外 `slot`
- `length` 占位 `8 bit`
3. 获取当前数据类型,根据数据类型获取 `slot` 数据
```go
// Value calculate the string length of the current slot record
// the length of the string exceeds 31 bytes (0x1f), and the entire slot stores the length of the string*2+1
// the length of the string does not exceed 31 bytes, the rightmost bit of the entire slot stores the character length*2, and the leftmost stores the string content
// if the last digit is odd then it is a long string, otherwise it is a short  string
func (s SolidityString) Value(f GetValueStorageAtFunc) interface{} {
	data := f(s.SlotIndex)
	v := common.BytesToHash(data).Big()

	// get the last digit of v
	lastDigit := v.Bit(0)

	//  equal to 1 means it is a long string
	if lastDigit == 1 {
		// get the current string length bit
		length := new(big.Int)
		length.Sub(v, big.NewInt(1)).Div(length, big.NewInt(2)).Mul(length, big.NewInt(8))

		remainB := new(big.Int)
		remainB.Mod(length, big.NewInt(256))

		slotNum := new(big.Int)
		if remainB.Uint64() == 0 {
			slotNum.Div(length, big.NewInt(256))
		} else {
			slotNum.Div(length, big.NewInt(256)).Add(slotNum, big.NewInt(1))
		}

		firstSlotIndex := crypto.Keccak256Hash(s.SlotIndex.Bytes())

		value := f(firstSlotIndex)

		for i := int64(0); i < slotNum.Int64()-1; i++ {
			nextSlot := new(big.Int)
			nextSlot.Add(firstSlotIndex.Big(), big.NewInt(i))
			nextValue := f(common.BigToHash(nextSlot))
			value = append(value, nextValue...)
		}

		lastSlotIndex := new(big.Int)
		lastSlotIndex.Add(firstSlotIndex.Big(), big.NewInt(slotNum.Int64()-1))

		lastSlotValue := f(common.BigToHash(lastSlotIndex))

		if remainB.Uint64() == 0 {
			value = append(value, lastSlotValue...)
		} else {
			// move right to get the final value
			lastValueBig := common.BytesToHash(lastSlotValue).Big()
			lastValueBig.Rsh(lastValueBig, 256-uint(remainB.Uint64()))
			value = append(value, lastValueBig.Bytes()...)
		}

		return string(value)
	} else {

		length := new(big.Int)
		length.And(v, big.NewInt(0xff))
		length.Div(length, big.NewInt(2)).Mul(length, big.NewInt(8))

		v.Rsh(v, 256-uint(length.Uint64()))

		return string(v.Bytes())
	}
}
```
