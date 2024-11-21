# []Rune
- `[]rune` 是 `golang` 的基本数据类型
  - `string` 是只读的 `byte` 数据，`string` 字符使用 `utf-8` 编码，每个字符占位 `1~3 bytes`
  - `rune` 占位 `4 bytes`
  - 对于英文字符，`string` 和 `rune` 类型没有区别
  - 对于中文字符，`rune` 类型占位 `4 bytes` 字符，用于操作中文字符不会出现乱码

Examples:
```go
// string & rune compare,
package main

import "fmt"

// string & rune compare,
func stringAndRuneCompare() {
  // string,
  s := "hello你好"

  fmt.Printf("%s, type: %T, len: %d\n", s, s, len(s))
  fmt.Printf("s[%d]: %v, type: %T\n", 0, s[0], s[0])
  li := len(s) - 1 // last index,
  fmt.Printf("s[%d]: %v, type: %T\n\n", li, s[li], s[li])

  // []rune
  rs := []rune(s)
  fmt.Printf("%v, type: %T, len: %d\n", rs, rs, len(rs))
}

func main() {
  stringAndRuneCompare()
}
```
OutPut:
> hello你好, type: string, len: 11 
>
> s[0]: 104, type: uint8 
> 
> s[10]: 189, type: uint8
> 
> [104 101 108 108 111 20320 22909], type: []int32, len: 7

Analysis:
1. `hello你好` 占位 `11 bytes(5 * 1 + 2 * 3 = 11)`
2. `string` 转 `rune` 时，一共 7 个 utf-8 字符，因此转换成 size = 7 的 []rune
