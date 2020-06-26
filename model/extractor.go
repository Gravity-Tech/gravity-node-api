package model

type RawData = byte

type IExtractor interface {
	extractData() []RawData
	mapData(data []RawData) interface{}
}

type Extractor struct {}