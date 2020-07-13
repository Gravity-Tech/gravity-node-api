package model

// Target chain

type ChainType = int
type SubscriptionFee = string

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
	TargetChain ChainType `json:"target_chain",pg:"target_chain"`

	SubscriptionFee SubscriptionFee `json:"subscription_fee",pg:"subscription_fee"`

	//Extractor *abstract.IExtractor `json:"extractor"`
	NodesUsing []Node `json:"nodes_using",pg:"nodes_using"`

	// Data feed subscription charge regularity
	// Represents minutes. For i.g. 1440 - one day
	//
	// required: true
	Regularity int64 `json:"regularity"`
}

func (nebula *Nebula) Matches (str string) bool {
	fieldValues := []string { nebula.Name, nebula.Description }

	return MatchStrList(fieldValues, str)
}
