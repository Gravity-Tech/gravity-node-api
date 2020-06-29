package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"../utils"
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

	bytes, _ := json.Marshal(nebulasList)

	addBaseHeaders(w.Header())

	fmt.Fprint(w, string(bytes))
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

	bytes, _ := json.Marshal(nodeList)
	addBaseHeaders(w.Header())

	fmt.Fprint(w, string(bytes))
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

	fmt.Fprint(w, string(bytes))
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

	bytes, _ := json.Marshal(rewards)

	addBaseHeaders(w.Header())

	fmt.Fprint(w, string(bytes))
}