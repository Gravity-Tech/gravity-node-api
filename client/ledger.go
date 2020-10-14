package client

import (
	"encoding/json"
	"fmt"
	"github.com/Gravity-Tech/gravity-core/config"
	"github.com/ethereum/go-ethereum/common/hexutil"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"io/ioutil"
	"net/http"
)

type LedgerClient struct {
	EndpointURL string
}

func NewLedgerClient(endpoint string) *LedgerClient {
	return &LedgerClient{ EndpointURL:endpoint }
}

type fetchResponse struct {
	Result abcitypes.ResponseQuery `json:"result"`
}

func (ledger *LedgerClient) extractData(response fetchResponse) string {
	return string(response.Result.Value)
}

func (ledger *LedgerClient) FetchValidatorDetails() (*core_config.ValidatorDetails, error) {

	// http://ledger.gravityhub.org/abci_query?path="validatorDetails"
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

	validatorDetailsBytes, err := hexutil.Decode(ledger.extractData(*parsedResponse))

	if err != nil {
		return nil, err
	}

	var validatorDetails core_config.ValidatorDetails
	err = json.Unmarshal(validatorDetailsBytes, validatorDetails)

	if err != nil {
		return nil, err
	}

	return &validatorDetails, nil
}