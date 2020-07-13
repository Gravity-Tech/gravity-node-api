package utils

import (
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/model"
	petname "github.com/dustinkirkland/golang-petname"
	"time"
)


// Pure functions
func AddNodes (nebula model.Nebula, nodes ...model.Node) model.Nebula {
	nebula.AddNodes(nodes...)

	return nebula
}
func AddNebulas (node model.Node, nebulas ...model.Nebula) model.Node {
	//node.NebulasUsing = append(node.NebulasUsing, nebulas...)
	node.AddNebulas(nebulas...)

	return node
}

func stringifyBigInt(any interface{}) string {
	return fmt.Sprintf("%d", any)
}

func stringify(any interface{}) string {
	return fmt.Sprintf("%v", any)
}

// Mockup functions
func GetMockup () (*[]model.Nebula, *[]model.Node)  {
	var demoNode = model.Node{
		Name:          "Demo Node #1",
		Description:   "Demo Desc",
		Score:         5,
		DepositChain:  model.WAVES_TARGET_CHAIN,
		DepositAmount: 25,
		JoinedAt:      time.Time{}.Unix(),
		//NebulasUsing:  nil,
	}
	var binanceNode = model.Node{
		Name:          "Binance Node #1",
		Description:   "Binance Desc",
		Score:         3,
		DepositChain:  model.ETH_TARGET_CHAIN,
		DepositAmount: 25,
		JoinedAt:      time.Time{}.Unix(),
		//NebulasUsing:  nil,
	}
	var huobiNode = model.Node{
		Name:          "LinkPool Node",
		Description:   `
			LinkPool is a leading Chainlink node service provider with the goal 
			of providing tools and services that benefit the Chainlink ecosystem. 
			Our aims include lowering the barrier to entry to staking on Chainlink nodes,
			easing the amount of technical experience required to run a Chainlink node and
			providing smart contract creators with the tools to easily search and identify Chainlink
			nodes that can suit their data requirements.
		`,
		Score:         7,
		DepositChain:  model.ETH_TARGET_CHAIN,
		DepositAmount: 25,
		JoinedAt:      time.Time{}.Unix(),
		//NebulasUsing:  nil,
	}

	var demoNebula = model.Nebula{
		Name:            "Demo Nebula",
		Status:          model.NebulaPendingStatus,
		Description:     "",
		Score:           50,
		SubscriptionFee: stringifyBigInt(10 * WavesDecimal),
		TargetChain:     model.WAVES_TARGET_CHAIN,
		Regularity:		 1440,
		//NodesUsing:      nil,
	}
	var binanceNebula = model.Nebula{
		Name:            "Binance Nebula",
		Status:          model.NebulaActiveStatus,
		Description:     "",
		Score:           100,
		TargetChain:     model.ETH_TARGET_CHAIN,
		Regularity:		 1440,
		SubscriptionFee: stringifyBigInt(10 * EthDecimal),
		//NodesUsing:      nil,
	}
	var coinbaseNebula = model.Nebula{
		Name:            "Coinbase Nebula",
		Status:          model.NebulaActiveStatus,
		Description:     "Coinbase",
		Score:           100,
		TargetChain:     model.ETH_TARGET_CHAIN,
		Regularity:		 1440,
		SubscriptionFee: stringifyBigInt(10 * EthDecimal),
		//NodesUsing:      nil,
	}

	nebulaList := []model.Nebula {
		AddNodes(demoNebula, demoNode),
		AddNodes(binanceNebula, binanceNode),
		coinbaseNebula,
	}
	nebulaList = append(nebulaList, *DuplicateToSimilarNebulas(&coinbaseNebula, 100)...)

	nodeList := []model.Node {
		AddNebulas(demoNode, demoNebula),
		AddNebulas(binanceNode, binanceNebula),
		huobiNode,
	}
	nodeList = append(nodeList, *DuplicateToSimilarNodes(&huobiNode, 100)...)

	return &nebulaList, &nodeList
}

