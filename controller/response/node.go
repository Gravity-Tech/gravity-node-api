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
//		 404: notFoundError
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

// swagger:route GET /transactions/all Transactions getAllTransactions
//
// Returns all available gravity transactions
//
// This will show all available gravity transactions
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
//       200: []Transaction
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
// swagger:route GET /transactions/exact Transactions getExactTransaction
//
// Returns exact transaction by txId
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
//       200: Transaction
//		 404: notFoundError
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

// swagger:route GET /transactions/swap/all Swaps getAllSwaps
//
// Returns all available gravity swaps
//
// This will show all available gravity swaps
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
//       200: []Swap
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

// swagger:route GET /transactions/commit/all Commit getAllCommit
//
// Returns all available gravity commit transactions
//
// This will show all available gravity commit transactions
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
//       200: []CommitTransaction
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
// swagger:route GET /transactions/commit/exact Commit getExactCommit
//
// Returns exact commit by txId
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
//       200: CommitTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/reveal/all Reveal getAllReveal
//
// Returns all available gravity reveal transactions
//
// This will show all available gravity reveal transactions
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
//       200: []RevealTransaction
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
// swagger:route GET /transactions/reveal/exact Reveal getExactReveal
//
// Returns exact reveal by txId
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
//       200: RevealTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/add-oracle/all AddOracle getAllAddOracle
//
// Returns all available gravity add-oracle transactions
//
// This will show all available gravity add-oracle transactions
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
//       200: []AddOracleTransaction
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
// swagger:route GET /transactions/add-oracle/exact AddOracle getExactAddOracle
//
// Returns exact add-oracle by txId
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
//       200: AddOracleTransaction
//		 404: notFoundError
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
// swagger:route GET /transactions/add-oracle-in-nebula/all AddOracleInNebula getAllAddOracleInNebula
//
// Returns all available gravity add-oracle-in-nebula transactions
//
// This will show all available gravity add-oracle-in-nebula transactions
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
//       200: []AddOracleInNebulaTransaction
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
// swagger:route GET /transactions/add-oracle-in-nebula/exact AddOracleInNebula getExactAddOracleInNebula
//
// Returns exact add-oracle-in-nebula by txId
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
//       200: AddOracleInNebulaTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/result/all Result getAllResult
//
// Returns all available gravity result transactions
//
// This will show all available gravity result transactions
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
//       200: []ResultTransaction
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
// swagger:route GET /transactions/result/exact Result getExactResult
//
// Returns exact result by txId
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
//       200: ResultTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/new-round/all NewRound getAllNewRound
//
// Returns all available gravity new-round transactions
//
// This will show all available gravity new-round transactions
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
//       200: []NewRoundTransaction
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
// swagger:route GET /transactions/new-round/exact NewRound getExactNewRound
//
// Returns exact new-round by txId
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
//       200: NewRoundTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/vote/all Vote getAllVote
//
// Returns all available gravity vote transactions
//
// This will show all available gravity vote transactions
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
//       200: []VoteTransaction
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
// swagger:route GET /transactions/vote/exact Vote getExactVote
//
// Returns exact vote by txId
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
//       200: VoteTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/add-nebula/all AddNebula getAllAddNebula
//
// Returns all available gravity add-nebula transactions
//
// This will show all available gravity add-nebula transactions
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
//       200: []AddNebulaTransaction
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
// swagger:route GET /transactions/add-nebula/exact AddNebula getExactAddNebula
//
// Returns exact add-nebula by txId
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
//       200: AddNebulaTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/drop-nebula/all DropNebula getAllDropNebula
//
// Returns all available gravity drop-nebula transactions
//
// This will show all available gravity drop-nebula transactions
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
//       200: []DropNebulaTransaction
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
// swagger:route GET /transactions/drop-nebula/exact DropNebula getExactDropNebula
//
// Returns exact drop-nebula by txId
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
//       200: DropNebulaTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/sign-new-consuls/all SignNewConsuls getAllSignNewConsuls
//
// Returns all available gravity sign-new-consuls transactions
//
// This will show all available gravity sign-new-consuls transactions
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
//       200: []SignNewConsulsTransaction
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
// swagger:route GET /transactions/sign-new-consuls/exact SignNewConsuls getExactSignNewConsuls
//
// Returns exact sign-new-consuls by txId
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
//       200: SignNewConsulsTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/sign-new-oracles/all SignNewOracles getAllSignNewOracles
//
// Returns all available gravity sign-new-oracles transactions
//
// This will show all available gravity sign-new-oracles transactions
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
//       200: []SignNewOraclesTransaction
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
// swagger:route GET /transactions/sign-new-oracles/exact SignNewOracles getExactSignNewOracles
//
// Returns exact sign-new-oracles by txId
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
//       200: SignNewOraclesTransaction
//		 404: notFoundError
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

// swagger:route GET /transactions/approve-last-round/all ApproveLastRound getAllApproveLastRound
//
// Returns all available gravity approve-last-round transactions
//
// This will show all available gravity approve-last-round transactions
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
//       200: []ApproveLastRoundTransaction
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
// swagger:route GET /transactions/approve-last-round/exact ApproveLastRound getExactApproveLastRound
//
// Returns exact approve-last-round by txId
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
//       200: ApproveLastRoundTransaction
//		 404: notFoundError
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
