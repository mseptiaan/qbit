package service

import (
	"context"
	"errors"
	"github.com/mseptian/qbit/internal/pb"
	"github.com/mseptian/qbit/internal/repository"
	"github.com/mseptian/qbit/pkg/auth"
	"github.com/mseptian/qbit/pkg/hash"
)

type ServiceProduct struct {
	Repository repository.Repositories
	Hash       hash.Hashing
	Token      auth.JWTManager
}

func NewServiceProduct(auth repository.Repositories, hash hash.Hashing, token auth.JWTManager) *ServiceProduct {
	return &ServiceProduct{Repository: auth, Hash: hash, Token: token}
}

func (s *ServiceProduct) GetProduct(ctx context.Context, input *pb.Empty) (*pb.MsgArrProduct, error) {
	all, err := s.Repository.Product.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var product []*pb.MsgProduct

	for _, m := range all {
		product = append(product, &pb.MsgProduct{
			Uuid:  m.ProductId,
			Name:  m.Name,
			Qty:   uint32(m.Qty),
			Price: uint32(m.Price),
		})
	}

	return &pb.MsgArrProduct{Items: product}, nil
}

func (s *ServiceProduct) GetProductSearch(ctx context.Context, input *pb.MsgReqGetProductSearch) (*pb.MsgArrProduct, error) {
	all, err := s.Repository.Product.GetBySearch(ctx, input.Search)
	if err != nil {
		return nil, err
	}

	var product []*pb.MsgProduct

	for _, m := range all {
		product = append(product, &pb.MsgProduct{
			Uuid:  m.ProductId,
			Name:  m.Name,
			Qty:   uint32(m.Qty),
			Price: uint32(m.Price),
		})
	}

	return &pb.MsgArrProduct{Items: product}, nil
}

func (s *ServiceProduct) GetProductId(ctx context.Context, input *pb.MsgReqGetProductId) (*pb.MsgProduct, error) {
	product, RowsAffected := s.Repository.Product.GetById(ctx, input.ProductId)
	if RowsAffected < 1 {
		return nil, errors.New("product not found")
	}

	return &pb.MsgProduct{
		Uuid:  product.ProductId,
		Name:  product.Name,
		Qty:   uint32(product.Qty),
		Price: uint32(product.Price),
	}, nil
}
