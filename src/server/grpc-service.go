package server

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/SeaOfWisdom/sow_proto/contractor-srv"
	"google.golang.org/grpc/reflection"

	"github.com/SeaOfWisdom/sow_contractor/src/config"
	srv "github.com/SeaOfWisdom/sow_contractor/src/service"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	proto.UnsafeContractorServiceServer
	listener net.Listener
	service  *srv.Service
	server   *grpc.Server
}

func NewGrpcServer(config *config.Config, service *srv.Service) *GrpcServer {
	listener, err := net.Listen("tcp", config.GrpcAddress)
	if err != nil {
		log.Fatalf("could not listen to the address %s, err: %v", config.GrpcAddress, err)
	}
	serverOps := []grpc.ServerOption{
		grpc.StreamInterceptor(grpcprometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpcprometheus.UnaryServerInterceptor),
	}
	instance := &GrpcServer{
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

func (gs *GrpcServer) PublishWork(ctx context.Context, req *proto.PublishWorkRequest) (*proto.TxHashResponse, error) {
	return nil, nil
}

// MakeAdmin ...
func (gs *GrpcServer) MakeAuthor(ctx context.Context, req *proto.AccountRequest) (*proto.TxHashResponse, error) {
	txHash, err := gs.service.MakeAuthor(req.Address)
	if err != nil {
		return nil, err
	}

	return &proto.TxHashResponse{
		TxHash: txHash,
	}, nil
}

func (gs *GrpcServer) MakeReviewer(ctx context.Context, req *proto.AccountRequest) (*proto.TxHashResponse, error) {
	txHash, err := gs.service.MakeReviewer(req.Address)
	if err != nil {
		return nil, err
	}

	return &proto.TxHashResponse{
		TxHash: txHash,
	}, nil
}

func (gs *GrpcServer) MakeAdmin(ctx context.Context, req *proto.AccountRequest) (*proto.TxHashResponse, error) {
	txHash, err := gs.service.MakeAdmin(req.Address)
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

func (gs *GrpcServer) PurchaseWork(ctx context.Context, req *proto.PurchaseWorkRequest) (*proto.PurchaseWorkResponse, error) {
	readerTxHash, authorTxHash, err := gs.service.PurchaseWork(req.ReaderAddress, req.AuthorAddress, req.Price)
	if err != nil {
		return nil, err
	}

	return &proto.PurchaseWorkResponse{
		ReaderTxStatus: &proto.TxHashResponse{
			TxHash: readerTxHash,
		},
		AuthorTxStatus: &proto.TxHashResponse{
			TxHash: authorTxHash,
		},
	}, nil
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

func (gs *GrpcServer) Start() {
	go func() {
		fmt.Println("Contractor service started at ", gs.listener.Addr().String())
		if err := gs.server.Serve(gs.listener); err != nil {
			panic(fmt.Errorf("failed to serve gRPC: %v", err))
		}
	}()
}

func (gs *GrpcServer) Stop() {
	gs.server.Stop()
	_ = gs.listener.Close()
}
