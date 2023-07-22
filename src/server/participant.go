package server

import (
	"context"

	proto "github.com/SeaOfWisdom/sow_proto/contractor-srv"
)

func (gs *GrpcServer) AddParticipant(ctx context.Context, req *proto.AccountRequest) (*proto.TxHashResponse, error) {
	txHash, err := gs.service.AddParticipant(ctx, req.Address)
	if err != nil {
		return nil, err
	}

	return &proto.TxHashResponse{TxHash: txHash, ErrorMsg: ""}, nil
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

func (gs *GrpcServer) GetParticipantRole(ctx context.Context, req *proto.AccountRequest) (*proto.RoleResponse, error) {
	role, err := gs.service.GetParticipantRole(req.Address)
	if err != nil {
		return nil, err
	}

	return &proto.RoleResponse{
		Role: uint64(role),
	}, nil
}
