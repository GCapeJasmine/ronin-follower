package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GCapeJasmine/ronin-follower/config"
	"github.com/GCapeJasmine/ronin-follower/infras/jsonrpc/ronin"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/usecases"
	"github.com/GCapeJasmine/ronin-follower/repos"
	"github.com/GCapeJasmine/ronin-follower/server"
	"github.com/GCapeJasmine/ronin-follower/worker"
)

func main() {
	var configFile, port string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.StringVar(&port, "port", "", "Specify port")
	flag.Parse()

	//load config
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Printf("load config fail with err: %v", err)
		panic(err)
	}

	repo := repos.NewRepoInMem(cfg.InMem)
	repoBlock := repo.Blocks()
	roninClient := ronin.NewRoninClient(cfg.RoninConfig)
	blockUsecase := usecases.NewBlock(cfg, repoBlock, roninClient)

	s := server.NewServer(cfg)
	domains := &server.Domains{
		Block: blockUsecase,
	}
	s.InitCORS()
	s.InitRouter(domains)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	w := worker.NewWorker(cfg, blockUsecase)
	go func() {
		w.RunJob(ctx)
	}()

	if err := s.ListenHTTP(); err != nil {
		log.Printf("start server fail with err: %v", err)
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Printf("Receive os.Signal: %s", sig.String())
}
