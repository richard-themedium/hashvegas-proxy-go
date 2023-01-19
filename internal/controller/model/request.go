package model

type InitGenesisLedgerForTableRequest struct {
	TableId string
}

type GetGameLastStateRequest struct {
	TableId string
	GameSeq uint64
}

type GetGameHistoryRequest struct {
	TableId string
	GameSeq uint64
}

type GetDeckInitStateRequest struct {
	TableId string
	GameSeq uint64
}

type AskNewCardsRequest struct {
	TableId         string
	GameSeq         uint64
	DealSeq         uint8
	RequestCardList []map[string]interface{}
}
