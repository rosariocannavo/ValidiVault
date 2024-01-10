package models

var VerifySignatureResponse struct {
	Nonce     string `json:"message"`
	Address   string `json:"address"`
	Signature string `json:"signature"`
}
