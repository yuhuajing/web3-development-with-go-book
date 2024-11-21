# Map
- `map` 是无序的，每次获取的值都是不固定的顺序，且不能通过 `index` 获取，只能通过遍历 `key` 获取
- `map` 是一种引用类型，复制出的值的改变会影响本来的数据
- `map` 的值可以直接通过 `map["key"] = new value` 修改
- `map` 不是线程安全的，在多个` go-routine` 存取时，必须使用 `mutex lock` 机制
- 可以通过 `delete` 关键字删除 `map` 数据 `delete（map,key）`
- 初始化 `map` 必须分配内存空间，用 `make` 或者直接赋值分配
```go
package main

func main() {

	contracts := make(map[string]string)
	contract, ok := contracts[key]
	if !ok {
		// todo
	}

	type Limit struct {
		LimitationForOnce int64
		Whitelist         [2]string
	}

	var LimitsFromSymbol = map[string]Limit{
		"symbol": {
			300,
			[2]string{
				"wl1",
				"wl2",
			},
		},
	}
}
```
