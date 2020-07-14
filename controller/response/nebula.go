package response

import (
	"encoding/json"
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/model"
	"net/http"
)


func addBaseHeaders (headers http.Header) {
	headers.Add("Content-Type", "application/json")
	headers.Set("Access-Control-Allow-Origin", "*")
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

	var matchesList []*model.Nebula
	for _, nebula := range *nebulasList {
		if queryString == "" { break }

		//if !nebula.Matches(queryString) {
		//	*nebulasList = append((*nebulasList)[:i], (*nebulasList)[i+1:]...)
		//}

		if nebula.Matches(queryString) {
			matchesList = append(matchesList, nebula)
		}
	}
	if queryString != "" {
		nebulasList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1

	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*nebulasList), itemsPerPage, pageIndex)
	*nebulasList = (*nebulasList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(nebulasList)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}


// swagger:route GET /nebulas/exact Nebulas getExactNebula
//
// Returns exact Nebula by address
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
//       200: Nebula
//		 404: null
func (rc *ResponseController) GetExactNebula (w http.ResponseWriter, req *http.Request) {
	address := req.URL.Query().Get("q")
	exactNebula := rc.DBControllerDelegate.ExactNebula(address)

	bytes, _ := json.Marshal(exactNebula)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}