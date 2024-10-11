package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RPCRequest is the structure for JSON-RPC request
type RPCRequest struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// RPC sends/received request/result
func RPC(method string, params []interface{}) (map[string]interface{}, error) {
	url := "https://ethereum-rpc.publicnode.com"
	requestBody, _ := json.Marshal(RPCRequest{
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}
