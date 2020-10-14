package model

type DBTableNames struct {
	Nebulas, Nodes, NodeRewards, NodeSocials, Datafeeds, CommonStats, NodeIPsMap string
}

var DefaultExtendedDBTableNames = &DBTableNames{
	Nebulas:     "nebulas",
	Nodes:       "nodes",
	NodeRewards: "node_rewards",
	NodeSocials: "node_socials",
	Datafeeds:   "data_feeds",
	CommonStats: "common_stats",
	NodeIPsMap:  "node_ips_map",
}