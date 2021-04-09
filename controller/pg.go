package controller

import (
	"fmt"
	"github.com/Gravity-Tech/gravity-core/config"
	"github.com/Gravity-Tech/gravity-node-api/client"
	"github.com/Gravity-Tech/gravity-node-api/migrations/common"
	"github.com/Gravity-Tech/gravity-node-api/model"
	"github.com/Gravity-Tech/gravity-node-api/utils"
	"github.com/go-pg/pg/v10"
)

type DBController struct {
	DB *pg.DB
}

const (
	materializedViewPostfix = "_materialized_view"
)

func NewDBController() *DBController {
	return &DBController{ DB: utils.ConnectToPG() }
}

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
		_, err := dbc.DB.Model(&nebula).Insert()

		dbc.errorHandle("Nebula", err)
	}
}

func (dbc *DBController) persistNodes(nodes *[]model.Node) {

	for _, node := range *nodes {
		_, err := dbc.DB.Model(&node).Insert()

		dbc.errorHandle("Node", err)
	}
}

func (dbc *DBController) persistDatafeedsList(datafeedsList *[]*model.Extractor) {
	for _, datafeed := range *datafeedsList {
		_, err := dbc.DB.Model(datafeed).Insert()

		dbc.errorHandle("Datafeed", err)
	}
}

func (dbc *DBController) UpdateNodeDetails(publicKey string, details *config.ValidatorDetails, status *client.ValidatorStatus) error {
	db := dbc.DB

	node := &model.Node{ PublicKey: publicKey }
	node.UpdateByValidatorDetails(details, status)

	var selectTo model.Node
	err := db.Model(&selectTo).Where("public_key = ?", publicKey).Select()

	doesExist := err == nil

	if doesExist {
		// Update existing
		_, err := db.Model(node).WherePK().Update()
		dbc.errorHandle("UpdateNodeDetails - update existing", err)
		if err != nil {
			return err
		}
	} else {
		// Create
		_, err := db.Model(node).Insert()

		dbc.errorHandle("UpdateNodeDetails - create new", err)
		if err != nil {
			return err
		}
	}

	dbc.errorHandle("UpdateNodeDetails", err)
	return err
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

func (dbc *DBController) CommonStats () *model.CommonStats {
	var stats model.CommonStats

	_, err := dbc.DB.Query(&stats, fmt.Sprintf("SELECT * FROM %v LIMIT 1;", model.DefaultExtendedDBTableNames.CommonStats))
	dbc.errorHandle("CommonStats", err)

	return &stats

}

func (dbc *DBController) AllNodeIPsRecords () *[]*model.NodeIPMapRecord {
	var list []*model.NodeIPMapRecord

	_, err := dbc.DB.Query(&list, fmt.Sprintf("SELECT * FROM %v;", model.DefaultExtendedDBTableNames.NodeIPsMap))
	dbc.errorHandle("AllNodeIPsRecords", err)

	return &list
}

func (dbc *DBController) AllNodesList () *[]*model.Node {
	var list []*model.Node

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *, %[1]v.ip_address
		FROM %[1]v
		INNER JOIN  %[2]v
		ON  %[1]v.public_key =  %[2]v.public_key;
	`, model.DefaultExtendedDBTableNames.NodeIPsMap, model.DefaultExtendedDBTableNames.Nodes))
	dbc.errorHandle("AllNodesList", err)

	return &list
}

func (dbc *DBController) mapTableToMaterializedView(tableName string) string {
	return tableName + materializedViewPostfix
}

func (dbc *DBController) ExactNodeByPubKey (address string) *model.Node {
	var node model.Node

	destination := model.DefaultExtendedDBTableNames.Nodes

	result, err := dbc.DB.Query(&node, fmt.Sprintf("SELECT * FROM %v WHERE public_key = '%v';", destination, address))
	dbc.errorHandle("ExactNode", err)

	if result.RowsReturned() == 0 { return nil }

	return &node
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
func (dbc *DBController) AllTransactionsList() *[]*model.Transaction {
	var list []*model.Transaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v
 ORDER BY height DESC;
	`, model.DefaultExtendedDBTableNames.Transactions))
	dbc.errorHandle("AllTransactionsList", err)

	return &list
}

