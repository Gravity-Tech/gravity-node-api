package controller

import (
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/migrations/common"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/model"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/utils"
	"github.com/go-pg/pg"
)

type DBController struct {
	DB *pg.DB
}

const (
	materializedViewPostfix = "_materialized_view"
)

func (dbc *DBController) PersistMockup () {
	nebulas, nodes := utils.GetMockup()
	datafeeds := utils.GetDatafeedsMockup(250)

	dbc.persistNebulas(nebulas)
	dbc.persistNodes(nodes)

	dbc.persistDatafeedsList(datafeeds)
}

func (dbc *DBController) RefreshNebulasAndNodesMaterializedView () {
	queries := []string {
		common.UpdateMaterializedViewQuery(model.DefaultExtendedDBTableNames.Nebulas),
		common.UpdateMaterializedViewQuery(model.DefaultExtendedDBTableNames.Nodes),
	}

	for _, query := range queries {
		_, err := dbc.DB.Query(nil, query)
		if err != nil {
			fmt.Printf("Error on refresh: %v;\n", err)
		}
	}
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

func (dbc *DBController) persistDatafeedsList(datafeedsList *[]*model.Extractor) {
	for _, datafeed := range *datafeedsList {
		err := dbc.DB.Insert(datafeed)

		dbc.errorHandle("Datafeed", err)
	}
}

func (dbc *DBController) AllDatafeedsList() *[]*model.Extractor {
	var list []*model.Extractor

	_, err := dbc.DB.Query(&list, fmt.Sprintf("SELECT * FROM %v;", model.DefaultExtendedDBTableNames.Datafeeds))
	dbc.errorHandle("AllNebulasList", err)

	return &list
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

	destination := model.DefaultExtendedDBTableNames.Nodes

	result, err := dbc.DB.Query(&node, fmt.Sprintf("SELECT * FROM %v WHERE address = '%v';", destination, address))
	dbc.errorHandle("ExactNode", err)

	if result.RowsReturned() == 0 { return nil }

	return &node
}

func (dbc *DBController) ExactNebula (address string) *model.Nebula {
	var nebula model.Nebula

	destination := model.DefaultExtendedDBTableNames.Nebulas

	result, err := dbc.DB.Query(&nebula, fmt.Sprintf("SELECT * FROM %v WHERE address = '%v';", destination, address))
	dbc.errorHandle("ExactNebula", err)

	if result.RowsReturned() == 0 { return nil }

	return &nebula
}