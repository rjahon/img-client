package main

import (
	"github.com/rjahon/img-client/api"
	"github.com/rjahon/img-client/api/handlers"
	"github.com/rjahon/img-client/config"
	"github.com/rjahon/img-client/grpc/client"
	"github.com/rjahon/img-client/logger"
)

func main() {
	cfg := config.Load()

	log := logger.NewLogger(cfg.ServiceName)
	defer logger.Cleanup(log)

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	h := handlers.NewHandler(cfg, log, svcs)
	r := api.SetUpRouter(h, cfg)
	log.Info("HTTP started on port ", logger.String("port", cfg.HTTPPort))

	r.Run(cfg.HTTPPort)
}
