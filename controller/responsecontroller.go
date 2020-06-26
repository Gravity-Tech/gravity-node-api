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

func GetAllNebulas(w http.ResponseWriter, req *http.Request) {
	nebulasList, _ := utils.GetMockup()

	bytes, _ := json.Marshal(nebulasList)

	addBaseHeaders(w.Header())

	fmt.Fprint(w, string(bytes))
}

func GetAllNodes(w http.ResponseWriter, req *http.Request) {
	_, nodeList := utils.GetMockup()

	bytes, _ := json.Marshal(nodeList)

	addBaseHeaders(w.Header())

	fmt.Fprint(w, string(bytes))
}