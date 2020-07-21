package response

import (
	"encoding/json"
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/model"
	"net/http"
)

// swagger:route GET /data/datafeeds/all Datafeeds getAvailableDataFeeds
//
// Returns available data feeds
//
// Data feeds
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
//       200: []Extractor
func (rc *ResponseController) GetAvailableDatafeedsList(w http.ResponseWriter, req *http.Request) {
	datafeedList := rc.DBControllerDelegate.AllDatafeedsList()

	queryString, queryPage, queryItemsPerPage := HandleParams(req)

	var matchesList []*model.Extractor
	for _, df := range *datafeedList {
		if queryString == "" { break }

		//if !df.Matches(queryString) {
		//	*datafeedList = append((*datafeedList)[:i], (*datafeedList)[i+1:]...)
		//}

		if df.Matches(queryString) {
			matchesList = append(matchesList, df)
		}
	}
	if queryString != "" {
		datafeedList = &matchesList
	}

	currentPage, itemsPerPage := RevealParams(queryPage, queryItemsPerPage)
	pageIndex := currentPage - 1

	pageIndexStart, pageIndexEnd := ComputeSliceRange(len(*datafeedList), itemsPerPage, pageIndex)
	*datafeedList = (*datafeedList)[pageIndexStart:pageIndexEnd]

	bytes, _ := json.Marshal(datafeedList)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}