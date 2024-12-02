# Signature
简介：

[https://yuhuajing.github.io/solidity-book/milestone_6/signature-ECDSA-validation.html](https://yuhuajing.github.io/solidity-book/milestone_6/signature-ECDSA-validation.html)

## sign_sha256
预编译合约地址 0x2 [https://yuhuajing.github.io/solidity-book/milestone_5/contracts-precompile.html](https://yuhuajing.github.io/solidity-book/milestone_5/contracts-precompile.html)实现 `sha256` 哈希校验
```go
package sign

import (
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func Sha256EncodeNumber(number int64) {
	s := hexutil.EncodeBig(big.NewInt(number))
	prefix := ""
	num := 64 - len(s[2:])
	for index := 0; index < num; index++ {
		prefix += "0"
	}
	s = s[:2] + prefix + s[2:]
	byteD, err := hexutil.Decode(s)
	if err != nil {
		fmt.Println(err)
	}
	h := sha256.New()
	h.Write(byteD)
	bs := h.Sum(nil)
	fmt.Println(hexutil.Encode(bs))
}
```

## sign_ecdsa
1. 编码 `hash` 待签名数据
```go
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message)))
	messageBytes := milestone4.UnsafeBytes(message)

	// Hash the prefix and message using Keccak-256
	hash := crypto.Keccak256Hash(prefix, messageBytes)
```
编码多种类型的代签数据,基于开源库`solsha3 "github.com/miguelmota/go-solidity-sha3"`
```go
	mes := solsha3.SoliditySHA3(
		[]string{"uint32[]", "uint32", "uint64", "uint64", "uint64", "address", "address", "address"},
		[]interface{}{
			nftId,
			chainId,
			timestamp,
			uuid,
			signId,
			nft,
			sender,
			contract,
		},
	)
	hash := solsha3.SoliditySHA3WithPrefix(mes)
```
2. 私钥签名数据hash
```go
	// Sign the hashed message
	sig, err := crypto.Sign(hash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		log.Fatalln(err)
	}

	// Adjust signature ID to Ethereum's format
	sig[64] += 27
```
3. 校验数据，根据待签数据和签名逆向反推签名地址
- 签名不匹配的话，会返回错误的签名地址
```go
func VerifySig(signature, address, message string) bool {
	// Decode the signature into bytes
	sig, err := hexutil.Decode(signature)
	if err != nil {
		log.Fatalln(err)
	}

	// Adjust signature to standard format (remove Ethereum's recovery ID)
	sig[64] = sig[64] - 27

	// Construct the message prefix
	prefix := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message)))
	data := []byte(message)

	// Hash the prefix and data using Keccak-256
	hash := crypto.Keccak256Hash(prefix, data)

	// Recover the public key bytes from the signature
	sigPublicKeyBytes, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		log.Fatalln(err)
	}
	ecdsaPublicKey, err := crypto.UnmarshalPubkey(sigPublicKeyBytes)
	if err != nil {
		log.Fatalln(err)
	}

	// Derive the address from the recovered public key
	rAddress := crypto.PubkeyToAddress(*ecdsaPublicKey)

	// Check if the recovered address matches the provided address
	isSigner := strings.EqualFold(rAddress.String(), address)

	return isSigner
}
```
