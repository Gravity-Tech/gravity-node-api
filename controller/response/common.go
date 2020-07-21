package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	stats := rc.DBControllerDelegate.CommonStats()

	bytes, _ := json.Marshal(stats)

	addBaseHeaders(w.Header())

	_, _ = fmt.Fprint(w, string(bytes))
}