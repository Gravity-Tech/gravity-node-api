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
	http.HandleFunc(router.GET_ALL_NEBULAS, controller.GetAllNebulas)
	http.HandleFunc(router.GET_ALL_NODES, controller.GetAllNodes)

	fmt.Printf("Listening on port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}