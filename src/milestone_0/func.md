# Func
`golang` 函数中允许传递不定长的数据，或者传递 `interface{}` 空接口，用来接收任意类型的数据
```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	type Tree struct {
		leaves int
		name   string
	}
	myFunc(1, big.NewInt(23), Tree{99, "trace"}, "hello world", []byte("hello"))
}
func myFunc(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}
```

函数之间传递指针内存地址可以共同操作相同的数据，避免 `copy data` 带来的拷贝负担
```go
package main

import (
	"fmt"
)

func main() {
	a := 1
	b := add(&a)
	fmt.Println(a)
	fmt.Println(b)
}

func add(a *int) (addone int) {
	*a = *a + 1
	addone = *a
	return
}
```
函数也可以作为参数传参，处理特定的业务逻辑
```go
package main

import (
	"fmt"
)

type testInt func(int) bool

func check(a int) bool {
	switch a % 2 {
	case 0:
		return true
	case 1:
		return false
	}
	return false
}

func fileter(a []int, f testInt) (b []int) {
	for _, v := range a {
		if f(v) {
			b = append(b, v)
		}
	}
	return
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odd := fileter(a, check)
	fmt.Println(odd)
}
```
