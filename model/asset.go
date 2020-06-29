package model

// swagger:model
type Asset struct {
	Name string `json:"name"`
	Decimals int `json:"decimals"`
	Amount string `json:"amount"`
}