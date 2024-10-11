package testFunc

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"MaoDaGreith/CryptoAssignment/client"
	"MaoDaGreith/CryptoAssignment/data"
	"MaoDaGreith/CryptoAssignment/model"
)

// StartPoll starts the info getting
func StartPoll() {
	for {
		latestBlock := GetCurrentBlockFromBlockchain()
		if latestBlock > data.CurrentBlock {
			fetchBlockTransactions(latestBlock)
			data.CurrentBlock = latestBlock
		}
		time.Sleep(10 * time.Second) // Poll every 10 seconds
	}
}

// GetCurrentBlockFromBlockchain gets the current block
func GetCurrentBlockFromBlockchain() int {
	result, _ := client.RPC("eth_blockNumber", []interface{}{})
	blockNumberHex := result["result"].(string)
	blockNumber := hexToDec(blockNumberHex)
	return blockNumber
}

func fetchBlockTransactions(blockNumber int) {
	params := []interface{}{fmt.Sprintf("0x%x", blockNumber), true}
	result, _ := client.RPC("eth_getBlockByNumber", params)

	transactions := result["result"].(map[string]interface{})["transactions"].([]interface{})

	for _, tx := range transactions {
		processTransaction(tx.(map[string]interface{}))
	}
}

func processTransaction(txData map[string]interface{}) {
	from := txData["from"].(string)
	to := txData["to"].(string)

	tx := model.Transaction{
		From:  from,
		To:    to,
		Value: txData["value"].(string),
		Hash:  txData["hash"].(string),
	}

	if data.SubscribedAddresses[from] {
		data.Transactions[from] = append(data.Transactions[from], tx)
	}
	if data.SubscribedAddresses[to] {
		data.Transactions[to] = append(data.Transactions[to], tx)
	}
}

// hexToDec converts a hexadecimal string to a decimal integer
func hexToDec(hexStr string) int {
	// Remove the "0x" prefix if present
	hexStr = strings.TrimPrefix(hexStr, "0x")

	// Parse the hex string to a decimal integer
	dec, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return 0 // Handle error properly in real-world scenarios
	}

	return int(dec)
}
