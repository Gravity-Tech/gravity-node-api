package utils

import (
	"../model"
	"time"
)

func AddNodes (nebula model.Nebula, nodes ...model.Node) model.Nebula {
	nebula.NodesUsing = append(nebula.NodesUsing, nodes...)

	return nebula
}

func AddNebulas (node model.Node, nebulas ...model.Nebula) model.Node {
	node.NebulasUsing = append(node.NebulasUsing, nebulas...)

	return node
}


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
		Status:          model.PendingStatus,
		Description:     "",
		Score:           50,
		SubscriptionFee: 10 * WAVES_DECIMAL,
		TargetChain:     model.WAVES_TARGET_CHAIN,
		Extractor:       nil,
		NodesUsing:      nil,
	}
	var binanceNebula = model.Nebula{
		Name:            "Binance Nebula",
		Status:          model.ActiveStatus,
		Description:     "",
		Score:           100,
		TargetChain:     model.ETH_TARGET_CHAIN,
		SubscriptionFee: 10 * ETH_DECIMAL,
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