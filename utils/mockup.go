package utils

import (
	"../model"
	"time"
)


// Pure functions
func AddNodes (nebula model.Nebula, nodes ...model.Node) model.Nebula {
	nebula.NodesUsing = append(nebula.NodesUsing, nodes...)

	return nebula
}
func AddNebulas (node model.Node, nebulas ...model.Nebula) model.Node {
	node.NebulasUsing = append(node.NebulasUsing, nebulas...)

	return node
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
		NebulasUsing:  nil,
	}
	var binanceNode = model.Node{
		Name:          "Binance Node #1",
		Description:   "Binance Desc",
		Score:         3,
		DepositChain:  model.ETH_TARGET_CHAIN,
		DepositAmount: 25,
		JoinedAt:      time.Time{}.Unix(),
		NebulasUsing:  nil,
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
		NebulasUsing:  nil,
	}


	var demoNebula = model.Nebula{
		Name:            "Demo Nebula",
		Status:          model.NebulaPendingStatus,
		Description:     "",
		Score:           50,
		SubscriptionFee: 10 * WavesDecimal,
		TargetChain:     model.WAVES_TARGET_CHAIN,
		Extractor:       nil,
		NodesUsing:      nil,
	}
	var binanceNebula = model.Nebula{
		Name:            "Binance Nebula",
		Status:          model.NebulaActiveStatus,
		Description:     "",
		Score:           100,
		TargetChain:     model.ETH_TARGET_CHAIN,
		SubscriptionFee: 10 * EthDecimal,
		Extractor:       nil,
		NodesUsing:      nil,
	}

	nebulaList := []model.Nebula {
		AddNodes(demoNebula, demoNode),
		AddNodes(binanceNebula, binanceNode),
	}

	nodeList := []model.Node {
		AddNebulas(demoNode, demoNebula),
		AddNebulas(binanceNode, binanceNebula),
		huobiNode,
	}

	return &nebulaList, &nodeList
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
			Amount:    9523 * EthDecimal,
			Decimals:   18,
			Timestamp: time.Time{}.Unix(),
			Currency:  "ETH",
		},
	}

	return &rewards
}