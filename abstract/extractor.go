package abstract

import (
	"encoding/json"
	"time"
)

type RawData = byte

type IExtractor interface {
	DataFeedTag() string
	Description() string
	extractData() []RawData
	mapData(data []RawData) interface{}
}

/**
	- Concrete extractor implementations are down below.
	- However, interface declaration right at the top
	- is constraints abstraction that should every
	- gravity node extractor conform to.

	- For extractor description template consider
	- using the pattern "This extractor <does_this>. <additional_info>"

	- Parameters based on specific data are considered to be
	- resolved/parsed by certain extractor author

	- That code is not considered as *CORE* in Gravity node.
	- It's just an example for extractor as external service.
*/

type binancePriceExtractor struct {}
type binancePrice = int

func (e *binancePriceExtractor) DataFeedTag() string {
	return "binance-price_WAVES_BTC"
}

func (e *binancePriceExtractor) Description() string {
	return "This extractor represents WAVES_BTC binance price"
}

func (e *binancePriceExtractor) extractData() []RawData {
	return []RawData {0,1}
}

// Return value is decimal
func (e *binancePriceExtractor) mapData(data []RawData) binancePrice {
	currentPrice := binancePrice(data[0])
	return currentPrice
}

type commonWeatherDataExtractor struct {}
type commonWeatherData struct {
	WeekDay time.Weekday
	Temperature int
}

func (e *commonWeatherDataExtractor) DataFeedTag() string {
	return "global-weather-forecast"
}

func (e *commonWeatherDataExtractor) Description() string {
	return "This extractor resolves approximate weather forecast on certain date/city"
}

func (e *commonWeatherDataExtractor) extractData() []RawData {
	return []RawData {0,1}
}

func (e *commonWeatherDataExtractor) mapData(data []RawData) *commonWeatherData {
	var weatherData commonWeatherData
	_ = json.Unmarshal(data, &weatherData)
	return &weatherData
}