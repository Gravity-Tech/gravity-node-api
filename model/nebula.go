package model

// Target chain

type ChainType = int

const (
	WAVES_TARGET_CHAIN ChainType = iota
	ETH_TARGET_CHAIN
)

// Nebula states
const (
	PendingStatus uint = iota
	ActiveStatus
)

type Score = uint16

type Nebula struct {
	// fundamentals
	Name string `json:"name"`
	Status uint `json:"status"`
	Description string `json:"description"`
	Score Score `json:"score"`

	TargetChain ChainType `json:"target_chain"`

	SubscriptionFee uint64 `json:"subscription_fee"`

	// complex
	Extractor *IExtractor `json:"extractor"`
	NodesUsing []Node `json:"nodes_using"`
}