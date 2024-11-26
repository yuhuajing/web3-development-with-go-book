# New account
## ECDSA 生成新的地址
要首先生成一个新的钱包，需要导入 `go-ethereum crypto` 包，该包提供用于生成随机私钥的 `GenerateKey` 方法
1. 生成全新的私钥，并将私钥存储在本地文件
2. 基于私钥导出公钥信息
3. 基于公钥导出账户地址
```go
package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strings"
)

func NewAccount() common.Address {
	ecdsaPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	newAddress := EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	err = crypto.SaveECDSA(newAddress.Hex()+"_key.txt", ecdsaPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	return newAddress
}

func AddressFromKey(key string) common.Address {
	if strings.HasPrefix(key, "0x") {
		key = strings.Trim(key, "0x")
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}
	rAddress := EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	return rAddress
}

func EcdsaAddressFromPrivateKey(ecdsaPrivateKey *ecdsa.PrivateKey) common.Address {
	publicKeyBytes := crypto.FromECDSAPub(ecdsaPrivateKey.Public().(*ecdsa.PublicKey))
	pub, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	rAddress := crypto.PubkeyToAddress(*pub)
	return rAddress
}
```
## 生成助记词，从助记词导出新地址
地址胜场方式：
1. 生成助记词
2. 拼接助记词和 `secret`， 构造哈希种子 `seed`
3. 基于 `seed` 逐步生成 `ecdsa` 私钥
4. 基于私钥导出钱包地址

导出账户密钥
1. 基于助记词导出密钥时：
- 必须提供相匹配的 `secret` 值
- 拼接出生成账户用的哈希种子 `seed`
- 基于 `seed` 产生相同的私钥和地址
- 因此，`secret` 用于进一步保证助记词的安全
```go
func main() {
    secret := "ert"
    mnemonic, address := KeyFromMnemonic(secret)
    fmt.Println(mnemonic, address)
    KeyFromMnemonicInput(mnemonic, secret)
}
func KeyFromMnemonic(secret string) (string, common.Address) {
	// Generate a mnemonic
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied passphrase
	seed := bip39.NewSeed(mnemonic, secret)
	masterPrivateKey, _ := bip32.NewMasterKey(seed)
	ecdsaPrivateKey := crypto.ToECDSAUnsafe(masterPrivateKey.Key)
	address := EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	return mnemonic, address
}

func KeyFromMnemonicInput(mnemonic, secret string) string {
	seed := bip39.NewSeed(mnemonic, secret)
	masterPrivateKey, _ := bip32.NewMasterKey(seed)
	ecdsaPrivateKey := crypto.ToECDSAUnsafe(masterPrivateKey.Key)
	privateKeyHex := fmt.Sprintf("%x", ecdsaPrivateKey.D)
	address := EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	fmt.Println(address)
	return privateKeyHex
}
```
## 从Keystore导出地址
1. 创建新的钱包地址，账户信息通过 `secret` 加密存储并导出 `keystore` 文件
2. 基于 `secret` 解密 `keystore` 文件

私钥签名
1. `keystore` 直接基于 `secret` 签名数据
2. `keystore` 可以先基于 secret `TimedUnlock` 解锁一段时间后可以直接用于签名
```go
import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"os"
)

func NewKeystoreAccount(secret string) common.Address {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(secret)
	if err != nil {
		log.Fatal(err)
	}
	return account.Address
}

func AddressFromKeystore(file, secret string) common.Address {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	account, err := ks.Import(jsonBytes, secret, secret)
	if err != nil {
		log.Fatal(err)
	}
	return account.Address
}

func SignatureFromKeystore(file, secret string) string {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	account, err := ks.Import(jsonBytes, secret, secret)
	if err != nil {
		log.Fatal(err)
	}
	signatureBytes, err := ks.SignHashWithPassphrase(account, secret, []byte("ww"))
	return hexutil.Encode(signatureBytes)
}
```
