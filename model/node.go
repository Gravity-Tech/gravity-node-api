package model

import "time"

type NodeActionType = int

const (
	NodeDataSent NodeActionType = iota
	NodeConsensusReached
	NodeAssetsReceived
	NodeVoteReceived
)

type NodeSocials struct {
	Telegram string  `json:"tg"`
	Facebook string  `json:"fb"`
	Linked string  `json:"linkedin"`
	Twitter string  `json:"twitter"`
}

type NodeContacts struct {
	Website string  `json:"website"`
}

type Node struct {
	Name string  `json:"name"`
	Description string  `json:"description"`
	Score Score `json:"score"`

	DepositChain ChainType `json:"deposit_chain"`
	DepositAmount int64 `json:"deposit_amount"`

	JoinedAt int64 `json:"joined_at"`
	LockedUntil int64 `json:"locked_until"`
	// JoinedAt > LockedUntil - node is active

	NebulasUsing []Nebula `json:"nebulas_using"`

	Сontacts NodeContacts `json:"contacts"`
	Socials NodeSocials `json:"socials"`
}

type NodeHistoryRecord struct {
	Name string  `json:"name"`
	Type NodeActionType `json:"type"`
	Asset Asset `json:"asset"`
	Status CommonStatus `json:"status"`
	Timestamp time.Time `json:"time"`
}