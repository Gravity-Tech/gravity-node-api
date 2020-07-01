package model

type RawData = byte

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