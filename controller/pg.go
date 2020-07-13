package controller

import (
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/model"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/utils"
	"github.com/go-pg/pg"
)

type DBController struct {
	DB *pg.DB
}

func (dbc *DBController) PersistMockup () {
	nebulas, nodes := utils.GetMockup()

	dbc.persistNebulas(nebulas)
	dbc.persistNodes(nodes)
}

func (dbc *DBController) errorHandle (prefix string, err error) {
	if err != nil {
		fmt.Printf("%v; Error occured: %v\n", prefix, err)
	}
}

func (dbc *DBController) persistNebulas(nebulas *[]model.Nebula) {

	for _, nebula := range *nebulas {
		err := dbc.DB.Insert(&nebula)

		dbc.errorHandle("Nebula", err)
	}
}

func (dbc *DBController) persistNodes(nodes *[]model.Node) {

	for _, node := range *nodes {
		err := dbc.DB.Insert(&node)

		dbc.errorHandle("Node", err)
	}
}

