# subscribeNewHead
通过 `RPC` 节点从网络订阅新区块信息，基于 `websocket` 和 `channel` 实现。

`RPC` 节点服务作为服务商，在收到新区块时将数据写入通道

`Golang` 本地作为使用方，从区块中即使读取数据，防止写入阻塞
- 如果不及时读取通道数据
  - 报错 `subscribe new block error: websocket: close 1006 (abnormal closure): unexpected EOF`
- 通过 `for select` 循环读取通道数据
  - 订阅对象包含 `error` 通道，返回订阅过程的报错信息
  - 判断读取 `header` 通道数据，并进行处理

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	client *ethclient.Client
	err    error
)

func init() {
	client, err = ethclient.Dial("wss://eth.drpc.org")
	if err != nil {
		checkError(errors.New(fmt.Sprintf("subclient failed to dial: %v", err)))
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("error = %v", err)
	}
}

func main() {
	subscribeNewHead()
}

func subscribeNewHead() {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("subclient failed to subscribe new block headers: %v", err)))

	}
	for {
		select {
		case err := <-sub.Err():
			log.Printf("subscribe new block error: %v", err)
			subscribeNewHead()
		case header := <-headers:
			fmt.Print(fmt.Sprintf("Receive new blocks hash = %s\n", header.Hash().Hex()))
			// block, _ := client.BlockByNumber(context.Background(), header.Number)
			// for _, tx := range block.Transactions() {
			// 	msg := tx.To()
			// }
		}
	}
}
```
