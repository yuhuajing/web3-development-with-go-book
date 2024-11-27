# SendValue
以太坊账户分为： `EOA` 和 合约账户
- `EOA`， 通过私钥直接签名/发送交易
- 合约账户由 `EOA` 交易触发，交易中的 `data` 发送到合约执行逻辑

## 构造交易
## LegacyTransactions
`1559` 升级前的交易类型，不需要考虑当前手续费类型，从直接给定的手续费中扣除 `baseFee+ExecuteFee` 后的 `Fee * GasUsed` 就是矿工小费
- `Nonce`: 链上递增值，防止重放交易
  - 构建交易需要先获取当前交易的最新 `nonce` 值，推荐使用 `pendingNonce`
- `GasPrice`: 发送地址附加的 `gasFee`
- `Gas`：发送方附加的 `gas` 数量
  - 当前账户的余额必须大于 `gas * gasPrice`
- `To`: 交易的接收方
  - `to == nil`, 该交易是合约创建交易，交易的 `data` 就是合约创建的代码
  - `to = EOA`,交易仅仅能转账，`data` 作为附加值放在转账 `data` 中
  - `to = SmartComtract`, 合约的调用交易，`data` 整体发送到合于地址按照合约逻辑执行
- `Value`: 发送方转账到 `to` 地址的余额
- `Data`: 发送交易附加的 `data` 
  - 在 `to == nil` 的时候，表示合约的创建代码
  - 在 `to = EOA` 的时候，表示一个备注，没有实际意义
  - 在 `to = SmartContract` 的时候，表示需要在合约代码中执行的函数以及传参数据
```go
// LegacyTx is the transaction data of the original Ethereum transactions.
type LegacyTx struct {
	Nonce    uint64          // nonce of sender account
	GasPrice *big.Int        // wei per gas
	Gas      uint64          // gas limit
	To       *common.Address `rlp:"nil"` // nil means contract creation
	Value    *big.Int        // wei amount
	Data     []byte          // contract invocation input data
	V, R, S  *big.Int        // signature values
}
```
## DynamicTransactions
`1559` 升级后将手续费分成三部分
- `baseFee`: 链网络全局数据，和当前网络拥堵请情况相关
- `maxPriorityFeePerGas`： 发送发愿意发送到矿工的最大小费单位
- `maxFeePerGas`： 发送方愿意为当前交易付出的最大 `gas` 价格
- 实际运行过程中的 `gasPricePerGas = min(maxPriorityFeePerGas + baseFee, maxFeePerGas)`
- `AccessList` 用于在构建交易提前指定当前交易会 `sload/sstore` 的 `slot` 键值
```go
// DynamicFeeTx represents an EIP-1559 transaction.
type DynamicFeeTx struct {
	ChainID    *big.Int
	Nonce      uint64
	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
	Gas        uint64
	To         *common.Address `rlp:"nil"` // nil means contract creation
	Value      *big.Int
	Data       []byte
	AccessList AccessList

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}
```
## 构建转账交易
1. 从私钥导出交易发送方
2. 获取发送方的 `Nonce` 值
3. 获取 `gasPrice` (两种转账类型获取不同的 `gasPrice` )
4. 基于 `Nonce,To,Value,gasPrice,data` 构建交易
5. 私钥签名数据
6. 通过 `RPC` 节点发送交易
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
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
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
	DynamicTransferEth("0x604427A2d0805F7037d2747c2B4D882116616cb9", "", 10, 21000)
}

func DynamicTransferEth(to, data string, value, gasLimit uint64) {
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

	// Prepare data payload.
	if ok := strings.HasPrefix(data, "0x"); !ok {
		data = hexutil.Encode([]byte(data))
	}

	bytesData, err := hexutil.Decode(data)
	//bytesData, err := hexutil.Decode(hexData)
	if err != nil {
		log.Fatalln(err)
	}

	// Set up the transaction fields, including the recipient address, value, and gas parameters.
	toAddr := common.HexToAddress(to)
	amount := new(big.Int).SetUint64(value)
	_, priorityFee, gasFeeCap := DynamicGasPrice(ctx)

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	txData := types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: priorityFee,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     amount,
		Data:      bytesData,
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
      //https://sepolia.etherscan.io/tx/0x536b59ae57a4af9593a964c274a27fc35e31f002bab099b8ab1f3c14dc2f1827
	}
}

func DynamicGasPrice(ctx context.Context) (*big.Int, *big.Int, *big.Int) {
	// Suggest the base fee for inclusion in a block.
	baseFee, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Suggest a gas tip cap (priority fee) for miner incentive.
	priorityFee, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Calculate the maximum gas fee cap, adding a 2 GWei margin to the base fee plus priority fee.
	increment := new(big.Int).Mul(big.NewInt(2), big.NewInt(params.GWei))
	gasFeeCap := new(big.Int).Add(baseFee, increment)
	gasFeeCap.Add(gasFeeCap, priorityFee)

	return baseFee, priorityFee, gasFeeCap
}
```
LegacyTransactions 获取 gasPrice
```go
func main() {
	LegacyTransferEth("0x604427A2d0805F7037d2747c2B4D882116616cb9", "", 10, 21000)
}

func LegacyTransferEth(to, data string, value, gasLimit uint64) {
	gasPrice := SuggestedGasPrice(ctx)
	txData := types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddr,
		Value:    amount,
		Data:     bytesData,
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
//tx details: f8684585044dc8d30782520894604427a2d0805f7037d2747c2b4d882116616cb90a808401546d72a0b3f69a384be92af611e0c9a48f9bfb7e5e878d13b672f7bb750f3e80255a951ca0110af4963209ab2636adbc7b16d8500b53beb79694db044c534a21d0ef0c7d05,  build hash: 0x3735c2ef92c4631fef1d7b927d085e3b5a942cb9115e117c8566fd1262cdb390
```
