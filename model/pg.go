package model

type DBTableNames struct {
	Nebulas, Nodes, NodeRewards, NodeSocials string
}

var DefaultExtendedDBTableNames = &DBTableNames{
	Nebulas:     "nebulas",
	Nodes:       "nodes",
	NodeRewards: "node_rewards",
	NodeSocials: "node_socials",
}