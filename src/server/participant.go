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
