package model

// Holdem Card Dealing History to be stored
type HoldemTableGame struct {
	TableId          string         `json:"tableId"` //<PK>
	GameSeq          uint64         `json:"gameSeq"` //<PK> Sequence of New Game (for Tracking Game History) (1,2,3 ...)
	DealSeq          uint8          `json:"dealSeq"` //Sequence of Dealing for each Card (1,2,3 ...)
	CardsRemain      string         `json:"cardsRemain"`
	CardsRemainCount uint8          `json:"cardsRemainCount"`
	RequestCards     []RequestCard  `json:"requestCards,omitempty" metadata:",optional"`  //dealder(banker)도 id를 부여받아야 한다.
	ResponseCards    []ResponseCard `json:"responseCards,omitempty" metadata:",optional"` //@TODO:추후 삭제 필요
	ResponseCardsEnc string         `json:"responseCardsEnc,omitempty" metadata:",optional"`
	InitCardsOrder   []string       `json:"initCardsOrder,omitempty" metadata:",optional"`
	EncKey           string         `json:"encKey,omitempty" metadata:",optional"`
}

type RequestCard struct {
	PlayerId     string `json:"playerId"`
	RequestCount uint8  `json:"requestCount"`
}

type ResponseCard struct {
	PlayerId  string   `json:"playerId"`
	DealCards []string `json:"dealCards"`
}

// Holdem Card Service Response format
type HoldemServicerResponse struct {
	TableId          string         `json:"tableId"` //
	GameSeq          uint64         `json:"gameSeq"` //Sequence of New Game (for Tracking Game History) (1,2,3 ...)
	DealSeq          uint8          `json:"dealSeq"` //Sequence of Dealing for each Card (1,2,3 ...)
	RequestCards     []RequestCard  `json:"requestCards,omitempty" metadata:",optional"`
	ResponseCards    []ResponseCard `json:"responseCards,omitempty" metadata:",optional"`
	ResponseCardsEnc string         `json:"responseCardsEnc,omitempty" metadata:",optional"`
}
