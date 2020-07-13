package controller

import (
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/model"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/utils"
	"github.com/go-pg/pg"
	migrations "github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/migrations"
)

type DBController struct {
	DB *pg.DB
}

const (
	materializedViewPostfix = "_materialized_view"
)

func (dbc *DBController) PersistMockup () {
	nebulas, nodes := utils.GetMockup()

	dbc.persistNebulas(nebulas)
	dbc.persistNodes(nodes)

	migrations.UpdateMaterializedViewQuery(model.DefaultExtendedDBTableNames.Nebulas)
	migrations.UpdateMaterializedViewQuery(model.DefaultExtendedDBTableNames.Nodes)
}

func (dbc *DBController) errorHandle (prefix string, err error) {
	if err != nil {
		fmt.Printf("Method: %v; Error occured: %v\n", prefix, err)
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

func (dbc *DBController) AllNebulasList () *[]*model.Nebula {
	var list []*model.Nebula

	_, err := dbc.DB.Query(&list, fmt.Sprintf("SELECT * FROM %v;", model.DefaultExtendedDBTableNames.Nebulas))
	dbc.errorHandle("AllNebulasList", err)

	return &list
}

func (dbc *DBController) AllNodesList () *[]*model.Node {
	var list []*model.Node

	_, err := dbc.DB.Query(&list, fmt.Sprintf("SELECT * FROM %v;", model.DefaultExtendedDBTableNames.Nodes))
	dbc.errorHandle("AllNodesList", err)

	return &list
}

func (dbc *DBController) mapTableToMaterializedView(tableName string) string {
	return tableName + materializedViewPostfix
}

func (dbc *DBController) ExactNode (address string) *model.Node {
	var node model.Node

	destination := dbc.mapTableToMaterializedView(model.DefaultExtendedDBTableNames.Nodes)

	_, err := dbc.DB.Query(&node, fmt.Sprintf("SELECT * FROM %v WHERE address = '%v';", destination, address))
	dbc.errorHandle("ExactNode", err)

	return &node
}

func (dbc *DBController) ExactNebula (address string) *model.Nebula {
	var nebula model.Nebula

	destination := dbc.mapTableToMaterializedView(model.DefaultExtendedDBTableNames.Nebulas)

	_, err := dbc.DB.Query(&nebula, fmt.Sprintf("SELECT * FROM %v WHERE address = '%v';", destination, address))
	dbc.errorHandle("ExactNebula", err)

	return &nebula
}