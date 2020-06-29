package model

type RawData = byte

type IExtractor interface {
	extractData() []RawData
	mapData(data []RawData) interface{}
}


// swagger: model
type Extractor struct {
	// Data feed tag for distinct usage
	//
	// required: true
	DataFeedTag string `json:"data_feed_tag"`

	// Common extractor description
	//
	// required: true
	Description string `json:"description"`
}

func (e *Extractor) extractData() []RawData {
	return []RawData{
		byte(0),
	}
}

func (e *Extractor) mapData() interface{} {
	return 0
}