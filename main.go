package main

import (
	"log"
	"net/http"

	"MaoDaGreith/CryptoAssignment/controller"
	"MaoDaGreith/CryptoAssignment/testFunc"
)

func main() {
	go testFunc.StartPoll() // Start background polling for new blocks

	// Define HTTP routes
	http.HandleFunc("/currentblock", controller.GetCurrentBlockHandler)
	http.HandleFunc("/subscribe", controller.SubscribeHandler)
	http.HandleFunc("/transactions", controller.GetTransactionsHandler)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
