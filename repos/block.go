package repos

import (
	"context"
	"fmt"
	"github.com/GCapeJasmine/ronin-follower/utils"
	"sync"

	"github.com/GCapeJasmine/ronin-follower/config"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/models"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/repos"
)

type blockRepo struct {
	capacity             int
	listBlock            *models.ListBlock
	hashToTransactionMap map[string]*models.Transaction
	gasPriceCounter      map[float64]int
	mu                   sync.Mutex
}

func NewBlockRepo(cfg *config.InMemConfig) repos.IBlockRepo {
	listBlock := &models.ListBlock{
		Head: &models.Block{},
		Tail: &models.Block{},
		List: make(map[string]*models.Block),
	}
	listBlock.Head.Next = listBlock.Tail
	listBlock.Tail.Prev = listBlock.Head
	return &blockRepo{
		capacity:             cfg.Capacity,
		listBlock:            listBlock,
		hashToTransactionMap: make(map[string]*models.Transaction),
		gasPriceCounter:      make(map[float64]int),
		mu:                   sync.Mutex{},
	}
}

func (b *blockRepo) GetBlockByBlockNumber(ctx context.Context, blockNumber string) (*models.Block, error) {
	if block, found := b.listBlock.List[blockNumber]; found {
		return block, nil
	}
	return nil, nil
}

func (b *blockRepo) GetTransactionByHash(ctx context.Context, transactionHash string) (*models.Transaction, error) {
	if transaction, found := b.hashToTransactionMap[transactionHash]; found {
		return transaction, nil
	}
	return nil, nil
}

func (b *blockRepo) InsertBlock(ctx context.Context, block *models.Block) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.listBlock.Insert(block)
	// handle add transactions
	for _, transaction := range block.Transactions {
		b.hashToTransactionMap[transaction.Hash] = transaction
		gasFee := b.getGasFee(transaction)
		b.gasPriceCounter[gasFee]++
	}
	// handle evicted block
	if len(b.listBlock.List) > b.capacity {
		// handle evicted transactions
		evictedBlock := b.listBlock.Head.Next
		for _, transaction := range evictedBlock.Transactions {
			delete(b.hashToTransactionMap, transaction.Hash)
			gasFee := b.getGasFee(transaction)
			b.gasPriceCounter[gasFee]--
			if b.gasPriceCounter[gasFee] == 0 {
				delete(b.gasPriceCounter, gasFee)
			}
		}
		b.listBlock.RemoveHead()
	}
	fmt.Println(b.listBlock.List)
	return nil
}

func (b *blockRepo) ListTransactionInRange(ctx context.Context, from, to int) ([]*models.Transaction, error) {
	res := make([]*models.Transaction, 0)
	curBlock := b.listBlock.Tail.Prev
	amount := to - from + 1

	// find first transaction in range
	for curBlock != nil && len(curBlock.Transactions) < from {
		curBlock = curBlock.Prev
		from -= len(curBlock.Transactions)
	}
	// get all transactions in range
outerLoop:
	for curBlock != nil {
		for i := len(curBlock.Transactions) - from - 1; i >= 0; i-- {
			res = append(res, curBlock.Transactions[i])
			if len(res) == amount {
				break outerLoop
			}
		}
		from = 0
		curBlock = curBlock.Prev
	}
	return res, nil
}

func (b *blockRepo) GetNumberOfTransactions(ctx context.Context) (int, error) {
	return len(b.hashToTransactionMap), nil
}

func (b *blockRepo) GetNumberOfTransactionsWhichHaveGasFeeLessThan(ctx context.Context, input float64) (int, error) {
	res := 0
	for gasFee, amount := range b.gasPriceCounter {
		if gasFee < input {
			res += amount
		}
	}
	return res, nil
}

func (b *blockRepo) GetLatestBlock(ctx context.Context) (*models.Block, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.listBlock.Tail.Prev, nil
}

func (b *blockRepo) getGasFee(transaction *models.Transaction) float64 {
	gasUsed := utils.HexWithOxPrefixToInt64(transaction.Gas)
	gasPrice := utils.HexWithOxPrefixToInt64(transaction.GasPrice)
	return float64(gasUsed * gasPrice)
}
