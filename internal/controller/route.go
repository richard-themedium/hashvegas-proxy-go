package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/the-medium-tech/platform-externals/log"
)

func (h *HvController) RegisterRoute() error {
	r := gin.Default()

	r.GET("/1", h.InitGenesisLedgerForTable)
	r.GET("/2", h.GetGameLastState)
	r.GET("/3", h.GetGameHistory)
	r.GET("/4", h.GetDeckInitState)
	r.GET("/5", h.AskNewCards)

	h.router = r
	return nil
}

func (h *HvController) Start(addr string) error {
	err := h.router.Run(addr)
	if err != nil {
		log.Error("Running the rest service failed: ", err.Error())
		return err
	}
	return nil
}
