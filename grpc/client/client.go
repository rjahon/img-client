package client

import (
	"github.com/rjahon/img-client/config"
	"github.com/rjahon/img-client/genproto/img_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	ImgService() img_service.ServiceClient
}

type grpcClients struct {
	imgService img_service.ServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connImgService, err := grpc.Dial(
		cfg.ImgServiceHost+cfg.ImgGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		imgService: img_service.NewServiceClient(connImgService),
	}, nil
}

func (g *grpcClients) ImgService() img_service.ServiceClient {
	return g.imgService
}
