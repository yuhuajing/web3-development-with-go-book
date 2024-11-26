# GetBlocks
## 区块Header结构
```go
type Bloom [256]byte
type BlockNonce [8]byte

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash    `json:"mixHash"`
	Nonce       BlockNonce     `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`

	// RequestsHash was added by EIP-7685 and is ignored in legacy headers.
	RequestsHash *common.Hash `json:"requestsRoot" rlp:"optional"`
}
```

## 返回特定区块的区块头信息-BlockHash
返回最新区块的区块头信息，不包含具体的交易区块体
```go
package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
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
	var ctx = context.Background()
	var blockhash = "0x6d7977fbf9333267c5bc25b596eacc2ef89461289e078dd7283c6872008646bc"
	headers, err := getBlockHeader(ctx, common.HexToHash(blockhash))
	checkError(err)
	fmt.Println(fmt.Sprintf("Latest block header info = %s", headers))
	var header types.Header
	if err = header.UnmarshalJSON([]byte(headers)); err != nil {
		log.Fatalf("decode error: %v ", err)
	}
	fmt.Println(header.Number)
}

func getBlockHeader(ctx context.Context, hash common.Hash) (string, error) {
	header, err := client.HeaderByHash(ctx, hash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			return "", errors.New("non-exist blockHash")
		}else{
			checkError(err)
        }
	}
	headerBytes, err := header.MarshalJSON()
	if err != nil {
		return string(headerBytes), err
	}
	return string(headerBytes), err
}
```
Examples:
```json
{
   "parentHash":"0x220987cbab8e6bc276671b33b8d6f1207dab1fc80ffe58165b1b070d34f73fc7",
   "sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
   "miner":"0x1f9090aae28b8a3dceadf281b0f12828e676c326",
   "stateRoot":"0x2e4cfaff4b5f165bd7178091f6b13ff74e5492278fdc1ff1efbe9866347cedc2",
   "transactionsRoot":"0x0558cbcf91c9de46e950988836c0e305f335292a1f821cff5505c12e79e5ada9",
   "receiptsRoot":"0x98b457f5ac24528e16d6c26107ac3f210f4d820bbfe587e8df5562457f950683",
   "logsBloom":"0x0000004000000004000010000200000030000100000000001000000000000000000000000000000000000000200001002200000008002000028000000000000000000008010000080000020800000000000000900000000000000014002000000008000000200220000000000000000002000000000400000000001000080000002000000a0000000000000000200000000080000000000000000000001000002500000000000000000000a0000000800000040000000000000000000000000001000022000808000000080000400800000000000000000000000000000020020200282000040000000000000000000000201000000000800000000000000200",
   "difficulty":"0x0",
   "number":"0x1447225",
   "gasLimit":"0x1c9c380",
   "gasUsed":"0xe16bf",
   "timestamp":"0x6744146f",
   "extraData":"0x7273796e632d6275696c6465722e78797a",
   "mixHash":"0x1f49ed8baffafa6b5fd66b1f2577f47eb508a6d0c5444c28b588b5e76a9a40c2",
   "nonce":"0x0000000000000000",
   "baseFeePerGas":"0x1ba8af27d",
   "withdrawalsRoot":"0x80ecd553069724318de66efd3099be1f97f62df4ff3fb94c30f16c3aee275f62",
   "blobGasUsed":"0x0",
   "excessBlobGas":"0x4b20000",
   "parentBeaconBlockRoot":"0xffd644329d46413352298de6bde7ca71752862a2bae34f0e3ead6970a3e9912a",
   "requestsRoot":null,
   "hash":"0x6d7977fbf9333267c5bc25b596eacc2ef89461289e078dd7283c6872008646bc"
}
```
## 返回特定区块的区块头信息-BlockHeight
1. 提供的区块高度 `< 0  || == nil`，按照最新区块高度处理
2. 提供的区块高度 `> latestBlockHeight`，报错 `ethereum.NotFound`
```go
func checkError(err error) {
	if err != nil {
		log.Fatalf("error = %v", err)
	}
}
func main() {
	var ctx = context.Background()
	headers, err := getTargetBlockHeader(ctx, big.NewInt(-2))//nil is the latest block height
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println("invalid block height")
		} else {
			checkError(err)
		}
	} else {
		headerBytes, err := headers.MarshalJSON()
		if err == nil {
			fmt.Println(fmt.Sprintf("Target block header info = %s", string(headerBytes)))
			var header types.Header
			if err = header.UnmarshalJSON([]byte(string(headerBytes))); err != nil {
				log.Fatalf("decode error: %v ", err)
			}
			fmt.Println(header.Number)
		} else {
			checkError(err)
		}
	}
}

func getTargetBlockHeader(ctx context.Context, number *big.Int) (*types.Header, error) {
	header, err := client.HeaderByNumber(ctx, number)
	return header, err
}
```
## BlocksDataEncode
`rlp` 编码区块信息
```go
func main() {
	var ctx = context.Background()
	hash := "0x6d7977fbf9333267c5bc25b596eacc2ef89461289e078dd7283c6872008646bc"
	blockInfo := getBlock(ctx, common.HexToHash(hash))
	fmt.Println(fmt.Printf("Blocks data %s  within block hash = %s", blockInfo, hash))
	txs := decodeBlock(blockInfo)
	fmt.Println(txs)
}

func getBlock(ctx context.Context, hash common.Hash) string {
	block, err := client.BlockByHash(ctx, hash)
    if err != nil {
        if errors.Is(err, ethereum.NotFound) {
            return "", errors.New("non-exist blockHash")
        }else{
            checkError(err)
        }
    }
	BlockEnc, err := rlp.EncodeToBytes(&block)
	if err != nil {
		log.Fatalf("Encode blocks err = %v", err)
	}
	blockStr := common.Bytes2Hex(BlockEnc)
	return blockStr
}
func decodeBlock(blockInfo string) []string {
	blockEnc := common.FromHex(blockInfo)
	var block types.Block
	if err := rlp.DecodeBytes(blockEnc, &block); err != nil {
		log.Fatalf("decode error: %v", err)
	}
	tx := make([]string, len(block.Body().Transactions))
	for _, trans := range block.Body().Transactions {
		transB, _ := trans.MarshalJSON()
		tx = append(tx, string(transB))
	}
	return tx
	//check := func(f string, got, want interface{}) {
	//	if !reflect.DeepEqual(got, want) {
	//		log.Fatalf("%s mismatch: got %v, want %v", f, got, want)
	//	}
	//}
}
```
## BlocksReceipt
获取当前区块体中全部交易的收据信息，用于判断每条交易的执行状态
```go
func checkError(err error) {
	if err != nil {
		log.Fatalf("error = %v", err)
	}
}

func main() {
	var ctx = context.Background()
	hash := "0x6d7977fbf9333267c5bc25b596eacc2ef89461289e078dd7283c6872008646bc"
	var filter rpc.BlockNumberOrHash = rpc.BlockNumberOrHashWithHash(common.HexToHash(hash), false)
	count := getBlockReceipt(ctx, filter)
	fmt.Println(count)
}

func getBlockReceipt(ctx context.Context, filter rpc.BlockNumberOrHash) int {
	receipts, err := client.BlockReceipts(ctx, filter)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			fmt.Println("invalid block height")
		} else {
			checkError(err)
		}
	}
	return len(receipts)
}
```
