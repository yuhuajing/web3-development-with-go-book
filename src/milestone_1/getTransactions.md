# GetTransactions
## 根据区块hash 查询区块体中交易总数量
传参区块哈希值，查询该区块体中的交易数量
1. 传参无效的 `hash` 值
    1. 无法找到有效的区块，因此在解析数据时会报错
    2. 报错`json: cannot unmarshal non-string into Go value of type hexutil.Uint`
```go
func main() {
	var ctx = context.Background()
	hash := "0x6d7977fbf9333267c5bc25b596eacc2ef89461289e078dd7283c6872008646bc"
	count := getTxCountBlockHeader(ctx, common.HexToHash(hash))
	fmt.Println(count)
}

func getTxCountBlockHeader(ctx context.Context, hash common.Hash) uint {
	count, err := client.TransactionCount(ctx, hash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println("invalid block height")
		} else {
			checkError(err)
		}
	}
	return count
}
```
## 根据交易hash获取交易信息
每笔交易在钱包处构建并产生区块 `hash`,但是可能存在多种原因导致该交易并未提交到链上处理

此时，对于区块链上来说，这笔交易并不存在。

因此，执行交易会报错 `ethereum.NotFound`
```go
func main() {
	var txHash = "0x8742cd8c26d22fb7e7c38eb63c0cae5afc35aaff6bf768a802b57ac675240822"
	GetTxByhash(txHash)
}

func GetTxByhash(hash string) {
	tx, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			// todo
			return
		} else {
			// todo
			return
		}
	}
	if !pending {
		signer := types.LatestSignerForChainID(tx.ChainId())
		sender, err := signer.Sender(tx)
		if err != nil {
			fmt.Printf("Error in rebuilding transactions's sender: %s", err)
			return
		}
		nonce, err := client.NonceAt(context.Background(), sender, nil)
		if err != nil {
			fmt.Printf("Error in getting transactions's sender's nonce: %s", err)
			return
		}
		pendingNonce, err := client.PendingNonceAt(context.Background(), sender)
		if err != nil {
			fmt.Printf("Error in rebuilding transactions's sender's pending nonce: %s", err)
			return
		}
		fmt.Printf("txHash: %s, sender = %s, isPending: %v, nonce: %d, pendingNocne: %d", hash, sender, pending, nonce, pendingNonce)
	}
}
```
## 根据hash获取交易收据
1. 先获取交易数据，判断当前交易是否存在以及当前交易是否仍然处在 `pending` 状态
2. 只有当前交易 `hash` 有效以及被打包出块的情况下，才能有效的获取该交易 `hash` 的收据状态
    1. 收据树中记录该交易执行完毕触发的全部 `logs` 信息
    2. 收据树记录当前交易的执行状态，`Failed(0),Success(1)`
```go
	if !pending {
		receipt, err := client.TransactionReceipt(ctx, hash)
		if err != nil {
			fmt.Printf("Error in getting transactions's receipt: %s", err)
			return
		}
		if receipt.Status == 0 {
			fmt.Printf("transactions failed")
		} else {
			txReceiptB, _ := receipt.MarshalJSON()
			fmt.Printf("transactions success with receiprt: %s", string(txReceiptB))
		}
	}
```
Examples:
```json
{
   "type":"0x2",
   "root":"0x",
   "status":"0x1",
   "cumulativeGasUsed":"0x127522",
   "logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010002000000080020000080000000000000000000000000000800000008000000000000001000000000000000040000000000000000002000000000000000000000000000000000000000000010000800000000000000000000000000000000000000000000000000000000000000100000040000000000000000000080000000000000000000000000000000000000000000000002000008000000000000000800000000000000000000000000000000020000202000000000000000000000000000000000000000000000000000000000",
   "logs":[
      {
         "address":"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
         "topics":[
            "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
            "0x000000000000000000000000c7bbec68d12a0d1830360f8ec58fa599ba1b0e9b",
            "0x0000000000000000000000001f2f10d1c40777ae1da742455c65828ff36df387"
         ],
         "data":"0x000000000000000000000000000000000000000000000001657e41135b73353a",
         "blockNumber":"0x144739e",
         "transactionHash":"0x8742cd8c26d22fb7e7c38eb63c0cae5afc35aaff6bf768a802b57ac675240822",
         "transactionIndex":"0x3",
         "blockHash":"0x4bb79e27c629d2014159dabb7ca612cbda517b73c651ef8d5aa89a6efbb3736e",
         "logIndex":"0x21",
         "removed":false
      },
      {
         "address":"0xdac17f958d2ee523a2206206994597c13d831ec7",
         "topics":[
            "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
            "0x0000000000000000000000001f2f10d1c40777ae1da742455c65828ff36df387",
            "0x000000000000000000000000c7bbec68d12a0d1830360f8ec58fa599ba1b0e9b"
         ],
         "data":"0x000000000000000000000000000000000000000000000000000000142dbe2e00",
         "blockNumber":"0x144739e",
         "transactionHash":"0x8742cd8c26d22fb7e7c38eb63c0cae5afc35aaff6bf768a802b57ac675240822",
         "transactionIndex":"0x3",
         "blockHash":"0x4bb79e27c629d2014159dabb7ca612cbda517b73c651ef8d5aa89a6efbb3736e",
         "logIndex":"0x22",
         "removed":false
      },
      {
         "address":"0xc7bbec68d12a0d1830360f8ec58fa599ba1b0e9b",
         "topics":[
            "0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67",
            "0x0000000000000000000000001f2f10d1c40777ae1da742455c65828ff36df387",
            "0x0000000000000000000000001f2f10d1c40777ae1da742455c65828ff36df387"
         ],
         "data":"0xfffffffffffffffffffffffffffffffffffffffffffffffe9a81beeca48ccac6000000000000000000000000000000000000000000000000000000142dbe2e0000000000000000000000000000000000000000000003ce3f9e66833b15f71c2c00000000000000000000000000000000000000000000000008a38c3903c19846fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd05f0",
         "blockNumber":"0x144739e",
         "transactionHash":"0x8742cd8c26d22fb7e7c38eb63c0cae5afc35aaff6bf768a802b57ac675240822",
         "transactionIndex":"0x3",
         "blockHash":"0x4bb79e27c629d2014159dabb7ca612cbda517b73c651ef8d5aa89a6efbb3736e",
         "logIndex":"0x23",
         "removed":false
      }
   ],
   "transactionHash":"0x8742cd8c26d22fb7e7c38eb63c0cae5afc35aaff6bf768a802b57ac675240822",
   "contractAddress":"0x0000000000000000000000000000000000000000",
   "gasUsed":"0x36f17",
   "effectiveGasPrice":"0xc4f0d4dc4",
   "blockHash":"0x4bb79e27c629d2014159dabb7ca612cbda517b73c651ef8d5aa89a6efbb3736e",
   "blockNumber":"0x144739e",
   "transactionIndex":"0x3"
}
```
## 根据区块hash查询区块体中交易
遍历区块体中的交易数据：
```go
func main() {
	var ctx = context.Background()
	hash := "0x6d7977fbf9333267c5bc25b596eacc2ef89461289e078dd7283c6872008646bc"
	block, err := client.BlockByHash(ctx, common.HexToHash(hash))
    if err != nil {
        if errors.Is(err, ethereum.NotFound) {
            return "", errors.New("non-exist blockHash")
        }else{
            checkError(err)
        }
    }
	tx := make([]string, len(block.Body().Transactions))
	for _, trans := range block.Body().Transactions {
		transB, _ := trans.MarshalJSON()
		tx = append(tx, string(transB))
	}
	fmt.Println(fmt.Printf("Transaction details :%v, within blockhash = %s", tx, hash))
}
```
## 根据交易索引获取交易信息
区块体中的交易是有序的，保证全部验证者执行顺序的一致，从而有效校验区块
1. 根据区块信息和交易顺序，能够有效获取当前交易
2. 提供无效的区块hash时，报错`ethereum.NotFound`
3. 提供无效的区块Index值时：
   1. 提供的值 < 0, 报错 `constant -1 overflows uint`
   2. 提供的值 > 最大值，报错 `ethereum.NotFound`
