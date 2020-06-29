package model

// swagger:model
type NodeReward struct {
	Amount uint64 `json:"amount"`
	Decimals int `json:"decimals"`
	Timestamp int64 `json:"timestamp"`
	Currency string `json:"currency"`
}

func (reward *NodeReward) Matches (str string) bool {
	fieldValues := []string { reward.Currency }

	return MatchStrList(fieldValues, str)
}
