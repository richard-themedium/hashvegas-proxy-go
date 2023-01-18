package chaincodes

import (
	model_dto "github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/dto/model"
)

type HashVegasContract interface {
	InitGenesisLedgerForTable(tableId string) (*model_dto.HoldemTableGame, error)
	GetGameLastState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error)
	GetGameHistory(tableId string, gameSeq uint64) ([]*model_dto.HoldemTableGame, error)
	getDeckInitState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error)
	AskNewCards(tableId string, gameSeq uint64, dealSeq uint8, requestCardList []map[string]interface{}) (*model_dto.HoldemServicerResponse, error)
}

func NewHashVegasContract() HashVegasContract {
	return &hashVegasContract{}
}

type hashVegasContract struct {
}

// AskNewCards implements HashVegasContract
func (*hashVegasContract) AskNewCards(tableId string, gameSeq uint64, dealSeq uint8, requestCardList []map[string]interface{}) (*model_dto.HoldemServicerResponse, error) {
	panic("unimplemented")
}

// GetGameHistory implements HashVegasContract
func (*hashVegasContract) GetGameHistory(tableId string, gameSeq uint64) ([]*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

// GetGameLastState implements HashVegasContract
func (*hashVegasContract) GetGameLastState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

// InitGenesisLedgerForTable implements HashVegasContract
func (*hashVegasContract) InitGenesisLedgerForTable(tableId string) (*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

// getDeckInitState implements HashVegasContract
func (*hashVegasContract) getDeckInitState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}
