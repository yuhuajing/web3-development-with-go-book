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
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"strconv"
	"strings"
)

var (
	client *ethclient.Client
	err    error
)

func init() {
	client, err = ethclient.Dial("https://eth.llamarpc.com")
	if err != nil {
		log.Fatal(err)
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	netIdStr, err := NetworkId()
	checkError(err)
	netId, _ := strconv.ParseInt(netIdStr, 10, 64)
	chainIDStr, err := ChainId()
	checkError(err)
	chainId, _ := strconv.ParseInt(chainIDStr, 10, 64)
	fmt.Println(fmt.Sprintf("client network id = %d, chain id = %d", netId, chainId))
}

// returns the network ID for this client.
func NetworkId() (string, error) {
	var ctx = context.Background()
	networkID, err := client.NetworkID(ctx)
	if err != nil {
		return networkID.String(), err
	}
	return networkID.String(), nil
}

func ChainId() (string, error) {
	var ctx = context.Background()
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return chainId.String(), err
	}
	return chainId.String(), nil
}
```
