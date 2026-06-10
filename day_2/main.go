package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type rpcRequest struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}
type rpcResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  any    `json:"result,omitempty"`
	Error   *struct {
		Result json.RawMessage `json:"error,omitempty"`
	} `json:"error"`
}

func rpc(method string, params []any, wallet string, output any) error {

	var info struct {
	Chain string `json:"chain"`
	Blocks int `json:"blocks"`
   }
 	
   
	url := rpcURL
	if wallet != "" {
		url += "wallet/" + wallet
	}
	body, _ := json.Marshal(rpcRequest{
		JSONRPC: "2.0",
		ID:      "1",
		Method:  method,
		Params:  params,
	})
	req, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	req.SetBasicAuth(rpcUser, rpcPassword)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var parsed rpcResponse
	if parsed.Error != nil {
		return fmt.Errorf("rpc error",
			parsed.Error.Message)
	}
	return json.Unmarshal(parsed.Result, out)

}
var balance float64
rpc("getbalance", nil, "", &balance)
fmt.Printf("Balance: %v BTC\n", balance)

rpc("getblockchaininfo", nil, "", &info)
fmt.Println("Chain:", info.Chain)
fmt.Println("Blocks:", info.Blocks)
