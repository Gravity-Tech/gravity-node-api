package decoder

import (
	"strings"
)

type ID [32]byte

type Type string

const (
	String Type = "string"
	Int    Type = "int"
	Bytes  Type = "bytes"
)

type Arg struct {
	Type  Type
	Value []byte
}

type TxFunc string

const (
	Commit            TxFunc = "commit"
	Reveal            TxFunc = "reveal"
	AddOracle         TxFunc = "addOracle"
	AddOracleInNebula TxFunc = "addOracleInNebula"
	Result            TxFunc = "result"
	NewRound          TxFunc = "newRound"
	Vote              TxFunc = "vote"
	AddNebula         TxFunc = "setNebula"
	DropNebula        TxFunc = "dropNebula"
	SignNewConsuls    TxFunc = "signNewConsuls"
	SignNewOracles    TxFunc = "signNewOracles"
	ApproveLastRound  TxFunc = "approveLastRound"
)

type Transaction struct {
	Id           ID
	SenderPubKey [32]byte
	Signature    [72]byte
	Func         TxFunc
	Timestamp    uint64
	Args         []Arg
}

type ChainType byte

func (ch ChainType) String() string {
	switch ch {
	case Ethereum:
		return "ethereum"
	case Waves:
		return "waves"
	case Binance:
		return "bsc"
	default:
		return "ethereum"
	}
}

const (
	Ethereum ChainType = iota
	Waves
	Binance
)

type OraclesPubKey [33]byte
type NebulaId [32]byte

const (
	NebulaIdLength        = 32
	EthereumAddressLength = 20
	BSCAddressLength      = 20
	WavesAddressLength    = 26
)

type DataType string
type ExtractorData struct {
	Type  DataType
	Value string
}

const (
	StringDataType DataType = "string"
	Int64DataType  DataType = "int64"
	Base64DataType DataType = "base64"
)

type ExtractorType uint8

const (
	Int64ExtractorType ExtractorType = iota
	StringExtractorType
	BytesExtractorType
)

func ParseExtractorType(extractorType string) (ExtractorType, error) {
	switch strings.ToLower(extractorType) {
	case "int64":
		return Int64ExtractorType, nil
	case "string":
		return StringExtractorType, nil
	case "bytes":
		return BytesExtractorType, nil
	default:
		return 0, nil
	}
}
