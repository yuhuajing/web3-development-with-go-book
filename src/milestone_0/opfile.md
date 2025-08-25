# OperateFile
## ReadCSV
```go
	file, err := os.Open("file.csv")
	if err != nil {
	}

	reader := csv.NewReader(file)
	reader.Comma = ','

	lines, err := reader.ReadAll()
	if err != nil {
		
	}

	for index, line := range lines {}
```
## WriteCSV
```go
func ParseTxToCSV(res *[]gettxbyaddress.Result, apiKey string) {
	file, err := os.OpenFile("file.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal()
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, txRes := range *res {
		ERC20 := checkERC20(txRes.To, apiKey)
		err := writer.Write([]string{
			txRes.BlockNumber,
			txRes.TimeStamp,
			txRes.Hash,
			txRes.Nonce,
			txRes.BlockHash,
			txRes.TransactionIndex,
			txRes.From,
			txRes.To,
			ERC20,
			txRes.Value,
			txRes.Gas,
			txRes.GasPrice,
			txRes.IsError,
			txRes.TxReceiptStatus,
			txRes.Input,
			txRes.ContractAddress,
			txRes.CumulativeGasUsed,
			txRes.GasUsed,
			txRes.Confirmations,
			txRes.MethodId,
			txRes.FunctionName,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
```
## HttpWithSameSession
```go
func createHttpClient() (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("cookiejar.New error %s", err.Error())
	}
	return &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			DisableKeepAlives:   true,
			MaxIdleConnsPerHost: -1,
		},
		Timeout: 0,
	}, nil
}
```

