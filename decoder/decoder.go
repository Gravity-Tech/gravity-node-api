package decoder

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func ParseTx(s string) Transaction {

	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("error:", err)
	}

	var tx Transaction
	err = json.Unmarshal(b, &tx)
	if err != nil {
		fmt.Println("error:", err)
	}

	return tx
}

func ParseStringValue(value []byte) string {
	return string(value[:])
}
func ParseIntValue(value []byte) uint64 {
	return binary.BigEndian.Uint64(value)
}
func ParseBytesValue(value []byte) []byte {
	return value
}
func ParseChainType(value []byte) ChainType {
	return ChainType(value[0])
}
func ParseNebulaId(value []byte) string {

	trim := bytes.TrimLeft(value, "\x00")
	l := len(trim)
	if l == 20 {
		return hexutil.Encode(trim[:])
	} else if l == 26 {
		return base58.Encode(trim[:])
	}

	return ""
}
func NebulaToBytes(value []byte, chainType ChainType) []byte {
	switch chainType {
	case Binance:
		return value[NebulaIdLength-BSCAddressLength:]
	case Ethereum:
		return value[NebulaIdLength-EthereumAddressLength:]
	case Waves:
		return value[NebulaIdLength-WavesAddressLength:]
	}

	return nil
}
func ParseOraclePubKey(value []byte) string {

	trim := bytes.TrimLeft(value, "\x00")
	l := len(trim)
	if l == 33 {
		return hexutil.Encode(trim)
	} else if l == 32 {
		return base58.Encode(trim)
	}

	return ""
}
func ParseDataBytes(value []byte) (
	rq string,
	amount string,
	receiver string) {

	if value[0] == byte('m') {
		return ParseDataBytesEth(value)
	} else {
		return ParseDataBytesWaves(value)
	}
}

func ParseDataBytesEth(value []byte) (
	rq string,
	amount string,
	receiver string) {

	rest := value
	rest = value[1:]
	rqIntBytes := rest[:32]
	rest = rest[32:]
	amountBytes := rest[:32]
	rest = rest[32:]
	receiverBytes := rest[:20]
	rest = rest[20:]

	var rqInt big.Int
	rqInt.SetBytes(rqIntBytes)
	rqIntStr := base58.Encode(rqIntBytes)
	var amountBig big.Int
	amountBig.SetBytes(amountBytes)
	receiver = hexutil.Encode(receiverBytes)

	return rqIntStr, amountBig.String(), receiver
}
func ParseDataBytesWaves(value []byte) (
	rq string,
	amount string,
	receiver string) {

	rest := value
	rest = rest[8:]
	rqIntBytes := rest[:32]
	rest = rest[32:]
	amountBytes := rest[:8]
	rest = rest[8:]
	receiverBytes := rest[:26]
	rest = rest[26:]

	rqIntStr := base58.Encode(rqIntBytes)

	var amountBig big.Int
	amountBig.SetBytes(amountBytes)

	receiver = base58.Encode(receiverBytes)

	return rqIntStr, amountBig.String(), receiver
}

func OraclePubKeyToBytes(value []byte, chainType ChainType) []byte {
	var v []byte
	switch chainType {
	case Ethereum, Binance:
		v = value[:33]
	case Waves:
		v = value[1:33]
	}
	return v
}
func ParseCommit(value []byte) string {
	return hexutil.Encode(value)
}
func ParseSign(value []byte) string {
	return base58.Encode(value)
}

func ParseArgsCommit(args []Arg) (
	nebulaId string,
	pulseId uint64,
	tcHeight uint64,
	commit string,
	oraclePubKey string) {
	nebulaId = ParseNebulaId(args[0].Value)
	pulseId = ParseIntValue(args[1].Value)
	tcHeight = ParseIntValue(args[2].Value)
	commit = ParseCommit(args[3].Value)
	oraclePubKey = ParseOraclePubKey(args[4].Value)
	return nebulaId,
		pulseId,
		tcHeight,
		commit,
		oraclePubKey
}
func ParseArgsReveal(args []Arg) (
	commit string,
	nebulaId string,
	pulseId uint64,
	tcHeight uint64,
	oraclePubKey string,
	rq string,
	amount string,
	receiver string) {
	commit = ParseCommit(args[0].Value)
	nebulaId = ParseNebulaId(args[1].Value)
	pulseId = ParseIntValue(args[2].Value)
	tcHeight = ParseIntValue(args[3].Value)
	oraclePubKey = ParseOraclePubKey(args[5].Value)
	rq, amount, receiver = ParseDataBytes(args[4].Value)
	return commit,
		nebulaId,
		pulseId,
		tcHeight,
		oraclePubKey,
		rq,
		amount,
		receiver
}
func ParseArgsAddOracle(args []Arg) (
	chainType ChainType,
	oraclePubKey string) {
	chainType = ParseChainType(args[0].Value)
	oraclePubKey = ParseOraclePubKey(args[1].Value)
	return chainType, oraclePubKey
}
func ParseArgsAddOracleInNebula(args []Arg) (
	nebulaId string,
	oraclePubKey string) {
	nebulaId = ParseNebulaId(args[0].Value)
	oraclePubKey = ParseOraclePubKey(args[1].Value)
	return nebulaId, oraclePubKey
}
func ParseArgsResult(args []Arg) (
	nebulaId string,
	pulseId uint64,
	sign string,
	chainType ChainType,
	oraclePubKey string) {
	nebulaId = ParseNebulaId(args[0].Value)
	pulseId = ParseIntValue(args[1].Value)
	sign = ParseSign(args[2].Value)
	chainType = ParseChainType(args[3].Value)
	oraclePubKey = ParseOraclePubKey(args[4].Value)
	return nebulaId,
		pulseId,
		sign,
		chainType,
		oraclePubKey
}
func ParseArgsNewRound(args []Arg) []Arg {
	return []Arg{}
}
func ParseArgsVote(args []Arg) []Arg {
	return []Arg{}
}
func ParseArgsAddNebula(args []Arg) (
	nebulaId string,
	b string) {
	nebulaId = ParseNebulaId(args[0].Value)
	b = ParseSign(args[1].Value)
	return nebulaId, b
}
func ParseArgsDropNebula(args []Arg) (
	nebulaId string) {
	nebulaId = ParseNebulaId(args[0].Value)
	return nebulaId
}
func ParseArgsSignNewConsuls(args []Arg) (
	chainType ChainType,
	roundId uint64,
	sign string) {
	chainType = ParseChainType(args[0].Value)
	roundId = ParseIntValue(args[1].Value)
	sign = ParseSign(args[2].Value)
	return chainType, roundId, sign
}
func ParseArgsSignNewOracles(args []Arg) (
	roundId uint64,
	sign string,
	nebulaId string) {
	roundId = ParseIntValue(args[0].Value)
	sign = ParseSign(args[1].Value)
	nebulaId = ParseNebulaId(args[2].Value)
	return roundId, sign, nebulaId
}
func ParseArgsApproveLastRound(args []Arg) []Arg {
	return []Arg{}
}
