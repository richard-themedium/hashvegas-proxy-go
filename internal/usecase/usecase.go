package usecase

import (
	model_dto "github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/dto/model"
	"github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/usecase/model"
)

type HashVegasService interface {
	InitGenesisLedgerForTable(tableId string) (*model_dto.HoldemTableGame, error)
	GetGameLastState(param *model.GetGameLastStateParam) (*model_dto.HoldemTableGame, error)
	GetGameHistory(param *model.GetGameHistoryParam) ([]*model_dto.HoldemTableGame, error)
	GetDeckInitState(param *model.GetDeckInitStateParam) (*model_dto.HoldemTableGame, error)
	AskNewCards(param *model.AskNewCardsParam) (*model_dto.HoldemServicerResponse, error)
}

func NewHashVegasService() HashVegasService {
	return newHashVegas()
}
