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
	newClient := client.NewLedgerClient(ledgerNodeEndpoint)

	validatorStatus, err := newClient.FetchValidatorStatus()

	handleInitError(err)
	if err != nil { return }

	nodeIPMapRecord := &model.NodeIPMapRecord{
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
	client := client.NewLedgerClient(nodeIPMapRecord.IPAddress)

	validatorDetails, err := client.FetchValidatorDetails()
	handleTestError(err, t)

	db := controller.NewDBController()

	db.UpdateNodeDetails(nodeIPMapRecord.PublicKey, validatorDetails)


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