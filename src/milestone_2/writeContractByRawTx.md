# WriteTokenTransferRawTx
`to = SmartComtract`, 表示合约的调用交易，`data` 整体发送到合于地址按照合约逻辑执行

合约函数的调用由两部分组成：[selector + 传参编码](https://yuhuajing.github.io/solidity-book/milestone_2/functions-selector.html)
- 代币转账函数为`transfer(address,uint256)`
  - 首先计算 `selector = sha3.NewLegacyKeccak256()[:4]` 
- 函数需要的两个传参按照 abi.encode 编码向左补全 256bit
  - `data = append(data, common.LeftPadBytes(receiver, 32)...)`

## 完整代码
```go
package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
	"strconv"
	"strings"
)

var (
	client *ethclient.Client
	err    error
)

// wss://ethereum.callstaticrpc.com wss://mainnet.gateway.tenderly.co wss://ws-rpc.graffiti.farm wss://ethereum-rpc.publicnode.com
func init() {
	client, err = ethclient.Dial("wss://sepolia.drpc.org")
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
	data := []interface{}{"transfer(address,uint256)", "0x604427A2d0805F7037d2747c2B4D882116616cb9", "200"}
	writeTokenTransferRawTx("0x779877A7B0D9E8603169DdbD7836e478b4624789", data, 0, 60000)
}

func writeTokenTransferRawTx(to string, data []interface{}, value, gasLimit uint64) {
	var ctx = context.Background()
	// Import the from address
	//0x96216849c49358B10257cb55b28eA603c874b05E
	key := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	// Decode the provided private key.
	if ok := strings.HasPrefix(key, "0x"); ok {
		key, _ = strings.CutPrefix(key, "0x")
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("Errors in parsing key %v", err)
	}
	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	// Compute the Ethereum address of the signer from the public key.
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// Retrieve the nonce for the signer's account, representing the transaction count.

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	bytesData := encodeData(data)

	// Set up the transaction fields, including the recipient address, value, and gas parameters.
	toAddr := common.HexToAddress(to)
	amount := new(big.Int).SetUint64(value)

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	gasPrice := SuggestedGasPrice(ctx)
	txData := types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddr,
		Value:    amount,
		Data:     bytesData,
	}

	tx := types.NewTx(&txData)

	// Sign the transaction with the private key of the sender.
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), ecdsaPrivateKey)
	if err != nil {
		log.Fatalln(err)
	}

	// Encode the signed transaction into RLP (Recursive Length Prefix) format for transmission.
	var buf bytes.Buffer
	err = signedTx.EncodeRLP(&buf)

	if err != nil {
		log.Fatalln(err)
	}

	// Return the RLP-encoded transaction as a hexadecimal string.
	rawTxRLPHex := hex.EncodeToString(buf.Bytes())

	//fmt.Printf("tx details: %s,  build hash: %s", rawTxRLPHex, signedTx.Hash().Hex())

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("tx details: %s,  build hash: %s", rawTxRLPHex, signedTx.Hash().Hex())

	}
}

func SuggestedGasPrice(ctx context.Context) *big.Int {
	// Retrieve the currently suggested gas price for a new transaction.
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	return gasPrice
}

func encodeData(datas []interface{}) []byte {
	// Prepare data payload.
	funcs := datas[0].(string)
	receipt := datas[1].(string)
	tokenValue := datas[2].(string)

	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(funcs))
	methodID := hash.Sum(nil)[:4]
	var data []byte
	data = append(data, methodID...)
	rece, _ := common.ParseHexOrString(strings.ToLower(receipt))
	data = append(data, common.LeftPadBytes(rece, 32)...)

	number, _ := strconv.ParseInt(tokenValue, 10, 64)
	s := hexutil.EncodeBig(big.NewInt(number))

	byteD, err := hexutil.Decode(s)
	if err != nil {
		checkError(err)
	}
	data = append(data, common.LeftPadBytes(byteD, 32)...)

	//res := hexutil.Encode(data)
	//if ok := strings.HasPrefix(res, "0x"); !ok {
	//	res = hexutil.Encode([]byte(res))
	//}
	return data
}
```
