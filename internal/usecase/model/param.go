package model

type GetGameLastStateParam struct {
	TableId string
	GameSeq uint64
}

type GetGameHistoryParam struct {
	TableId string
	GameSeq uint64
}

type GetDeckInitStateParam struct {
	TableId string
	GameSeq uint64
}

type AskNewCardsParam struct {
	TableId         string
	GameSeq         uint64
	DealSeq         uint8
	RequestCardList []map[string]interface{}
}
