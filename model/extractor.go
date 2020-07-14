package model

type RawData = byte

// swagger: model
type Extractor struct {
	tableName struct{} `sql:"data_feeds"`

	// Data feed tag for distinct usage
	//
	// required: true
	DataFeedTag string `json:"datafeed_tag" pg:"data_feed_tag"`

	// Common extractor description
	//
	// required: true
	Description string `json:"description"`
}


func (ext *Extractor) Matches (str string) bool {
	fieldValues := []string { ext.DataFeedTag, ext.Description }

	return MatchStrList(fieldValues, str)
}