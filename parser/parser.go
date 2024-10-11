package parser

import (
	"MaoDaGreith/CryptoAssignment/data"
	"MaoDaGreith/CryptoAssignment/model"
)

// Subscribe adds address
func Subscribe(address string) bool {
	if _, exists := data.SubscribedAddresses[address]; exists {
		return false // Address already subscribed
	}
	data.SubscribedAddresses[address] = true
	return true
}

// GetCurrentBlock returns last block
func GetCurrentBlock() int {
	return data.CurrentBlock
}

// GetTransactions returns
func GetTransactions(address string) []model.Transaction {
	return data.Transactions[address]
}
