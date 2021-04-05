package response

import (
	"encoding/json"
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/model"
	"github.com/Gravity-Tech/gravity-node-api/utils"
	"net/http"
)


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

	var matchesList []*model.Node
	for _, node := range *nodeList {
		if queryString == "" { break }

		if node.Matches(queryString) {
			matchesList = append(matchesList, node)
		}
	}
	if queryString != "" {
		nodeList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*nodeList), itemsPerPage, pageIndex)
	*nodeList = (*nodeList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(nodeList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}


// swagger:route GET /nodes/exact Nodes getExactNode
//
// Returns exact node by name
//
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
	address := req.URL.Query().Get("q")
	publicKey := req.URL.Query().Get("pubKey")

	var exactNode *model.Node

	if address != ""  {
		exactNode = rc.DBControllerDelegate.ExactNode(address)
	} else if publicKey != "" {
		exactNode = rc.DBControllerDelegate.ExactNodeByPubKey(publicKey)
	}

	bytes, _ := json.Marshal(exactNode)

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

func (rc *ResponseController) GetAllTransactions(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllTransactionsList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.Transaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactTransaction(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.Transaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactTransaction(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllSwaps(w http.ResponseWriter, req *http.Request) {
	swapList := rc.DBControllerDelegate.AllSwapsList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.Swap
	for _, swap := range *swapList {
		if queryString == "" {
			break
		}

		if swap.Matches(queryString) {
			matchesList = append(matchesList, swap)
		}
	}
	if queryString != "" {
		swapList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*swapList), itemsPerPage, pageIndex)
	*swapList = (*swapList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(swapList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllCommit(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllCommitList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.CommitTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactCommit(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.CommitTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactCommit(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllReveal(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllRevealList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.RevealTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactReveal(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.RevealTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactReveal(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllAddOracle(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllAddOracleList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.AddOracleTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactAddOracle(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.AddOracleTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactAddOracle(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllAddOracleInNebula(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllAddOracleInNebulaList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.AddOracleInNebulaTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactAddOracleInNebula(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.AddOracleInNebulaTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactAddOracleInNebula(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllResult(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllResultList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.ResultTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactResult(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.ResultTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactResult(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllNewRound(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllNewRoundList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.NewRoundTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactNewRound(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.NewRoundTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactNewRound(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllVote(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllVoteList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.VoteTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactVote(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.VoteTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactVote(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllAddNebula(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllAddNebulaList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.AddNebulaTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactAddNebula(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.AddNebulaTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactAddNebula(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllDropNebula(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllDropNebulaList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.DropNebulaTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactDropNebula(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.DropNebulaTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactDropNebula(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllSignNewConsuls(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllSignNewConsulsList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.SignNewConsulsTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactSignNewConsuls(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.SignNewConsulsTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactSignNewConsuls(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllSignNewOracles(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllSignNewOraclesList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.SignNewOraclesTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactSignNewOracles(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.SignNewOraclesTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactSignNewOracles(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetAllApproveLastRound(w http.ResponseWriter, req *http.Request) {
	txList := rc.DBControllerDelegate.AllApproveLastRoundList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.ApproveLastRoundTransaction
	for _, transaction := range *txList {
		if queryString == "" {
			break
		}

		if transaction.Matches(queryString) {
			matchesList = append(matchesList, transaction)
		}
	}
	if queryString != "" {
		txList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1
	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*txList), itemsPerPage, pageIndex)
	*txList = (*txList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(txList)
	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
func (rc *ResponseController) GetExactApproveLastRound(w http.ResponseWriter, req *http.Request) {
	txId := req.URL.Query().Get("q")

	var exactTx *model.ApproveLastRoundTransaction

	if txId != "" {
		exactTx = rc.DBControllerDelegate.ExactApproveLastRound(txId)
	}

	bytes, _ := json.Marshal(exactTx)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}
