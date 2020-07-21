package model

import (
	"regexp"
	"strings"
)

type CommonStatus = uint

const (
	CommonPendingStatus CommonStatus = iota
	CommonFailedStatus
	CommonDoneStatus
)

// swagger:model
type CommonStats struct {
	NodesCount uint `json:"nodes_count"`
	Pulses uint `json:"pulses"`
	DataFeeds uint `json:"data_feeds"`
}

type ISearchable interface {
	MatchStrList([]string, string) bool
}

func MatchStrList (fieldValues []string, str string) bool {
	str = strings.ToLower(str)
	regex, err := regexp.Compile(str)

	if err != nil { return false }

	for _, value := range fieldValues {
		value = strings.ToLower(value)

		matched := regex.Match([]byte(value))

		if matched { return true }
	}

	return false
}
