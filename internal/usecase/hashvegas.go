package usecase

import (
	model_dto "github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/dto/model"
	"github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/usecase/model"
)

type hashVegas struct {
}

func newHashVegas() *hashVegas {
	return &hashVegas{}
}

// AskNewCards implements HashVegasService
func (*hashVegas) AskNewCards(param *model.AskNewCardsParam) (*model_dto.HoldemServicerResponse, error) {
	panic("unimplemented")
}

// GetDeckInitState implements HashVegasService
func (*hashVegas) GetDeckInitState(param *model.GetDeckInitStateParam) (*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

// GetGameHistory implements HashVegasService
func (*hashVegas) GetGameHistory(param *model.GetGameHistoryParam) ([]*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

// GetGameLastState implements HashVegasService
func (*hashVegas) GetGameLastState(param *model.GetGameLastStateParam) (*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

// InitGenesisLedgerForTable implements HashVegasService
func (*hashVegas) InitGenesisLedgerForTable(tableId string) (*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}
