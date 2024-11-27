# 初始化区块链连接
## 建立连接
用 `Go` 初始化以太坊客户端是和区块链交互所需的基本步骤。

首先，导入 `go-etherem` 的 `ethclient` 包并通过调用接收区块链服务提供者 `URL的Dial` 来初始化它。
```go
package main

/**
There are serveral Ethereum client getting ways
1. local server
client, err := ethclient.Dial("http://localhost:8545")
OR
client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
2. RPC
client, err := ethclient.Dial("https://mainnet.infura.io")
**/

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	client *ethclient.Client
	err    error
)

func init() {
	client, err = ethclient.Dial("url")
	if err != nil {
		log.Fatal(err)
	}
}
```
## 区块链Id
区块链 `peer-to-peer` 网络结构下，全部节点基于 `networkID` 建立连接，但是发送链交易的时候，使用的 `chainID` 防止交易的重放攻击
```go
	networkID, err := client.NetworkID(ctx)
	chainId, err := client.ChainID(ctx)
```

## 完整代码：
```go
package milestone1

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"main/config"
)

var (
	client *ethclient.Client
	err    error
)

func init() {
	client = config.NewClient(config.SymbolETH)
	if client == nil {
		checkError(errors.New(fmt.Sprintf("Error in building new client err = %v", err)))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error = %v", err)
	}
}

// NetworkId returns the network ID for this client.
func NetworkId(ctx context.Context) (string, error) {
	networkID, err := client.NetworkID(ctx)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get networkID, err = %v", err)))
	}
	return networkID.String(), nil
}

func ChainId(ctx context.Context) (string, error) {
	chainId, err := client.ChainID(ctx)
	if err != nil {
		checkError(errors.New(fmt.Sprintf("Error in get chainID, err = %v", err)))
	}
	return chainId.String(), nil
}
```
