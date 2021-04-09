package router

const (
	// Nebula routes
	GetAllNebulas = "/nebulas/all"
	GetExactNebula  = "/nebulas/exact"

	// Node Routes
	GetNodeRewards        = "/nodes/rewards/all"
	GetNodeActionsHistory = "/nodes/actions/history"
	GetAllNodes   = "/nodes/all"
	GetExactNode  = "/nodes/exact"

	// Common stats
	GetCommonStats        = "/common/stats"

	// Data feeds
	GetAvailableDataFeeds = "/data/datafeeds/all"

	// For extractors
	// GetAvailableExtractors = "/data/extractors/all"

	GetAllTransactions  = "/transactions/all"
	GetExactTransaction = "/transactions/exact"

	GetAllSwaps = "/transactions/swap/all"

	GetAllCommit   = "/transactions/commit/all"
	GetExactCommit = "/transactions/commit/exact"

	GetAllReveal   = "/transactions/reveal/all"
	GetExactReveal = "/transactions/reveal/exact"

	GetAllAddOracle   = "/transactions/add-oracle/all"
	GetExactAddOracle = "/transactions/add-oracle/exact"

	GetAllAddOracleInNebula   = "/transactions/add-oracle-in-nebula/all"
	GetExactAddOracleInNebula = "/transactions/add-oracle-in-nebula/exact"

	GetAllResult   = "/transactions/result/all"
	GetExactResult = "/transactions/result/exact"

	GetAllNewRound   = "/transactions/new-round/all"
	GetExactNewRound = "/transactions/new-round/exact"

	GetAllVote   = "/transactions/vote/all"
	GetExactVote = "/transactions/vote/exact"

	GetAllAddNebula   = "/transactions/add-nebula/all"
	GetExactAddNebula = "/transactions/add-nebula/exact"

	GetAllDropNebula   = "/transactions/drop-nebula/all"
	GetExactDropNebula = "/transactions/drop-nebula/exact"

	GetAllSignNewConsuls   = "/transactions/sign-new-consuls/all"
	GetExactSignNewConsuls = "/transactions/sign-new-consuls/exact"

	GetAllSignNewOracles   = "/transactions/sign-new-oracles/all"
	GetExactSignNewOracles = "/transactions/sign-new-oracles/exact"

	GetAllApproveLastRound   = "/transactions/approve-last-round/all"
	GetExactApproveLastRound = "/transactions/approve-last-round/exact"
)
