package model

import (
	"time"
)

// swagger:model
type Transaction struct {
	TxId     uint64    `json:"tx_id"`
	TxHash   string    `json:"tx_hash"`
	FuncType string    `json:"func_type"`
	Height   int64     `json:"height"`
	Time     time.Time `json:"time"`
}

func (tx *Transaction) Matches(str string) bool {
	fieldValues := []string{tx.FuncType}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type Swap struct {
	TxId         uint64    `json:"tx_id"`
	TxHash       string    `json:"tx_hash"`
	FuncType     string    `json:"func_type"`
	Height       int64     `json:"height"`
	Time         time.Time `json:"time"`
	DataId       uint64    `json:"data_id"`
	Commit       string    `json:"commit"`
	NebulaId     string    `json:"nebula_id"`
	PulseId      uint64    `json:"pulse_id"`
	TcHeight     uint64    `json:"tc_height"`
	OraclePubKey string    `json:"oracle_pubkey"`
	Rq           string    `json:"rq"`
	Amount       string    `json:"amount"`
	Receiver     string    `json:"receiver"`
}

func (swap *Swap) Matches(str string) bool {
	fieldValues := []string{swap.NebulaId, swap.Receiver}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type CommitTransaction struct {
	DataId       uint64 `json:"data_id"`
	TxId         uint64 `json:"tx_id"`
	NebulaId     string `json:"nebula_id"`
	PulseId      uint64 `json:"pulse_id"`
	TcHeight     uint64 `json:"tc_height"`
	Commit       string `json:"commit"`
	OraclePubKey string `json:"oracle_pubkey"`
}

func (tx *CommitTransaction) Matches(str string) bool {
	fieldValues := []string{tx.NebulaId, tx.OraclePubKey}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type RevealTransaction struct {
	DataId       uint64 `json:"data_id"`
	TxId         uint64 `json:"tx_id"`
	Commit       string `json:"commit"`
	NebulaId     string `json:"nebula_id"`
	PulseId      uint64 `json:"pulse_id"`
	TcHeight     uint64 `json:"tc_height"`
	OraclePubKey string `json:"oracle_pubkey"`
	Rq           string `json:"rq"`
	Amount       string `json:"amount"`
	Receiver     string `json:"receiver"`
}

func (tx *RevealTransaction) Matches(str string) bool {
	fieldValues := []string{tx.NebulaId, tx.OraclePubKey, tx.Receiver, tx.Rq}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type AddOracleTransaction struct {
	DataId       uint64 `json:"data_id"`
	TxId         uint64 `json:"tx_id"`
	ChainType    string `json:"chain_type"`
	OraclePubKey string `json:"oracle_pubkey"`
}

func (tx *AddOracleTransaction) Matches(str string) bool {
	fieldValues := []string{tx.ChainType, tx.OraclePubKey}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type AddOracleInNebulaTransaction struct {
	DataId       uint64 `json:"data_id"`
	TxId         uint64 `json:"tx_id"`
	NebulaId     string `json:"nebula_id"`
	OraclePubKey string `json:"oracle_pubkey"`
}

func (tx *AddOracleInNebulaTransaction) Matches(str string) bool {
	fieldValues := []string{tx.NebulaId, tx.OraclePubKey}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type ResultTransaction struct {
	DataId       uint64 `json:"data_id"`
	TxId         uint64 `json:"tx_id"`
	NebulaId     string `json:"nebula_id"`
	PulseId      uint64 `json:"pulse_id"`
	Sign         string `json:"sign"`
	ChainType    string `json:"chain_type"`
	OraclePubKey string `json:"oracle_pubkey"`
}

func (tx *ResultTransaction) Matches(str string) bool {
	fieldValues := []string{tx.NebulaId, tx.ChainType, tx.OraclePubKey}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type NewRoundTransaction struct {
	DataId uint64 `json:"data_id"`
	TxId   uint64 `json:"tx_id"`
}

func (tx *NewRoundTransaction) Matches(str string) bool {
	fieldValues := []string{}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type VoteTransaction struct {
	DataId uint64 `json:"data_id"`
	TxId   uint64 `json:"tx_id"`
}

func (tx *VoteTransaction) Matches(str string) bool {
	fieldValues := []string{}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type AddNebulaTransaction struct {
	DataId   uint64 `json:"data_id"`
	TxId     uint64 `json:"tx_id"`
	NebulaId string `json:"nebula_id"`
	B        string `json:"b"`
}

func (tx *AddNebulaTransaction) Matches(str string) bool {
	fieldValues := []string{tx.NebulaId}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type DropNebulaTransaction struct {
	DataId   uint64 `json:"data_id"`
	TxId     uint64 `json:"tx_id"`
	NebulaId string `json:"nebula_id"`
}

func (tx *DropNebulaTransaction) Matches(str string) bool {
	fieldValues := []string{tx.NebulaId}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type SignNewConsulsTransaction struct {
	DataId    uint64 `json:"data_id"`
	TxId      uint64 `json:"tx_id"`
	ChainType string `json:"chain_type"`
	RoundId   uint64 `json:"round_id"`
	Sign      string `json:"sign"`
}

func (tx *SignNewConsulsTransaction) Matches(str string) bool {
	fieldValues := []string{tx.ChainType}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type SignNewOraclesTransaction struct {
	DataId   uint64 `json:"data_id"`
	TxId     uint64 `json:"tx_id"`
	RoundId  uint64 `json:"round_id"`
	Sign     string `json:"sign"`
	NebulaId string `json:"nebula_id"`
}

func (tx *SignNewOraclesTransaction) Matches(str string) bool {
	fieldValues := []string{tx.NebulaId}

	return MatchStrList(fieldValues, str)
}

// swagger:model
type ApproveLastRoundTransaction struct {
	DataId uint64 `json:"data_id"`
	TxId   uint64 `json:"tx_id"`
}

func (tx *ApproveLastRoundTransaction) Matches(str string) bool {
	fieldValues := []string{}

	return MatchStrList(fieldValues, str)
}
