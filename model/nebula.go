package model

// Target chain

type ChainType = int

const (
	WAVES_TARGET_CHAIN ChainType = iota
	ETH_TARGET_CHAIN
)

// Nebula states
const (
	NebulaPendingStatus uint = iota
	NebulaActiveStatus
)

type Score = uint16

// swagger:model
type Nebula struct {

	// the name of the nebula
	//
	// required: true
	Name string `json:"name"`

	// the status of the nebula
	//
	// required: true
	Status uint `json:"status"`

	// the description of the nebula
	//
	// required: true
	Description string `json:"description"`

	// the score of the nebula
	//
	// required: true
	Score Score `json:"score"`

	// the target chain of the nebula
	// recently allowed: WAVES, ETH
	//
	// required: true
	TargetChain ChainType `json:"target_chain"`

	SubscriptionFee uint64 `json:"subscription_fee"`

	Extractor *IExtractor `json:"extractor"`
	NodesUsing []Node `json:"nodes_using"`
}