package controller

import (
	"net/http"

	"github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/controller/model"
	// model_dto "github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/dto/model"
	"github.com/gin-gonic/gin"
	"github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/usecase"
	model_usecase "github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/usecase/model"
)

const (
	StatusOk = 200
)

type HvController struct {
	hvService usecase.HashVegasService
	router    *gin.Engine
}

func NewHvController() *HvController {
	hvService := usecase.NewHashVegasService()
	return &HvController{
		hvService: hvService,
	}
}

func (h *HvController) InitGenesisLedgerForTable(c *gin.Context) {
	var request model.InitGenesisLedgerForTableRequest
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ht, err := h.hvService.InitGenesisLedgerForTable(request.TableId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &model.InitGenesisLedgerForTableResponse{
		Code: StatusOk,
		Data: ht,
	})
	return
}

func (h *HvController) GetGameLastState(c *gin.Context) {
	var request model.GetDeckInitStateRequest
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ht, err := h.hvService.GetGameLastState(&model_usecase.GetGameLastStateParam{
		TableId: request.TableId,
		GameSeq: request.GameSeq,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &model.GetGameLastStateResponse{
		Code: StatusOk,
		Data: ht,
	})
	return
}

func (h *HvController) GetGameHistory(c *gin.Context) {
	var request model.GetGameHistoryRequest
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ht, err := h.hvService.GetGameHistory(&model_usecase.GetGameHistoryParam{
		TableId: request.TableId,
		GameSeq: request.GameSeq,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &model.GetGameHistoryResponse{
		Code: StatusOk,
		Data: ht,
	})
	return
}

func (h *HvController) GetDeckInitState(c *gin.Context) {
	var request model.GetDeckInitStateRequest
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ht, err := h.hvService.GetDeckInitState(&model_usecase.GetDeckInitStateParam{
		TableId: request.TableId,
		GameSeq: request.GameSeq,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &model.GetDeckInitStateResponse{
		Code: StatusOk,
		Data: ht,
	})

	return
}

func (h *HvController) AskNewCards(c *gin.Context) {
	var request model.AskNewCardsRequest
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ht, err := h.hvService.AskNewCards(&model_usecase.AskNewCardsParam{
		TableId:         request.TableId,
		GameSeq:         request.GameSeq,
		DealSeq:         request.DealSeq,
		RequestCardList: request.RequestCardList,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &model.AskNewCardsResponse{
		Code: StatusOk,
		Data: ht,
	})

	return
}
