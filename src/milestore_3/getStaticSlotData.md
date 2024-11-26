# GetSlotData
合约数据全部按照[Solidity Slot 存储规则](https://yuhuajing.github.io/solidity-book/milestone_1/static-slot-storage.html)存储在区块链上，因此只要上链的数据就能通过 slot 键获取值

```go
func getSCstorage(address common.Address, slot int, blockNum int64) {
	t := common.BigToHash(big.NewInt(int64(slot)))
	int256 := new(big.Int)
	if blockNum != 0 {
		//fmt.Printf("get slot %d of the address %s in the block %d\n", slot, address.Hex(), blockNum)
		blocknumBigInt := big.NewInt(int64(blockNum))
		res, _ := client.StorageAt(context.Background(), address, t, blocknumBigInt)
		//	fmt.Println(res)
		int256.SetBytes(res)
	} else {
		//fmt.Printf("get slot %d of the address %s in the latest block\n", slot, address.Hex())
		res, _ := client.StorageAt(context.Background(), address, t, nil)
		//	fmt.Println(res)
		int256.SetBytes(res)
	}
	//fmt.Println()
	fmt.Printf("0x%x\n", int256)
	// fmt.Printf("uint256: %v\n", int256)
}

```