func (dbc *DBController) ExactTransaction(txId string) *model.Transaction {
	var tx model.Transaction

	destination := model.DefaultExtendedDBTableNames.Transactions

	result, err := dbc.DB.Query(&tx, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactTransaction", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &tx
}

// TODO should be a factory

func (dbc *DBController) AllSwapsList() *[]*model.Swap {
	var list []*model.Swap

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
        SELECT *
        FROM
          (SELECT DISTINCT ON (commit) *
           FROM %[1]v
           NATURAL JOIN %[2]v
           ORDER BY commit, tx_id DESC) as alias
        ORDER BY tx_id DESC;
	`, model.DefaultExtendedDBTableNames.RevealTable, model.DefaultExtendedDBTableNames.Transactions))
	dbc.errorHandle("AllSwapsList", err)

	return &list
}
func (dbc *DBController) AllCommitList() *[]*model.CommitTransaction {
	var list []*model.CommitTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.CommitTable))
	dbc.errorHandle("AllCommitList", err)

	return &list
}
func (dbc *DBController) ExactCommit(txId string) *model.CommitTransaction {
	var data model.CommitTransaction

	destination := model.DefaultExtendedDBTableNames.CommitTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactCommit", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllRevealList() *[]*model.RevealTransaction {
	var list []*model.RevealTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.RevealTable))
	dbc.errorHandle("AllRevealList", err)

	return &list
}
func (dbc *DBController) ExactReveal(txId string) *model.RevealTransaction {
	var data model.RevealTransaction

	destination := model.DefaultExtendedDBTableNames.RevealTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactReveal", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllAddOracleList() *[]*model.AddOracleTransaction {
	var list []*model.AddOracleTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.AddOracleTable))
	dbc.errorHandle("AddOracleList", err)

	return &list
}
func (dbc *DBController) ExactAddOracle(txId string) *model.AddOracleTransaction {
	var data model.AddOracleTransaction

	destination := model.DefaultExtendedDBTableNames.AddOracleTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactAddOracle", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllAddOracleInNebulaList() *[]*model.AddOracleInNebulaTransaction {
	var list []*model.AddOracleInNebulaTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.AddOracleInNebulaTable))
	dbc.errorHandle("AllAddOracleInNebulaList", err)

	return &list
}
func (dbc *DBController) ExactAddOracleInNebula(txId string) *model.AddOracleInNebulaTransaction {
	var data model.AddOracleInNebulaTransaction

	destination := model.DefaultExtendedDBTableNames.AddOracleInNebulaTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactAddOracleInNebula", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllResultList() *[]*model.ResultTransaction {
	var list []*model.ResultTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.ResultTable))
	dbc.errorHandle("AllResultList", err)

	return &list
}
func (dbc *DBController) ExactResult(txId string) *model.ResultTransaction {
	var data model.ResultTransaction

	destination := model.DefaultExtendedDBTableNames.ResultTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactResult", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllNewRoundList() *[]*model.NewRoundTransaction {
	var list []*model.NewRoundTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.NewRoundTable))
	dbc.errorHandle("AllNewRoundList", err)

	return &list
}
func (dbc *DBController) ExactNewRound(txId string) *model.NewRoundTransaction {
	var data model.NewRoundTransaction

	destination := model.DefaultExtendedDBTableNames.NewRoundTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactNewRound", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllVoteList() *[]*model.VoteTransaction {
	var list []*model.VoteTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.VoteTable))
	dbc.errorHandle("AllVoteList", err)

	return &list
}
func (dbc *DBController) ExactVote(txId string) *model.VoteTransaction {
	var data model.VoteTransaction

	destination := model.DefaultExtendedDBTableNames.VoteTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactVote", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllAddNebulaList() *[]*model.AddNebulaTransaction {
	var list []*model.AddNebulaTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.AddNebulaTable))
	dbc.errorHandle("AllAddNebulaList", err)

	return &list
}
func (dbc *DBController) ExactAddNebula(txId string) *model.AddNebulaTransaction {
	var data model.AddNebulaTransaction

	destination := model.DefaultExtendedDBTableNames.AddNebulaTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactAddNebula", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllDropNebulaList() *[]*model.DropNebulaTransaction {
	var list []*model.DropNebulaTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.DropNebulaTable))
	dbc.errorHandle("AllDropNebulaList", err)

	return &list
}
func (dbc *DBController) ExactDropNebula(txId string) *model.DropNebulaTransaction {
	var data model.DropNebulaTransaction

	destination := model.DefaultExtendedDBTableNames.DropNebulaTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactDropNebula", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllSignNewConsulsList() *[]*model.SignNewConsulsTransaction {
	var list []*model.SignNewConsulsTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.SignNewConsulsTable))
	dbc.errorHandle("AllSignNewConsulsList", err)

	return &list
}
func (dbc *DBController) ExactSignNewConsuls(txId string) *model.SignNewConsulsTransaction {
	var data model.SignNewConsulsTransaction

	destination := model.DefaultExtendedDBTableNames.SignNewConsulsTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactSignNewConsuls", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllSignNewOraclesList() *[]*model.SignNewOraclesTransaction {
	var list []*model.SignNewOraclesTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.SignNewOraclesTable))
	dbc.errorHandle("AllSignNewOraclesList", err)

	return &list
}
func (dbc *DBController) ExactSignNewOracles(txId string) *model.SignNewOraclesTransaction {
	var data model.SignNewOraclesTransaction

	destination := model.DefaultExtendedDBTableNames.SignNewOraclesTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactSignNewOracles", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
func (dbc *DBController) AllApproveLastRoundList() *[]*model.ApproveLastRoundTransaction {
	var list []*model.ApproveLastRoundTransaction

	_, err := dbc.DB.Query(&list, fmt.Sprintf(`
		SELECT *
		FROM %[1]v;
	`, model.DefaultExtendedDBTableNames.ApproveLastRoundTable))
	dbc.errorHandle("AllApproveLastRoundList", err)

	return &list
}
func (dbc *DBController) ExactApproveLastRound(txId string) *model.ApproveLastRoundTransaction {
	var data model.ApproveLastRoundTransaction

	destination := model.DefaultExtendedDBTableNames.ApproveLastRoundTable

	result, err := dbc.DB.Query(&data, fmt.Sprintf("SELECT * FROM %v WHERE tx_id = '%v';", destination, txId))
	dbc.errorHandle("ExactApproveLastRound", err)

	if result.RowsReturned() == 0 {
		return nil
	}

	return &data
}
