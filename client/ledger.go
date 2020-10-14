package client

import (
	"encoding/json"
	"fmt"
	coreconfig "github.com/Gravity-Tech/gravity-core/config"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	//ctypes "github.com/tendermint/tendermint/rpc/core/types"
	rpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	//rpcclient "github.com/tendermint/tendermint/rpc/client"

	"io/ioutil"
	"net/http"
)

type LedgerClient struct {
	EndpointURL string
}

func NewLedgerClient(endpoint string) *LedgerClient {
	return &LedgerClient{ EndpointURL:endpoint }
}

type responseResult struct {
	Response abcitypes.ResponseQuery `json:"response"`
}
type fetchResponse struct {
	Result *responseResult `json:"result"`
}

func (ledger *LedgerClient) extractData(response fetchResponse) string {
	return string(response.Result.Response.Value)
}

//func (ledger *LedgerClient) extractValidatorInfo(response *ctypes.ResultStatus) *ctypes.ValidatorInfo {
//	return &response.ValidatorInfo
//}

type ValidatorStatus struct {
	ValidatorInfo *ValidatorInfo `json:"validator_info"`
}
type ValidatorInfo struct {
	Address string
	PubKey struct {
		Type string `json:"type"`
		Value string `json:"value"`
	} `json:"pub_key"`
}

//{
//	"validator_info": {
//	"address": "C52EB90B1DF941E17CA50F73248A225C95FAAF2D",
//	"pub_key": {
//	"type": "tendermint/PubKeyEd25519",
//	"value": "iGUd/zzoOUHU9vg7wRTYQG9i6KwIUpA/Ke9aH+KZiVE="
//	},
//	"voting_power": "0"
//	}
//}

func (ledger *LedgerClient) FetchValidatorStatus() (*ValidatorStatus, error) {
	path := fmt.Sprintf("%v/status", ledger.EndpointURL)

	response, err := http.Get(path)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	byteValue, err := ioutil.ReadAll(response.Body)

	var statusResponse rpctypes.RPCResponse

	err = json.Unmarshal(byteValue, &statusResponse)

	if err != nil {
		return nil, err
	}

	var resultResponse ValidatorStatus
	err = json.Unmarshal(statusResponse.Result, &resultResponse)

	if err != nil {
		return nil, err
	}

	return &resultResponse, nil
}

func (ledger *LedgerClient) FetchValidatorDetails() (*coreconfig.ValidatorDetails, error) {
	path := fmt.Sprintf("%v/abci_query?path=\"validatorDetails\"", ledger.EndpointURL)

	response, err := http.Get(path)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	byteValue, err := ioutil.ReadAll(response.Body)

	var parsedResponse *fetchResponse
	err = json.Unmarshal(byteValue, &parsedResponse)

	if err != nil {
		return nil, err
	}

	ledgerDataExtracted := ledger.extractData(*parsedResponse)

	var validatorDetails coreconfig.ValidatorDetails
	err = json.Unmarshal([]byte(ledgerDataExtracted), &validatorDetails)

	if err != nil {
		return nil, err
	}

	return &validatorDetails, nil
}