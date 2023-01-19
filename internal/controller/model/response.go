package model

import (
	model_dto "github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/dto/model"
)

type InitGenesisLedgerForTableResponse struct {
	Code int                        `json:"code"`
	Data *model_dto.HoldemTableGame `json:"data"`
}

type GetGameLastStateResponse struct {
	Code int                        `json:"code"`
	Data *model_dto.HoldemTableGame `json:"data"`
}

type GetGameHistoryResponse struct {
	Code int                          `json:"code"`
	Data []*model_dto.HoldemTableGame `json:"data"`
}

type GetDeckInitStateResponse struct {
	Code int                        `json:"code"`
	Data *model_dto.HoldemTableGame `json:"data"`
}

type AskNewCardsResponse struct {
	Code int                               `json:"code"`
	Data *model_dto.HoldemServicerResponse `json:"data"`
}
