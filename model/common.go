package model

type CommonStatus = uint

const (
	CommonPendingStatus CommonStatus = iota
	CommonDoneStatus
)

// swagger:model
type CommonStats struct {
	NodesCount uint `json:"nodes_count"`
	Pulses uint `json:"pulses"`
	DataFeeds uint `json:"data_feeds"`
}