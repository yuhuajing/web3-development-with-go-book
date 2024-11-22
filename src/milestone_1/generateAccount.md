# New account
## ECDSA 生成新的地址
```go
package generateAccount

import (
	"crypto/ecdsa"
	"errors"
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
	//k := hex.EncodeToString(crypto.FromECDSA(ecdsaPrivateKey))
	//fmt.Println(k)

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(errors.New("Export ecdsaPublicKey error"))
	}
	newAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
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
1. 生成助记词
2. 拼接 助记词和 secret， 构造哈希种子
3. 基于哈希种子逐步生成 ecdsa 私钥
4. 基于 ecdsa 私钥导出钱包地址
5. 基于 助记词导出 ecdsa 密钥时：
- 必须提供相匹配的 secret 值
- 执行相同的拼接和计算，产生相同的私钥和地址
- 因此，secret 用于进一步保证助记词的安全
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
	//fmt.Println("Mnemonic (gen): ", mnemonic)
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied passphrase
	seed := bip39.NewSeed(mnemonic, secret)
	masterPrivateKey, _ := bip32.NewMasterKey(seed)
	ecdsaPrivateKey := crypto.ToECDSAUnsafe(masterPrivateKey.Key)
	address := EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	return mnemonic, address
}

func KeyFromMnemonicInput(mnemonic, secret string) string {
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied passphrase
	seed := bip39.NewSeed(mnemonic, secret)
	masterPrivateKey, _ := bip32.NewMasterKey(seed)
	ecdsaPrivateKey := crypto.ToECDSAUnsafe(masterPrivateKey.Key)
	privateKeyHex := fmt.Sprintf("%x", ecdsaPrivateKey.D)
	address := EcdsaAddressFromPrivateKey(ecdsaPrivateKey)
	fmt.Println(address)
	return privateKeyHex
}
```
## 生成Keystore，从Keystore导出地址
1. 创建新的钱包地址，账户信息通过 secret 加密存储并导出 keystore 文件
2. 基于 secret 解密 keystore 文件
3. keystore 直接基于 secret 签名数据
4. keystore 可以先基于 secret `TimedUnlock` 解锁一段时间后可以直接用于签名
```go
package generateAccount

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
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
	jsonBytes, err := ioutil.ReadFile(file)
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
	jsonBytes, err := ioutil.ReadFile(file)
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
