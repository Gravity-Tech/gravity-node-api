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

}