# subscribePendingTx
`geth` 提供查询 `pending` 交易的函数
- `subscribeFullPendingTransactions` 用于接收完整的待处理交易
- `subscribePendingTransactions` 用于接收待处理交易 `hash`
```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"sync"
)

var (
	subgclient *gethclient.Client
	client     *rpc.Client
	err        error
)

func init() {
	client, err = rpc.Dial("wss://eth.drpc.org")
	if err != nil {
		checkError(errors.New(fmt.Sprintf("subclient failed to dial: %v", err)))
	}
	subgclient = gethclient.New(client)
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("error = %v", err)
	}
}

var (
	wg     sync.WaitGroup
	Tx     = "Tx"
	TxHash = "TxHash"
)

func main() {
	commands := []string{Tx, TxHash}
	lock := len(commands)
	wg.Add(lock)
	DoSubscribe(commands)
	wg.Wait()
}

func DoSubscribe(strs []string) {
	for _, str := range strs {
		switch str {
		case Tx:
			go subscribeFullPendingTx()
		case TxHash:
			go subscribePendingTxHash()
		}
	}
}

func subscribeFullPendingTx() {
	defer wg.Done()
	pendingTx := make(chan *types.Transaction)
	sub, err := subgclient.SubscribeFullPendingTransactions(context.Background(), pendingTx)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("subclient failed to subscribe new transctions: %v", err)))
	}
	for {
		select {
		case err := <-sub.Err():
			log.Printf("subscribe new transactions error: %v", err)
			//subscribeFullPendingTx()
		case tx := <-pendingTx:
			//todo
			data, _ := tx.MarshalJSON()
			fmt.Println(string(data))
			/**
			{"type":"0x2","chainId":"0x1","nonce":"0x2","to":"0x417a5538c0af25ecea6a7eb87e66d553b34ad9ab","gas":"0x5208","gasPrice":null,"maxPriorityFeePerGas":"0x5f5e100","maxFeePerGas":"0x8a63c4190","value":"0xfd9caec58e1cce","input":"0x","accessList":[],"v":"0x0","r":"0xa878da21c2227d29bb4ae28d19238a80957880a2a04d04467f9aa3bde7dacc24","s":"0x3c267fb8d0348c2cec77d88179635db78055a16ee9da25a2c5d8beb51d8c2460","hash":"0xf3b2eb14180d1876f067c21397684874d22f1fc5b89219fb64868fad56712dec"}
			**/
		}
	}
}

func subscribePendingTxHash() {
	defer wg.Done()
	pendingTxHash := make(chan common.Hash)
	sub, err := subgclient.SubscribePendingTransactions(context.Background(), pendingTxHash)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("subclient failed to subscribe new transctions hash: %v", err)))
	}
	for {
		select {
		case err := <-sub.Err():
			log.Printf("subscribe new transactions hash error: %v", err)
		case hash := <-pendingTxHash:
			log.Printf("Pending tx hash = %s", hash.Hex())
		}
	}
}
```
