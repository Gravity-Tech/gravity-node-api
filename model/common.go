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

// A NotFoundError is the error message that is generated when server could not find what was requested.
//
// swagger:response notFoundError
type NotFoundError struct {
	// in: body
	Body struct {
		// HTTP status code
		// example: 404
		// default: 404
		Code    int32 `json:"code"`
		Message error `json:"message"`
	} `json:"body"`
}
