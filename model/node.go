package model


// NodeActionType
type NodeActionType = int

const (
	NodeDataSent NodeActionType = iota
	NodeConsensusReached
	NodeAssetsReceived
	NodeVoteReceived
)

// swagger:model
type NodeSocials struct {
	Telegram string  `json:"tg"`
	Facebook string  `json:"fb"`
	Linked string  `json:"linkedin"`
	Twitter string  `json:"twitter"`
}

// swagger:model
type NodeContacts struct {
	Website string  `json:"website"`
}

// swagger:model
type Node struct {
	Name string  `json:"name"`
	Description string  `json:"description"`
	Score Score `json:"score"`

	DepositChain ChainType `json:"deposit_chain"`
	DepositAmount int64 `json:"deposit_amount"`

	JoinedAt int64 `json:"joined_at"`
	LockedUntil int64 `json:"locked_until"` // JoinedAt > LockedUntil - node is active

	NebulasUsing []Nebula `json:"nebulas_using"`

	Сontacts NodeContacts `json:"contacts"`
	Socials NodeSocials `json:"socials"`
}

func (node *Node) Matches (str string) bool {
	fieldValues := []string { node.Name, node.Description }

	return MatchStrList(fieldValues, str)
}


// swagger:model
type NodeHistoryRecord struct {
	Name string  `json:"name"`
	Type NodeActionType `json:"type"`
	Asset Asset `json:"asset"`
	Status CommonStatus `json:"status"`
	Timestamp int64 `json:"time"`
}

func (record *NodeHistoryRecord) Matches (str string) bool {
	fieldValues := []string { record.Name, record.Asset.Name }

	return MatchStrList(fieldValues, str)
}
