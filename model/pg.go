package model

type DBTableNames struct {
	Nebulas, Nodes, NodeRewards, NodeSocials, Datafeeds, CommonStats, NodeIPsMap, Transactions,
	CommitTable, RevealTable, AddOracleTable, AddOracleInNebulaTable, ResultTable, NewRoundTable,
	VoteTable, AddNebulaTable, DropNebulaTable, SignNewConsulsTable, SignNewOraclesTable, ApproveLastRoundTable string
}

var DefaultExtendedDBTableNames = &DBTableNames{
	Nebulas:                "nebulas",
	Nodes:                  "nodes",
	NodeRewards:            "node_rewards",
	NodeSocials:            "node_socials",
	Datafeeds:              "data_feeds",
	CommonStats:            "common_stats",
	NodeIPsMap:             "node_ips_map",
	Transactions:           "transactions",
	CommitTable:            "commit_transactions",
	RevealTable:            "reveal_transactions",
	AddOracleTable:         "add_oracle_transactions",
	AddOracleInNebulaTable: "add_oracle_in_nebula_transactions",
	ResultTable:            "result_transactions",
	NewRoundTable:          "new_round_transactions",
	VoteTable:              "vote_transactions",
	AddNebulaTable:         "add_nebula_transactions",
	DropNebulaTable:        "drop_nebula_transactions",
	SignNewConsulsTable:    "sign_new_consuls_transactions",
	SignNewOraclesTable:    "sign_new_oracles_transactions",
	ApproveLastRoundTable:  "approve_last_round_transactions",
}
