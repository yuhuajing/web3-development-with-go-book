



## 获取链上GasPrice
```go
func SuggestedGasPrice(rpcUrl string) {
	// Retrieve the currently suggested gas price for a new transaction.
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// Print the suggested gas price to the terminal.
	fmt.Println("Suggested Gas Price:", gasPrice.String())
}
```
## 获取链上预估gas
```go
func EstimateGas(rpcUrl, from, to, data string, value uint64) uint64 {
	var ctx = context.Background()
	var err error
	var (
		fromAddr  = common.HexToAddress(from)     // Convert the from address from hex to an Ethereum address.
		toAddr    = common.HexToAddress(to)       // Convert the to address from hex to an Ethereum address.
		amount    = new(big.Int).SetUint64(value) // Convert the value from uint64 to *big.Int.
		bytesData []byte
	)

	// Encode the data if it's not already hex-encoded.
	if data != "" {
		if ok := strings.HasPrefix(data, "0x"); !ok {
			data = hexutil.Encode([]byte(data))
		}

		bytesData, err = hexutil.Decode(data)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Create a message which contains information about the transaction.
	msg := ethereum.CallMsg{
		From:  fromAddr,
		To:    &toAddr,
		Gas:   0x00,
		Value: amount,
		Data:  bytesData,
	}

	// Estimate the gas required for the transaction.
	gas, err := client.EstimateGas(ctx, msg)
	if err != nil {
		log.Fatalln(err)
	}
	return gas
}
```


// createRawTransaction creates a raw EIP-1559 transaction and returns it as a hex string.
func CreateRawTransaction(rpcURL, to, data, privKey string, gasLimit, wei uint64) string {
// Retrieve the chain ID for the target Ethereum network.
chainID, err := client.ChainID(context.Background())
if err != nil {
log.Fatalln(err)
}

	// Suggest the base fee for inclusion in a block.
	baseFee, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Suggest a gas tip cap (priority fee) for miner incentive.
	priorityFee, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Calculate the maximum gas fee cap, adding a 2 GWei margin to the base fee plus priority fee.
	increment := new(big.Int).Mul(big.NewInt(2), big.NewInt(params.GWei))
	gasFeeCap := new(big.Int).Add(baseFee, increment)
	gasFeeCap.Add(gasFeeCap, priorityFee)

	// Decode the provided private key.
	if ok := strings.HasPrefix(privKey, "0x"); !ok {
		privKey = hexutil.Encode([]byte(privKey))
	}

	pKeyBytes, err := hexutil.Decode(privKey)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert the private key bytes to an ECDSA private key.
	ecdsaPrivateKey, err := crypto.ToECDSA(pKeyBytes)
	if err != nil {
		log.Fatalln(err)
	}

	// Extract the public key from the ECDSA private key.
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
	amount := new(big.Int).SetUint64(wei)
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

	// Create a new transaction object from the prepared data.
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

	return rawTxRLPHex
}

