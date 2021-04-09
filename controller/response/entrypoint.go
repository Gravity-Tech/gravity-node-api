package response

import (
	"github.com/Gravity-Tech/gravity-node-api/controller"
	"github.com/Gravity-Tech/gravity-node-api/router"
	"net/http"
)

type ResponseController struct {
	DBControllerDelegate *controller.DBController
}


func (rc *ResponseController) Handle () {
	http.HandleFunc(router.GetAllNebulas, rc.GetAllNebulas)
	http.HandleFunc(router.GetExactNebula, rc.GetExactNebula)

	http.HandleFunc(router.GetNodeRewards, rc.GetNodeRewardsList)
	http.HandleFunc(router.GetNodeActionsHistory, rc.GetNodeActionsHistory)
	http.HandleFunc(router.GetAllNodes, rc.GetAllNodes)
	http.HandleFunc(router.GetExactNode, rc.GetExactNode)

	http.HandleFunc(router.GetCommonStats, rc.GetCommonStats)
	http.HandleFunc(router.GetAvailableDataFeeds, rc.GetAvailableDatafeedsList)

	http.HandleFunc(router.GetAllTransactions, rc.GetAllTransactions)
	http.HandleFunc(router.GetExactTransaction, rc.GetExactTransaction)

	http.HandleFunc(router.GetAllSwaps, rc.GetAllSwaps)

	http.HandleFunc(router.GetAllCommit, rc.GetAllCommit)
	http.HandleFunc(router.GetExactCommit, rc.GetExactCommit)
	http.HandleFunc(router.GetAllReveal, rc.GetAllReveal)
	http.HandleFunc(router.GetExactReveal, rc.GetExactReveal)
	http.HandleFunc(router.GetAllAddOracle, rc.GetAllAddOracle)
	http.HandleFunc(router.GetExactAddOracle, rc.GetExactAddOracle)
	http.HandleFunc(router.GetAllAddOracleInNebula, rc.GetAllAddOracleInNebula)
	http.HandleFunc(router.GetExactAddOracleInNebula, rc.GetExactAddOracleInNebula)
	http.HandleFunc(router.GetAllResult, rc.GetAllResult)
	http.HandleFunc(router.GetExactResult, rc.GetExactResult)
	http.HandleFunc(router.GetAllNewRound, rc.GetAllNewRound)
	http.HandleFunc(router.GetExactNewRound, rc.GetExactNewRound)
	http.HandleFunc(router.GetAllVote, rc.GetAllVote)
	http.HandleFunc(router.GetExactVote, rc.GetExactVote)
	http.HandleFunc(router.GetAllAddNebula, rc.GetAllAddNebula)
	http.HandleFunc(router.GetExactAddNebula, rc.GetExactAddNebula)
	http.HandleFunc(router.GetAllDropNebula, rc.GetAllDropNebula)
	http.HandleFunc(router.GetExactDropNebula, rc.GetExactDropNebula)
	http.HandleFunc(router.GetAllSignNewConsuls, rc.GetAllSignNewConsuls)
	http.HandleFunc(router.GetExactSignNewConsuls, rc.GetExactSignNewConsuls)
	http.HandleFunc(router.GetAllSignNewOracles, rc.GetAllSignNewOracles)
	http.HandleFunc(router.GetExactSignNewOracles, rc.GetExactSignNewOracles)
	http.HandleFunc(router.GetAllApproveLastRound, rc.GetAllApproveLastRound)
	http.HandleFunc(router.GetExactApproveLastRound, rc.GetExactApproveLastRound)
}
