package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/GCapeJasmine/ronin-follower/internal/domains/usecases"
)

type Handler struct {
	block *usecases.Block
}

func NewHandler(block *usecases.Block) *Handler {
	return &Handler{
		block: block,
	}
}

func (h *Handler) ConfigRouteAPI(router *gin.RouterGroup) {
	// transactions
	router.GET("/transactions", h.getTransaction())
	router.GET("/transactions/list", h.listTransactionInRange())
	router.GET("/transactions/gas", h.getPercentageTransactionGasFee())
	// blocks
	router.GET("/blocks/transactions", h.listBlockTransaction())

}