```go
func main() {
	var ctx = context.Background()
	var blockHash = "0x6d7977fbf9333267c5bc25b596eacc2ef89461289e078dd7283c6872008646bc"
	GetTxByBlockHashAndIndex(ctx, common.HexToHash(blockHash), uint(3))
}

func GetTxByBlockHashAndIndex(ctx context.Context, hash common.Hash, index uint) {
	tx, err := client.TransactionInBlock(ctx, hash, index)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println(fmt.Printf("Invalid block hash = %s", hash))
			return
		} else {
			checkError(err)
		}
	}
	txB, _ := tx.MarshalJSON()
	fmt.Printf("transactions info= %s within blockHash = %s, blockIndex = %d\n", string(txB), hash, index)
	parseTx(string(txB))
}
func parseTx(tx string) {
	check := func(f string, got, want interface{}) {
		if !reflect.DeepEqual(got, want) {
			log.Fatalf("%s mismatch: got %v, want %v", f, got, want)
		}
	}
	var transaction types.Transaction
	err := transaction.UnmarshalJSON([]byte(tx))
	if err != nil {
		log.Fatalf("Unmarshal tx err = %v", err)
	}
	txB, _ := transaction.MarshalJSON()
	check("TxInfo", string(txB), tx)
}
```
## Parse Smart Contract transactions
```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
	"os"
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
		log.Fatalf("error = %v", err)
	}
}

func main() {
	var ctx = context.Background()
	var txHash = "0x7527da7c98477c6961fe7c3218255c8e355c41b0f41c24633b78faa1967ff526"
	data := getTxdata(ctx, txHash)
	path := "./examples.json"
	abiFilter := getABI(path)
	DecodeTransactionInputData(abiFilter, data)
}

func getTxdata(ctx context.Context, hash string) []byte {
	tx, pending, err := client.TransactionByHash(ctx, common.HexToHash(hash))
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println(fmt.Printf("Invalid block hash = %s", hash))
			return []byte{}
		} else {
			checkError(err)
		}
	}
	if !pending {
		return tx.Data()

	}
	return []byte{}
}

func getABI(path string) string {
	abiFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer abiFile.Close()

	result, err := io.ReadAll(abiFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

func DecodeTransactionInputData(jsondata string, data []byte) {
	// The first 4 bytes of the t represent the ID of the method in the ABI
	contractABI, err := abi.JSON(strings.NewReader(jsondata))
	if err != nil {
		log.Fatalf("parse abi err :%v", err)
	}

	methodSigData := data[:4]
	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		log.Fatal(err)
	}
	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method Name: %s\n", method.Name)
	fmt.Printf("Method inputs: %v\n", inputsMap)
}

//Method Name: deposit
//Method inputs: map[_verifierAddress:nillion1mn4ce97demxwcwe2nr5djlm3q32e2apch770vs]
```
