package models

import "github.com/GCapeJasmine/ronin-follower/internal/domains/models"

type GetTransactionRequest struct {
	Hash string `form:"hash"`
}

type ListBlockTransactionRequest struct {
	BlockNumber string `form:"block_number"`
}

type ListTransactionsInRangeRequest struct {
	From int `form:"from"`
	To   int `form:"to"`
}

type ListTransactionsInRangeResponse struct {
	Transactions []*models.Transaction `json:"transactions"`
	Total        int                   `json:"total"`
}

type GetPercentageTransactionGasFeeRequest struct {
	GasFee float64 `form:"gas_fee"`
}

type GetPercentageTransactionGasFeeResponse struct {
	Result float64 `json:"result"`
}
