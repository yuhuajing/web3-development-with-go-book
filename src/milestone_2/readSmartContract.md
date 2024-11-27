# ReadContract
## solc
合约编译时会编译成两部分： 字节码(`bin`) + 合约函数的二进制接口(`abi`)
- `bin` 跟随交易部署到合约地址
- `abi` 文件用来调用合约函数

[Solc](https://docs.soliditylang.org/en/latest/installing-solidity.html)工具用于编译合约
- `solc --abi xx.sol`, 生成合约 `abi` 文件
- `solc --bin xx.sol`,生成合约的 `bin` 文件

[abigen](https://github.com/ethereum/go-ethereum/tree/master/cmd/abigen)工具可以基于 `abi` 和 `bin` 文件创建合约部署和调用的 `Golang` 文件
- `abigen --abi=xxx.abi --bin=xxx.bin --pkg=xxx --out=xxx.go`
  - 仅仅提供 `abi` 文件的话，只能生成合约的调用文件，不能部署合约

## 代码
```go
package milestone2

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strconv"
)

func StakingInfo(contract, nft, nftId string) (common.Address, int64, error) {
	var staker common.Address
	stakingEndTime := int64(0)
	address := common.HexToAddress(contract)
	instance, err := NewStakeCaller(address, client)
	if err != nil {
		return staker, stakingEndTime, err
	}
	nftid, err := strconv.ParseInt(nftId, 10, 64)
	if err != nil {
		return staker, stakingEndTime, err
	}
	staker, endts, err := instance.RegisterData(nil, common.HexToAddress(nft), big.NewInt(nftid))
	if err != nil {
		return staker, stakingEndTime, err
	}
	return staker, endts.Int64(), nil
}
```
