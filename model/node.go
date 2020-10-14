package model

import "github.com/Gravity-Tech/gravity-core/config"

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

type NodeContactInfo struct {
	NodeSocials
	NodeContacts
}

// swagger:model
type Node struct {
	// the address of the node
	//
	// required: true
	Address string `json:"address"`

	PublicKey string `json:"public_key"`

	Name string  `json:"name"`
	Description string  `json:"description"`
	Score Score `json:"score"`

	DepositChain ChainType `json:"deposit_chain",pg:"deposit_chain"`
	DepositAmount int64 `json:"deposit_amount",pg:"deposit_amount"`

	JoinedAt int64 `json:"joined_at",pg:"joined_at"`
	LockedUntil int64 `json:"locked_until",pg:"locked_until"` // JoinedAt > LockedUntil - node is active

	//NebulasUsing []Nebula `json:"nebulas_using",pg:"-"`
	NebulasUsing []string `pg:",array" json:"nebulas_using"`

	//Contacts NodeContacts `json:"contacts"`
	//Socials NodeSocials `json:"socials"`
}

func (node *Node) AddNebulas (nebulas ...Nebula) {
	node.NebulasUsing = []string {}

	for _, nebula := range nebulas {
		//node.NebulasUsing = append(node.NebulasUsing, nebula)
		node.NebulasUsing = append(node.NebulasUsing, nebula.Address)
	}
}

func (node *Node) Matches (str string) bool {
	fieldValues := []string { node.Name, node.Description }

	return MatchStrList(fieldValues, str)
}

func (node *Node) UpdateByValidatorDetails (details *config.ValidatorDetails) {
	node.Name = details.Name
	node.Description = details.Description
}


type NodeIPMapRecord struct {
	IPAddress string `json:"ip_address"`
	PublicKey string `json:"public_key"`
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

