
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
//     Host: node.gravityhub.org
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
	"flag"
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/controller"
	"github.com/Gravity-Tech/gravity-node-api/controller/response"
	"github.com/Gravity-Tech/gravity-node-api/updater"
	"net/http"
)

var port string
var isDebug, shouldFill bool

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func init() {
	flag.StringVar(&port, "port", "8090", "Path to config.toml")
	flag.BoolVar(&shouldFill, "fill", false, "Should fill the db")
	flag.BoolVar(&isDebug, "debug", false, "is debug mode enabled")

	flag.Parse()
}

func main () {
	db := controller.NewDBController()

	ch := make(chan int, 1)
	if shouldFill {
		ch <- 0
		go db.PersistMockup()
		<- ch
	}

	ch <- 0
	go db.RefreshNebulasAndNodesMaterializedView()
	<- ch

	http.HandleFunc("/hello", headers)

	responseController := &response.ResponseController{}
	responseController.DBControllerDelegate = db
	responseController.Handle()

	cacheUpdater := updater.NewNodesCacheUpdater()
	go cacheUpdater.Start()

	fmt.Printf("Listening on port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}