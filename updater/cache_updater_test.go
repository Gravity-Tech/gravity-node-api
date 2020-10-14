package updater

import (
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/client"
	"github.com/Gravity-Tech/gravity-node-api/controller"
	"github.com/Gravity-Tech/gravity-node-api/model"
	"github.com/go-pg/pg/v10"
	"testing"
)

var ledgerNodeEndpoint = "ledger.gravityhub.org"
var nodeIPMapRecord *model.NodeIPMapRecord

var dbController *controller.DBController
var oldNodeVersion, newNodeVersion *model.Node
var oldNodeFieldValue, newNodeFieldValue string

func InsertDefinedNodeIPRecord() {
	dbController = controller.NewDBController()
	dbController.DB = pg.Connect(&pg.Options{
		Addr: ":5432",
		User:     "postgres",
		Password: "123123123",
		Database: "gravity-api",
	})

	path := fmt.Sprintf("http://%v", ledgerNodeEndpoint)
	newClient := client.NewLedgerClient(path)

	validatorStatus, err := newClient.FetchValidatorStatus()

	handleInitError(err)
	if err != nil { return }

	fmt.Printf("Validator status: %+v \n", validatorStatus.ValidatorInfo.PubKey)

	nodeIPMapRecord = &model.NodeIPMapRecord{
		IPAddress: ledgerNodeEndpoint,
		PublicKey: validatorStatus.ValidatorInfo.PubKey.Value,
	}

	_, _ = dbController.DB.Model(nodeIPMapRecord).Insert()

	handleInitError(err)
}

func CleanupDefinedNodeIPRecord() {
	db := controller.NewDBController()

	_, err := db.DB.Model(&model.NodeIPMapRecord{
		IPAddress: nodeIPMapRecord.IPAddress,
	}).WherePK().Delete()

	handleInitError(err)
}

func TestUpdateNodeDetails(t *testing.T) {
	path := fmt.Sprintf("http://%v", nodeIPMapRecord.IPAddress)
	t.Logf("Ledger client: %v \n", path)
	ledgerClient := client.NewLedgerClient(path)

	validatorStatus, err := ledgerClient.FetchValidatorStatus()
	validatorDetails, err := ledgerClient.FetchValidatorDetails()
	handleTestError(err, t)

	// DROP existing node
	nodeToDrop := &model.Node{ PublicKey: nodeIPMapRecord.PublicKey }
	_, _ = dbController.DB.Model(nodeToDrop).WherePK().Delete()
	handleTestError(err, t)

	_ = dbController.UpdateNodeDetails(nodeIPMapRecord.PublicKey, validatorDetails, validatorStatus)
	handleTestError(err, t)

	t.Cleanup(CleanupDefinedNodeIPRecord)
}

func TestMain(m *testing.M) {
	InsertDefinedNodeIPRecord()
	m.Run()
}

func handleInitError(err error) {
	if err != nil {
		fmt.Printf("Error occured on init: %v \n", err)
	}
}
func handleTestError(err error, t *testing.T) {
	if err != nil {
		t.Log("Test failed: ", err)
		t.Fail()
	}
}