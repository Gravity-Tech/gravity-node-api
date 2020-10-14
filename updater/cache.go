package updater

import (
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/controller"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/client"
	core_config "github.com/Gravity-Tech/gravity-core/config"
	"time"
)

type entityUpdater interface {
	UpdateEntity()
}

type NodesCacheUpdater struct {}

func NewNodesCacheUpdater () *NodesCacheUpdater {
	return &NodesCacheUpdater{}
}

func (updater *NodesCacheUpdater) Start() {
	ledgerClient := client.NewLedgerClient("http://ledger.gravityhub.org/abci_query?path=\"validatorDetails\"")

	details, err := ledgerClient.FetchValidatorDetails()

	if err != nil {
		fmt.Printf("Error: %v \n", err)
		return
	}

	fmt.Printf("Details: %v \n", details)
}

/**
	Takes node_ips_map as base for updating
 */
func (updater *NodesCacheUpdater) UpdateEntity() {
	db := controller.NewDBController()
	nodeIPsRecords := db.AllNodeIPsRecords()

	if len(*nodeIPsRecords) == 0 {
		updater.log("Node IPs map is empty. Nothing to update")
		return
	}
	//
	//for _, nodeIPRecord := range *nodeIPsRecords {
	//	nodeIPRecord.IPAddress
	//}

}

//func (updater *NodesCacheUpdater) Fetch

func (updater *NodesCacheUpdater) log(message interface{}) {
	fmt.Printf("%v - NodesCacheUpdater: %v \n", time.Now().Format(time.RFC3339), message)
}