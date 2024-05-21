package worker

import (
	"context"
	"log"
	"time"

	"github.com/GCapeJasmine/ronin-follower/config"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/usecases"
)

type Worker struct {
	cfg          *config.AppConfig
	blockUsecase *usecases.Block
}

func NewWorker(cfg *config.AppConfig, blockUsecase *usecases.Block) *Worker {
	return &Worker{
		cfg:          cfg,
		blockUsecase: blockUsecase,
	}
}

func (w *Worker) RunJob(ctx context.Context) {
	interval := time.Duration(w.cfg.Worker.Interval) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Job canceled")
			return
		case <-ticker.C:
			log.Printf("Job running at: %v", time.Now())
			err := w.blockUsecase.InsertBlock(ctx)
			if err != nil {
				log.Printf("worker InsertBlock got err = %v", err)
			}
		}
	}

}
