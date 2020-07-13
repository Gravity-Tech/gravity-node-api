package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/model"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/utils"
	"net/http"
)

type ResponseController struct {
	DBControllerDelegate *DBController
}

func addBaseHeaders (headers http.Header) {
	headers.Add("Content-Type", "application/json")
}

// swagger:route GET /nebulas/all Nebulas getAllNebulas
//
// Returns all available gravity nebulas
//
// This will show all available gravity nebulas
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//       api_key:
//
//     Responses:
//       200: []Nebula
func (rc *ResponseController) GetAllNebulas(w http.ResponseWriter, req *http.Request) {
	nebulasList := rc.DBControllerDelegate.AllNebulasList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	for i, nebula := range *nebulasList {
		if queryString == "" { break }

		if !nebula.Matches(queryString) {
			*nebulasList = append((*nebulasList)[:i], (*nebulasList)[i+1:]...)
		}
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1

	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*nebulasList), itemsPerPage, pageIndex)
	*nebulasList = (*nebulasList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(nebulasList)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}


// swagger:route GET /nodes/all Nodes getAllNodes
//
// Returns all available gravity nodes
//
// This will show all available gravity nodes
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//       api_key:
//
//     Responses:
//       200: []Node
func (rc *ResponseController) GetAllNodes(w http.ResponseWriter, req *http.Request) {
	nodeList := rc.DBControllerDelegate.AllNodesList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	for i, node := range *nodeList {
		if queryString == "" { break }

		if !node.Matches(queryString) {
			*nodeList = append((*nodeList)[:i], (*nodeList)[i+1:]...)
		}
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*nodeList), itemsPerPage, pageIndex)
	*nodeList = (*nodeList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(nodeList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}


// swagger:route GET /common/stats Common getCommonStats
//
// Returns gravity node common statistics
//
// This will show all gravity common stats
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//       api_key:
//
//     Responses:
//       200: CommonStats
func (rc *ResponseController) GetCommonStats(w http.ResponseWriter, req *http.Request) {
	stats := utils.GetCommonStatsMockup()

	bytes, _ := json.Marshal(stats)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}

// swagger:route GET /nodes/rewards/all Nodes getNodeRewards
//
// Returns gravity node mockup rewards
//
// This will show all gravity node mockup rewards list
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//       api_key:
//
//     Responses:
//       200: []NodeReward
func (rc *ResponseController) GetNodeRewardsList(w http.ResponseWriter, req *http.Request) {
	rewards := utils.GetNodeRewardsListMockup()
	result := make([]model.NodeReward, DefaultItemsPerPage)

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	for _, reward := range *rewards {
		if queryString == "" {
			result = *rewards
			break
		}

		if reward.Matches(queryString) {
			result = append(result, reward)
		}
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(result), itemsPerPage, pageIndex)
	result = (result)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(&result)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}


// swagger:route GET /nodes/actions/history Nodes getNodeActionHistory
//
// Returns gravity node mockup actions history
//
// This will show all gravity node mockup actions history list
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//       api_key:
//
//     Responses:
//       200: []NodeHistoryRecord
func (rc *ResponseController) GetNodeActionsHistory(w http.ResponseWriter, req *http.Request) {
	rewards := utils.GetNodeActionsHistoryMockup()
	result := make([]model.NodeHistoryRecord, DefaultItemsPerPage)

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	for _, reward := range *rewards {
		if queryString == "" {
			result = *rewards
			break
		}

		if reward.Matches(queryString) {
			result = append(result, reward)
		}
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(result), itemsPerPage, pageIndex)
	result = (result)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(&result)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}


// swagger:route GET /nodes/exact Nodes getExactNode
//
// Returns exact node by name
//
// This will show all gravity node mockup actions history list
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Security:
//       api_key:
//
//     Responses:
//       200: Node
//		 404: null
func (rc *ResponseController) GetExactNode (w http.ResponseWriter, req *http.Request) {

}