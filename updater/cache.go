package updater

import (
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/controller"
	"github.com/Gravity-Tech/gravity-node-api/client"
	"sync"

	//core_config "github.com/Gravity-Tech/gravity-core/config"
	"time"
)

type entityUpdater interface {
	UpdateEntity()
}

type NodesCacheUpdater struct {
	DB *controller.DBController
}

func NewNodesCacheUpdater () *NodesCacheUpdater {
	return &NodesCacheUpdater{
		DB: controller.NewDBController(),
	}
}

func (updater *NodesCacheUpdater) Start() {
	updater.UpdateEntity()

	time.Sleep(time.Second * 10)
	updater.Start()
}


func (updater *NodesCacheUpdater) updateNode(endpoint string, wg *sync.WaitGroup, db *controller.DBController, pubKey string) error {
	defer wg.Done()

	ledgerClient := client.NewLedgerClient(endpoint)

	validatorStatus, err := ledgerClient.FetchValidatorStatus()
	details, err := ledgerClient.FetchValidatorDetails()

	if err != nil {
		fmt.Printf("Error on fetch: %v \n", err)
		return err
	}


	updater.log(fmt.Sprintf("Updated Node details: %+v; Status: %v; \n", details, validatorStatus))

	err = db.UpdateNodeDetails(pubKey, details, validatorStatus)

	return nil
}

/**
	Takes node_ips_map as base for updating
 */
func (updater *NodesCacheUpdater) UpdateEntity() {
	db := updater.DB
	nodeIPsRecords := db.AllNodeIPsRecords()

	if len(*nodeIPsRecords) == 0 {
		updater.log("Node IPs map is empty. Nothing to update")
		return
	}

	var wg sync.WaitGroup
	for _, nodeRecord := range *nodeIPsRecords {
		wg.Add(1)

		endpoint := fmt.Sprintf("http://%v", nodeRecord.IPAddress)
		go updater.updateNode(endpoint, &wg, db, nodeRecord.PublicKey)
		//go updater.updateNode(nodeRecord.IPAddress, &wg)
	}
	wg.Wait()
}

//func (updater *NodesCacheUpdater) Fetch

func (updater *NodesCacheUpdater) log(message interface{}) {
	fmt.Printf("%v - NodesCacheUpdater: %v \n", time.Now().Format(time.RFC3339), message)
}