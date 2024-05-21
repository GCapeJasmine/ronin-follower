package usecases

import (
	"context"
	"testing"

	"github.com/GCapeJasmine/ronin-follower/config"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/models"
	"github.com/GCapeJasmine/ronin-follower/repos"
	"github.com/stretchr/testify/require"
)

func TestBlock_GetPercentageOfTransactionsWhichHaveGasFeeLessThan(t *testing.T) {
	blockRepo := repos.NewBlockRepo(&config.InMemConfig{
		Capacity: 5,
	})
	ctx := context.Background()
	_ = blockRepo.InsertBlock(ctx, &models.Block{
		Transactions: []*models.Transaction{
			{
				Hash:     "0",
				Gas:      "0x19107",
				GasPrice: "0x59682f000",
			},
			{
				Hash:     "1",
				Gas:      "0x123fe",
				GasPrice: "0x59682f000",
			},
			{
				Hash:     "2",
				Gas:      "0x19107",
				GasPrice: "0x59682f000",
			},
			{
				Hash:     "3",
				Gas:      "0x123fe",
				GasPrice: "0x59682f000",
			},
			{
				Hash:     "4",
				Gas:      "0x19107",
				GasPrice: "0x59682f000",
			},
			{
				Hash:     "5",
				Gas:      "0x123fe",
				GasPrice: "0x59682f000",
			},
			{
				Hash:     "6",
				Gas:      "0x123fe",
				GasPrice: "0x59682f000",
			},
		},
	})
	res, err := blockRepo.GetNumberOfTransactionsWhichHaveGasFeeLessThan(ctx, float64(24000000000*102663))
	total, _ := blockRepo.GetNumberOfTransactions(ctx)
	require.NoError(t, err)
	require.Equal(t, total, 7)
	require.Equal(t, res, 4)
}

func TestBlock_ListTransactionInRange(t *testing.T) {
	blockRepo := repos.NewBlockRepo(&config.InMemConfig{
		Capacity: 5,
	})
	ctx := context.Background()
	_ = blockRepo.InsertBlock(ctx, &models.Block{
		Transactions: []*models.Transaction{
			{
				Hash: "0",
			},
			{
				Hash: "1",
			},
			{
				Hash: "2",
			},
			{
				Hash: "3",
			},
			{
				Hash: "4",
			},
			{
				Hash: "5",
			},
		},
	})

	_ = blockRepo.InsertBlock(context.Background(), &models.Block{
		Transactions: []*models.Transaction{
			{
				Hash: "6",
			},
			{
				Hash: "7",
			},
			{
				Hash: "8",
			},
			{
				Hash: "9",
			},
			{
				Hash: "10",
			},
		},
	})
	res, err := blockRepo.ListTransactionInRange(ctx, 4, 8)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, "6", res[0].Hash)
	require.Equal(t, "2", res[len(res)-1].Hash)
}
