package server

import (
	"fmt"
	"github.com/GCapeJasmine/ronin-follower/config"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
	router     *gin.Engine
	cfg        *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	router := gin.New()
	return &Server{
		router: router,
		cfg:    cfg,
	}
}

func (s *Server) ListenHTTP() error {
	_ = os.Setenv("PORT", "8090")
	listen, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Printf("err %v", err)
		panic(err)
	}
	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	log.Printf("starting http server at port %v ...", os.Getenv("PORT"))

	return s.httpServer.Serve(listen)
}
