package model


type NodeSocials struct {
	Telegram string  `json:"tg"`
	Facebook string  `json:"fb"`
	Linked string  `json:"linkedin"`
	Twitter string  `json:"twitter"`
}

type Node struct {
	Name string  `json:"name"`
	Description string  `json:"description"`
	Score Score `json:"score"`

	DepositChain ChainType `json:"deposit_chain"`
	DepositAmount int64 `json:"deposit_amount"`

	JoinedAt int64 `json:"joined_at"`

	NebulasUsing []Nebula `json:"nebulas_using"`

	Socials NodeSocials `json:"socials"`
}