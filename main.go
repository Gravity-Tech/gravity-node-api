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