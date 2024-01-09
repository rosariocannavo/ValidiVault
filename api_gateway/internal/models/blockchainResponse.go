package models

type BlockChainResponse struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
}
