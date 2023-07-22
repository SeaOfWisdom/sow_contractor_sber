package server

import (
	"context"
	"fmt"
	"net"

	"github.com/SeaOfWisdom/sow_contractor/src/config"
	"github.com/SeaOfWisdom/sow_contractor/src/log"
	srv "github.com/SeaOfWisdom/sow_contractor/src/service"
	proto "github.com/SeaOfWisdom/sow_proto/contractor-srv"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	proto.UnsafeContractorServiceServer
	logger   *log.Logger
	listener net.Listener
	service  *srv.Service
	server   *grpc.Server
}

func NewGrpcServer(config *config.Config, logger *log.Logger, service *srv.Service) *GrpcServer {
	listener, err := net.Listen("tcp", config.GrpcAddress)
	if err != nil {
		logger.Fatalf("could not listen to the address %s, err: %v", config.GrpcAddress, err)
	}

	serverOps := []grpc.ServerOption{
		grpc.StreamInterceptor(grpcprometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpcprometheus.UnaryServerInterceptor),
	}

	instance := &GrpcServer{
		logger:   logger,
		service:  service,
		server:   grpc.NewServer(serverOps...),
		listener: listener,
	}
	grpcprometheus.EnableHandlingTimeHistogram()
	grpcprometheus.Register(instance.server)
	proto.RegisterContractorServiceServer(instance.server, instance)

	reflection.Register(instance.server)

	return instance
}

func (gs *GrpcServer) Faucet(ctx context.Context, req *proto.FaucetRequest) (*proto.TxHashResponse, error) {
	txHash, err := gs.service.Faucet(req.Address, req.Amount)
	if err != nil {
		return nil, err
	}

	return &proto.TxHashResponse{
		TxHash: txHash,
	}, nil
}

func (gs *GrpcServer) GetStatus(ctx context.Context, req *proto.TxStatusRequest) (*proto.TxStatusResponse, error) {
	return nil, nil
}

func (gs *GrpcServer) Start() {
	go func() {
		gs.logger.Infof("Contractor service started at %s", gs.listener.Addr().String())
		if err := gs.server.Serve(gs.listener); err != nil {
			panic(fmt.Errorf("failed to serve gRPC: %v", err))
		}
	}()
}

func (gs *GrpcServer) Stop() {
	gs.server.Stop()
	_ = gs.listener.Close()
}