func DuplicateToSimilarNebulas(nebula *model.Nebula, amount int) *[]model.Nebula {
	var arr []model.Nebula

	index := 0

	for {
		if index >= amount { break }

		newNebula := *nebula
		newNebula.Name += " " + petname.Name() + " " + petname.Adjective()
		newNebula.Score += model.Score(5 * index)
		newNebula.Description = petname.Generate(2, " ")

		arr = append(arr, newNebula)

		index++
	}

	return &arr
}


func DuplicateToSimilarNodes(node *model.Node, amount int) *[]model.Node {
	var arr []model.Node

	index := 0

	for {
		if index >= amount { break }

		newNode := *node
		newNode.Name += " " + petname.Name() + " " + petname.Adjective()
		newNode.Score += 5
		newNode.Description = petname.Generate(2, " ")

		arr = append(arr, newNode)

		index++
	}

	return &arr
}


func GetCommonStatsMockup() *model.CommonStats {
	stats := model.CommonStats{
		NodesCount: 25,
		Pulses:     20 * 1000,
		DataFeeds:  125,
	}
	return &stats
}

func GetNodeRewardsListMockup() *[]model.NodeReward {
	rewards := []model.NodeReward {
		model.NodeReward{
			Amount:    10 * WavesDecimal,
			Decimals:   8,
			Timestamp: time.Time{}.Unix(),
			Currency:  "WAVES",
		},
		model.NodeReward{
			Amount:    11 * WavesDecimal,
			Timestamp: time.Time{}.Unix(),
			Currency:  "WAVES",
		},
		model.NodeReward{
			Amount:    12 * WavesDecimal,
			Decimals:   8,
			Timestamp: time.Time{}.Unix(),
			Currency:  "WAVES",
		},
		model.NodeReward{
			Amount:    15 * EthDecimal,
			Decimals:   18,
			Timestamp: time.Time{}.Unix(),
			Currency:  "ETH",
		},
		model.NodeReward{
			Amount:    3456 * WavesDecimal,
			Decimals:   8,
			Timestamp: time.Time{}.Unix(),
			Currency:  "WAVES",
		},
		model.NodeReward{
			Amount:    1 * EthDecimal,
			Decimals:   18,
			Timestamp: time.Time{}.Unix(),
			Currency:  "ETH",
		},
	}

	return &rewards
}

//
//NodeDataSent NodeActionType = iota
//NodeConsensusReached
//NodeAssetsReceived
//NodeVoteReceived
func GetNodeActionsHistoryMockup() *[]model.NodeHistoryRecord {
	list := []model.NodeHistoryRecord {
		model.NodeHistoryRecord{
			Name:      "A",
			Type:      model.NodeDataSent,
			Asset:     model.Asset{
				Name:     "WAVES",
				Decimals: WavesDecimal,
				Amount:   "534",
			},
			Status:    model.CommonFailedStatus,
			Timestamp: time.Time{}.Unix(),
		},
		model.NodeHistoryRecord{
			Name:      "B",
			Type:      model.NodeConsensusReached,
			Asset:     model.Asset{
				Name:     "ETH",
				Decimals: EthDecimal,
				Amount:   "123",
			},
			Status:    model.CommonDoneStatus,
			Timestamp: time.Time{}.Unix(),
		},
		model.NodeHistoryRecord{
			Name:      "C",
			Type:      model.NodeAssetsReceived,
			Asset:     model.Asset{
				Name:     "WAVES",
				Decimals: WavesDecimal,
				Amount:   "130259",
			},
			Status:    model.CommonPendingStatus,
			Timestamp: time.Time{}.Unix(),
		},
		model.NodeHistoryRecord{
			Name:      "D",
			Type:      model.NodeVoteReceived,
			Asset:     model.Asset{
				Name:     "ETH",
				Decimals: EthDecimal,
				Amount:   "5",
			},
			Status:    model.CommonFailedStatus,
			Timestamp: time.Time{}.Unix(),
		},
	}

	return &list
}