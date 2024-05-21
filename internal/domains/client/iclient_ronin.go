package client

import (
	"context"

	"github.com/GCapeJasmine/ronin-follower/internal/domains/models"
)

type IClientRonin interface {
	GetLatestBlockNumber(ctx context.Context, input *models.GetLatestBlockNumberInput) (*models.GetLatestBlockNumberOutput, error)
	GetBlockByNumber(ctx context.Context, input *models.GetBlockByNumberInput) (*models.GetBlockByNumberOutput, error)
	GetTransactionByHash(ctx context.Context, input *models.GetTransactionByHashInput) (*models.GetTransactionByHashOutput, error)
}
