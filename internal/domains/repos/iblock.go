package repos

import (
	"context"

	"github.com/GCapeJasmine/ronin-follower/internal/domains/models"
)

type IBlockRepo interface {
	GetBlockByBlockNumber(ctx context.Context, blockNumber string) (*models.Block, error)
	GetTransactionByHash(ctx context.Context, transactionHash string) (*models.Transaction, error)
	InsertBlock(ctx context.Context, block *models.Block) error
	ListTransactionInRange(ctx context.Context, from, to int) ([]*models.Transaction, error)
	GetNumberOfTransactions(ctx context.Context) (int, error)
	GetNumberOfTransactionsWhichHaveGasFeeLessThan(ctx context.Context, input float64) (int, error)
	GetLatestBlock(ctx context.Context) (*models.Block, error)
}
