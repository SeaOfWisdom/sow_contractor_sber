package server

import (
	"context"

	proto "github.com/SeaOfWisdom/sow_proto/contractor-srv"
)

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

func (gs *GrpcServer) PublishWork(ctx context.Context, req *proto.PublishWorkRequest) (*proto.TxHashResponse, error) {
	txHash, err := gs.service.PublishPaper(req.Authors, req.Name, req.Uri, req.WorkId, req.Price)
	if err != nil {
		return nil, err
	}

	return &proto.TxHashResponse{
		TxHash: txHash,
	}, nil
}

func (gs *GrpcServer) GetPaperById(ctx context.Context, req *proto.PaperByIdRequest) (*proto.PaperByIdResponse, error) {
	paperAddress, err := gs.service.GetPaperById(req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.PaperByIdResponse{
		Address: paperAddress.Hex(),
	}, nil
}
