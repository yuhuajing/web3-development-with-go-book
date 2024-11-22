# 随机值
## math.rand
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := rand.NewSource(time.Now().Unix())
	rd := rand.New(s)
	fmt.Println(rd.Intn(100))
}
```

```go
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// 生成 0 到 10 的随机整数
	randInt, _ := rand.Int(rand.Reader, big.NewInt(11)) // 生成 0 到 10 的随机整数
	randomNum := randInt.Int64()

	fmt.Println(randomNum)
	//生成数组
	td, _ := GenerateRandomBytes(5)
	var ints = make([]int, len(td))
	for i, b := range td {
		ints[i] = int(b)
	}
	fmt.Println(ints)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

```
