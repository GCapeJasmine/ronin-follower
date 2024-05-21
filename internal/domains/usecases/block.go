package usecases

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/GCapeJasmine/ronin-follower/config"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/client"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/models"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/repos"
	"github.com/GCapeJasmine/ronin-follower/utils"
)

const (
	Precision = 2
)

type Block struct {
	cfg         *config.AppConfig
	blockRepo   repos.IBlockRepo
	roninClient client.IClientRonin
}

func NewBlock(cfg *config.AppConfig, blockRepo repos.IBlockRepo, roninClient client.IClientRonin) *Block {
	return &Block{
		cfg:         cfg,
		blockRepo:   blockRepo,
		roninClient: roninClient,
	}
}

func (b *Block) GetTransactionByHash(ctx context.Context, transactionHash string) (*models.Transaction, error) {
	log.Printf("GetTransactionByHash with hash = %v", transactionHash)

	transaction, err := b.blockRepo.GetTransactionByHash(ctx, transactionHash)
	if err != nil {
		log.Printf("GetTransactionByHash got error = %v", err)
		return nil, err
	}

	if transaction == nil {
		resp, err := b.roninClient.GetTransactionByHash(ctx, &models.GetTransactionByHashInput{
			Params: []any{
				transactionHash,
			},
		})
		if err != nil {
			log.Printf("GetTransactionByHash got error = %v", err)
			return nil, err
		}

		return resp.ToTransaction(), nil
	}

	return transaction, nil
}

func (b *Block) ListTransactionByBlockNumber(ctx context.Context, blockNumber string) ([]*models.Transaction, error) {
	log.Printf("ListTransactionByBlockNumber with blockNumber = %v", blockNumber)

	block, err := b.blockRepo.GetBlockByBlockNumber(ctx, blockNumber)
	if err != nil {
		log.Printf("ListTransactionByBlockNumber got error = %v", err)
		return nil, err
	}

	if block == nil {
		resp, err := b.roninClient.GetBlockByNumber(ctx, &models.GetBlockByNumberInput{
			Params: []any{
				blockNumber,
				true,
			},
		})
		if err != nil {
			log.Printf("ListTransactionByBlockNumber got error = %v", err)
			return nil, err
		}

		return resp.ToBlock().Transactions, nil
	}

	return block.Transactions, nil
}

func (b *Block) ListTransactionInRange(ctx context.Context, from, to int) ([]*models.Transaction, int, error) {
	log.Printf("ListTransactionWithFromAndTo with from %d to %d", from, to)

	transactions, err := b.blockRepo.ListTransactionInRange(ctx, from, to)
	if err != nil {
		log.Printf("ListTransactionInRange got error = %v", err)
		return nil, 0, err
	}

	numberOfTransactions, _ := b.blockRepo.GetNumberOfTransactions(ctx)

	return transactions, numberOfTransactions, nil
}

func (b *Block) GetPercentageOfTransactionsWhichHaveGasFeeLessThan(ctx context.Context, gasFee float64) (float64, error) {
	log.Printf("GetPercentageOfTransactionsWhichHaveGasFeeLessThan ...")

	amount, _ := b.blockRepo.GetNumberOfTransactionsWhichHaveGasFeeLessThan(ctx, gasFee)
	numberOfTransactions, _ := b.blockRepo.GetNumberOfTransactions(ctx)
	if amount == 0 || numberOfTransactions == 0 {
		return 0, errors.New("integer divide by zero")
	}
	res := utils.Round(float64(amount*100/numberOfTransactions), Precision)

	return res, nil
}

func (b *Block) InsertBlock(ctx context.Context) error {
	log.Printf("InsertNewBlock ...")

	latestBlock, _ := b.blockRepo.GetLatestBlock(ctx)
	nextBlockNumber := ""
	// get first block number
	if len(latestBlock.Number) == 0 {
		firstBlockNumber, err := b.getFirstBlockNumber(ctx)
		if err != nil {
			log.Printf("getFirstBlockNumber got err = %v", err)
			return err
		}
		nextBlockNumber = firstBlockNumber
	} else {
		nextBlockNumber = b.getNextBlockNumber(latestBlock)
	}
	fmt.Println(nextBlockNumber)
	// insert block
	blockResp, err := b.roninClient.GetBlockByNumber(ctx, &models.GetBlockByNumberInput{
		Params: []any{
			nextBlockNumber,
			true,
		},
	})
	if err != nil {
		log.Printf("GetBlockByNumber got err = %v", err)
		return err
	}
	err = b.blockRepo.InsertBlock(ctx, blockResp.ToBlock())
	if err != nil {
		log.Printf("InsertBlock got err = %v", err)
		return err
	}
	return nil
}

func (b *Block) getFirstBlockNumber(ctx context.Context) (string, error) {
	log.Printf("getFirstBlockNumber ...")

	res, err := b.roninClient.GetLatestBlockNumber(ctx, &models.GetLatestBlockNumberInput{})
	if err != nil {
		log.Printf("GetLatestBlockNumber got err = %v", err)
		return "", err
	}

	return res.Result, nil
}

func (b *Block) getNextBlockNumber(latestBlock *models.Block) string {
	latestBlockNumberInt64 := utils.HexWithOxPrefixToInt64(latestBlock.Number)
	return utils.Int64ToHexWith0xPrefix(latestBlockNumberInt64 + 1)
}
