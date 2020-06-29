package model

// swagger:model
type NodeReward struct {
	Amount uint64 `json:"amount"`
	Decimals int `json:"decimals"`
	Timestamp int64 `json:"timestamp"`
	Currency string `json:"currency"`
}

