package controller

import (
	"encoding/json"
	"net/http"

	"MaoDaGreith/CryptoAssignment/parser"
)

// GetCurrentBlockHandler handler
func GetCurrentBlockHandler(w http.ResponseWriter, r *http.Request) {
	currentBlock := parser.GetCurrentBlock()
	json.NewEncoder(w).Encode(map[string]int{"current_block": currentBlock})
}

// SubscribeHandler handler
func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}
	subscribed := parser.Subscribe(address)
	json.NewEncoder(w).Encode(map[string]bool{"subscribed": subscribed})
}

// GetTransactionsHandler handler
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}
	transactions := parser.GetTransactions(address)
	json.NewEncoder(w).Encode(transactions)
}
