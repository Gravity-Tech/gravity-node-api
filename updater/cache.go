package updater

import (
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/client"
	"github.com/Gravity-Tech/gravity-node-api/controller"
	"github.com/Gravity-Tech/gravity-node-api/decoder"
	"github.com/Gravity-Tech/gravity-node-api/model"
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

	latestBlock, err := updater.LatestBlock()
	if err != nil {
		fmt.Printf("Error on fetching latestBlock: %v\n", err)
		return
	}
	fmt.Printf("latestBlock: %v\n", *latestBlock)

	lastCachedBlock, err := updater.LastCachedBlock()
	if err != nil {
		fmt.Printf("Error on fetching lastCachedBlock: %v\n", err)
		return
	}
	fmt.Printf("lastCachedBlock: %v\n", *lastCachedBlock)

	for i := *lastCachedBlock; i < *latestBlock; i++ {
		fmt.Printf("Caching block %v\n", i)
		err = updater.CacheTransactions(i)
		if err != nil {
			break
		}
	}

	time.Sleep(time.Minute * 10)
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

	err = db.UpdateNodeDetails(pubKey, details, validatorStatus)
	if err != nil {
		updater.log(fmt.Sprintf("Error occured on node details update: %+v; \n", err))
	} else {
		updater.log(fmt.Sprintf("Updated Node details: %+v; Status: %v; \n", details, validatorStatus))
	}

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
	}
	wg.Wait()
}

//func (updater *NodesCacheUpdater) Fetch

func (updater *NodesCacheUpdater) log(message interface{}) {
	fmt.Printf("%v - NodesCacheUpdater: %v \n", time.Now().Format(time.RFC3339), message)
}

func (updater *NodesCacheUpdater) CacheTransactions(height int64) error {
	var err error

	ledgerClient := client.NewLedgerClient("http://134.122.37.128:26657")

	blockInfo, err := ledgerClient.FetchBlock(height)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}

	db := updater.DB.DB
	txs := blockInfo.Block.Data.Txs
	blockTime := blockInfo.Block.Header.Time

	var tx model.Transaction

	for i := 0; i < len(txs); i++ {

		txHash := txs[i]
		txP := decoder.ParseTx(txHash)
		funcType := string(txP.Func)
		tx = model.Transaction{
			TxId:     0,
			TxHash:   txHash,
			FuncType: funcType,
			Height:   height,
			Time:     blockTime}

		_, err = db.Model(&tx).
			OnConflict("DO NOTHING").
			Insert()
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		err = updater.CacheArgs(txP.Func, txP.Args, tx.TxId)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	time.Sleep(time.Second * 1)
	return nil
}

func (updater *NodesCacheUpdater) LatestBlock() (*int64, error) {
	var err error

	ledgerClient := client.NewLedgerClient("http://134.122.37.128:26657")

	latestBlock, err := ledgerClient.LatestBlock()
	if err != nil {
		fmt.Printf("Error latestBlock %s\n", err)
		return nil, err
	}

	return latestBlock, nil
}

func (updater *NodesCacheUpdater) LastCachedBlock() (*int64, error) {

	var err error
	db := updater.DB.DB
	var tx model.Transaction

	err = db.Model(&tx).Order("height DESC").Limit(1).Select()
	if err != nil && err.Error() == "pg: no rows in result set" {
		fmt.Printf("Init cache %s\n")
		n := int64(1)
		return &n, nil
	} else if err != nil {
		fmt.Printf("Error lastCachedBlock %s\n", err)
		return nil, err
	}

	return &tx.Height, nil
}
