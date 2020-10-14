package updater

import (
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/client"
	"github.com/Gravity-Tech/gravity-node-api/controller"
	"github.com/Gravity-Tech/gravity-node-api/model"
	"testing"
)

var ledgerNodeEndpoint = "ledger.gravityhub.org"
var nodeIPMapRecord *model.NodeIPMapRecord

var oldNodeVersion, newNodeVersion *model.Node
var oldNodeFieldValue, newNodeFieldValue string

func InsertDefinedNodeIPRecord() {
	path := fmt.Sprintf("http://%v", ledgerNodeEndpoint)
	newClient := client.NewLedgerClient(path)

	validatorStatus, err := newClient.FetchValidatorStatus()

	handleInitError(err)
	if err != nil { return }

	fmt.Printf("Valiator status: %+v \n", validatorStatus.ValidatorInfo.PubKey)

	nodeIPMapRecord = &model.NodeIPMapRecord{
		IPAddress: ledgerNodeEndpoint,
		PublicKey: string(validatorStatus.ValidatorInfo.PubKey.Bytes()),
	}

	db := controller.NewDBController()

	err = db.DB.Insert(nodeIPMapRecord)

	handleInitError(err)
}

func CleanupDefinedNodeIPRecord() {
	db := controller.NewDBController()

	err := db.DB.Delete(nodeIPMapRecord)

	handleInitError(err)
}

func TestUpdateNodeDetailsOnNew(t *testing.T) {
	path := fmt.Sprintf("http://%v", nodeIPMapRecord.IPAddress)
	t.Logf("Ledger client: %v \n", path)
	ledgerClient := client.NewLedgerClient(path)

	validatorDetails, err := ledgerClient.FetchValidatorDetails()
	handleTestError(err, t)

	db := controller.NewDBController()

	// DROP existing node
	err = db.DB.Delete(&model.Node{ PublicKey: nodeIPMapRecord.PublicKey })
	handleTestError(err, t)

	err = db.UpdateNodeDetails(nodeIPMapRecord.PublicKey, validatorDetails)
	handleTestError(err, t)
}

func TestUpdateNodeDetailsExisting(t *testing.T) {

}

func TestMain(m *testing.M) {
	InsertDefinedNodeIPRecord()
	m.Run()
	CleanupDefinedNodeIPRecord()
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