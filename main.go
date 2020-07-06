
// Package classification Gravity Node API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: node.gravityhub.org:8090
//     BasePath: /
//     Version: 1.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: venlab.dev <shamil@venlab.dev> https://venlab.dev
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package main

import (
	"fmt"
	"flag"
	"net/http"
	"./router"
	"./controller"
)

var port string

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func init() {
	flag.StringVar(&port, "port", "8090", "Path to config.toml")
	flag.Parse()
}

func main () {
	http.HandleFunc("/hello", headers)
	http.HandleFunc(router.GetAllNebulas, controller.GetAllNebulas)
	http.HandleFunc(router.GetAllNodes, controller.GetAllNodes)
	http.HandleFunc(router.GetCommonStats, controller.GetCommonStats)
	http.HandleFunc(router.GetNodeRewards, controller.GetNodeRewardsList)
	http.HandleFunc(router.GetNodeActionsHistory, controller.GetAllNodes)

	fmt.Printf("Listening on port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}