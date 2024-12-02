# MerkleProof
简介链接：

[https://yuhuajing.github.io/solidity-book/milestone_6/merkle-proof-validation.html](https://yuhuajing.github.io/solidity-book/milestone_6/merkle-proof-validation.html)

[https://github.com/yuhuajing/solidity-book/tree/main/src/ContractsHub/merkle_tree_prove](https://github.com/yuhuajing/solidity-book/tree/main/src/ContractsHub/merkle_tree_prove)

## merkleSecurity
为保证默克尔树的安全性，防止攻击者仅传递非叶子节点的数据跳过校验

叶子节点和验证节点的采用不用的 `hash` 校验：

## 叶子节点双 hash
待验证的叶子节点采用双 `hash` 的方式，区别叶子节点和验证节点的数据
```go
func abiPackLeafHash(leafEncodings []string, values ...interface{}) ([]byte, error) {
	data, err := AbiPack(leafEncodings, values...)
	if err != nil {
		return nil, err
	}
	hash, err := standardLeafHash(data)
	return hash, err
}

func standardLeafHash(value []byte) ([]byte, error) {
	k1, err := Keccak256(value)
	if err != nil {
		return nil, err
	}
	k2, err := Keccak256(k1)
	return k2, err
}
```
## 单类型叶子节点
叶子节点仅采用单一数据类型,匹配的验证合约能够获取的数据有限
```go
func MerkleOnlyOneArgOZ() {
	leaf1 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
	}

	leaf2 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
	}

	leaf3 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
	}

	leaf4 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
	}
	leaf5 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
	}

	leaves := [][]interface{}{
		leaf1,
		leaf2,
		leaf3,
		leaf4,
		leaf5,
	}

	tree, err := smt.Of(
		leaves,
		[]string{
			smt.SOL_ADDRESS,
			//smt.SOL_UINT256,
		})

	if err != nil {
		fmt.Println("Of ERR", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("Merkle Root: ", root)

	proof, err := tree.GetProof(leaf1)
	strProof := make([]string, len(proof))
	if err != nil {
		fmt.Println("GetProof ERR", err)
	}
	for _, v := range proof {
		strProof = append(strProof, hexutil.Encode(v))
	}
	fmt.Println("02 proof: ", strProof)
}
```
## 多类型叶子节点
叶子节点拼接数据类型,可以将各种类型拼接（各种类型、类型数组等），匹配的验证合约能够获取足额数据，进行额外数据的处理

示例代码将地址和书来给你拼接，合约中通过额外的 `mapping` 记录，可以验证当前地址的剩余额度，
```go
func MerkleWithMultiArgOZ() {
	leaf1 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
		smt.SolNumber("5000000000000000000"),
	}

	leaf2 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
		smt.SolNumber("2500000000000000000"),
	}

	leaf3 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
		smt.SolNumber("5000000000000000000"),
	}

	leaf4 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
		smt.SolNumber("2500000000000000000"),
	}
	leaf5 := []interface{}{
		smt.SolAddress("0x0000000000000000000000000000000000000000"),
		smt.SolNumber("2500000000000000000"),
	}

	leaves := [][]interface{}{
		leaf1,
		leaf2,
		leaf3,
		leaf4,
		leaf5,
	}

	tree, err := smt.Of(
		leaves,
		[]string{
			smt.SOL_ADDRESS,
			smt.SOL_UINT256,
		})

	if err != nil {
		fmt.Println("Of ERR", err)
	}

	root := hexutil.Encode(tree.GetRoot())
	fmt.Println("Merkle Root: ", root)

	proof, err := tree.GetProof(leaf1)
	strProof := make([]string, len(proof))
	if err != nil {
		fmt.Println("GetProof ERR", err)
	}
	for _, v := range proof {
		strProof = append(strProof, hexutil.Encode(v))
	}
	fmt.Println("02 proof: ", strProof)
}
```
