package updater

import (
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/decoder"
	"github.com/Gravity-Tech/gravity-node-api/model"
)

func (updater *NodesCacheUpdater) CacheArgs(
	txFunc decoder.TxFunc,
	args []decoder.Arg,
	txId uint64) error {
	var err error

	switch txFunc {
	case decoder.Commit:
		err = updater.CacheArgsCommit(args, txId)
	case decoder.Reveal:
		err = updater.CacheArgsReveal(args, txId)
	case decoder.Result:
		err = updater.CacheArgsResult(args, txId)
	case decoder.AddOracleInNebula:
		err = updater.CacheArgsAddOracleInNebula(args, txId)
	case decoder.AddOracle:
		err = updater.CacheArgsAddOracle(args, txId)
	case decoder.NewRound:
		err = updater.CacheArgsNewRound(args, txId)
	case decoder.Vote:
		err = updater.CacheArgsVote(args, txId)
	case decoder.AddNebula:
		err = updater.CacheArgsAddNebula(args, txId)
	case decoder.DropNebula:
		err = updater.CacheArgsDropNebula(args, txId)
	case decoder.SignNewConsuls:
		err = updater.CacheArgsSignNewConsuls(args, txId)
	case decoder.SignNewOracles:
		err = updater.CacheArgsSignNewOracles(args, txId)
	case decoder.ApproveLastRound:
		err = updater.CacheArgsApproveLastRound(args, txId)
	default:
		fmt.Println("func not found")
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

func (updater *NodesCacheUpdater) CacheArgsCommit(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	nebulaId, pulseId, tcHeight, commit, oraclePubKey := decoder.ParseArgsCommit(args)
	metadata := model.CommitTransaction{
		TxId:         txId,
		NebulaId:     nebulaId,
		PulseId:      pulseId,
		TcHeight:     tcHeight,
		Commit:       commit,
		OraclePubKey: oraclePubKey}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsReveal(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	commit, nebulaId, pulseId, tcHeight, oraclePubKey, rq, amount, receiver := decoder.ParseArgsReveal(args)
	metadata := model.RevealTransaction{
		TxId:         txId,
		Commit:       commit,
		NebulaId:     nebulaId,
		PulseId:      pulseId,
		TcHeight:     tcHeight,
		OraclePubKey: oraclePubKey,
		Rq:           rq,
		Amount:       amount,
		Receiver:     receiver}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsResult(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	nebulaId, pulseId, sign, chainType, oraclePubKey := decoder.ParseArgsResult(args)
	metadata := model.ResultTransaction{
		TxId:         txId,
		NebulaId:     nebulaId,
		PulseId:      pulseId,
		Sign:         sign,
		ChainType:    chainType.String(),
		OraclePubKey: oraclePubKey}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsAddOracleInNebula(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	nebulaId, oraclePubKey := decoder.ParseArgsAddOracleInNebula(args)
	metadata := model.AddOracleInNebulaTransaction{
		TxId:         txId,
		NebulaId:     nebulaId,
		OraclePubKey: oraclePubKey}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsAddOracle(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	chainType, oraclePubKey := decoder.ParseArgsAddOracle(args)
	metadata := model.AddOracleTransaction{
		TxId:         txId,
		ChainType:    chainType.String(),
		OraclePubKey: oraclePubKey}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsNewRound(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	metadata := model.NewRoundTransaction{TxId: txId}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsVote(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	metadata := model.VoteTransaction{TxId: txId}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsAddNebula(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	nebulaId, b := decoder.ParseArgsAddNebula(args)
	metadata := model.AddNebulaTransaction{
		TxId:     txId,
		NebulaId: nebulaId,
		B:        b}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsDropNebula(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	nebulaId := decoder.ParseArgsDropNebula(args)
	metadata := model.DropNebulaTransaction{
		TxId:     txId,
		NebulaId: nebulaId}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsSignNewConsuls(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	chainType, roundId, sign := decoder.ParseArgsSignNewConsuls(args)
	metadata := model.SignNewConsulsTransaction{
		TxId:      txId,
		ChainType: chainType.String(),
		RoundId:   roundId,
		Sign:      sign}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsSignNewOracles(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	roundId, sign, nebulaId := decoder.ParseArgsSignNewOracles(args)
	metadata := model.SignNewOraclesTransaction{
		TxId:     txId,
		RoundId:  roundId,
		Sign:     sign,
		NebulaId: nebulaId}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
func (updater *NodesCacheUpdater) CacheArgsApproveLastRound(args []decoder.Arg, txId uint64) error {
	var err error
	db := updater.DB.DB
	metadata := model.ApproveLastRoundTransaction{TxId: txId}
	_, err = db.Model(&metadata).
	  Where("tx_id = ?tx_id").
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
