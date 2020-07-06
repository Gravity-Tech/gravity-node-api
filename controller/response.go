package controller

import (
	"../model"
	"../utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller interface {}

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
func GetAllNebulas(w http.ResponseWriter, req *http.Request) {
	nebulasList, _ := utils.GetMockup()

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
func GetAllNodes(w http.ResponseWriter, req *http.Request) {
	_, nodeList := utils.GetMockup()

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
func GetCommonStats(w http.ResponseWriter, req *http.Request) {
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
func GetNodeRewardsList(w http.ResponseWriter, req *http.Request) {
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