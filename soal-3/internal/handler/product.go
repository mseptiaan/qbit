package handler

import (
	"context"
	"github.com/mseptian/qbit/internal/pb"
	"github.com/mseptian/qbit/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductHandlerGrpc struct {
	pb.UnimplementedProductServer
	product service.ProductService
}

func NewProductHandlerGrpcHandler(product service.ProductService) *ProductHandlerGrpc {
	return &ProductHandlerGrpc{
		product: product,
	}
}

func (s *ProductHandlerGrpc) GetProduct(ctx context.Context, input *pb.Empty) (*pb.MsgArrProduct, error) {
	createdUser, err := s.product.GetProduct(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return createdUser, nil
}

func (s *ProductHandlerGrpc) GetProductSearch(ctx context.Context, input *pb.MsgReqGetProductSearch) (*pb.MsgArrProduct, error) {
	createdUser, err := s.product.GetProductSearch(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return createdUser, nil
}

func (s *ProductHandlerGrpc) GetProductId(ctx context.Context, input *pb.MsgReqGetProductId) (*pb.MsgProduct, error) {
	createdUser, err := s.product.GetProductId(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return createdUser, nil
}
