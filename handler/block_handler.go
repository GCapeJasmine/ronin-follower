package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GCapeJasmine/ronin-follower/handler/models"
)

func (h *Handler) getTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &models.GetTransactionRequest{}
		if err := ctx.ShouldBind(request); err != nil {
			log.Printf("parse request with err = %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := h.block.GetTransactionByHash(ctx, request.Hash)
		if err != nil {
			log.Printf("GetSiteUsecase fail with err = %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func (h *Handler) listBlockTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &models.ListBlockTransactionRequest{}
		if err := ctx.ShouldBind(request); err != nil {
			log.Printf("parse request with err = %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := h.block.ListTransactionByBlockNumber(ctx, request.BlockNumber)
		if err != nil {
			log.Printf("listBlockTransaction got err = %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func (h *Handler) listTransactionInRange() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &models.ListTransactionsInRangeRequest{}
		if err := ctx.ShouldBind(request); err != nil {
			log.Printf("parse request with err = %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, total, err := h.block.ListTransactionInRange(ctx, request.From, request.To)
		if err != nil {
			log.Printf("listTransactionInRange got err = %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &models.ListTransactionsInRangeResponse{
			Transactions: res,
			Total:        total,
		})
	}
}

func (h *Handler) getPercentageTransactionGasFee() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &models.GetPercentageTransactionGasFeeRequest{}
		if err := ctx.ShouldBind(request); err != nil {
			log.Printf("parse request with err = %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := h.block.GetPercentageOfTransactionsWhichHaveGasFeeLessThan(ctx, request.GasFee)
		if err != nil {
			log.Printf("GetPercentageOfTransactionsWhichHaveGasFeeLessThan got err = %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &models.GetPercentageTransactionGasFeeResponse{
			Result: res,
		})
	}
}
