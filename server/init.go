package server

import (
	"github.com/gin-contrib/cors"

	hdl "github.com/GCapeJasmine/ronin-follower/handler"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/usecases"
)

type Domains struct {
	Block *usecases.Block
}

func (s *Server) InitCORS() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Access-Token",
		"X-Google-Access-Token",
	}
	s.router.Use(cors.New(corsConfig))
}

func (s *Server) InitRouter(domains *Domains) {
	// init handler
	handler := hdl.NewHandler(domains.Block)

	// API groups
	router := s.router.Group("v1")
	handler.ConfigRouteAPI(router)
}
