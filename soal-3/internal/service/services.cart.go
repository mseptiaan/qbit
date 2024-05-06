package service

import (
	"context"
	"errors"
	"github.com/mseptian/qbit/internal/models"
	"github.com/mseptian/qbit/internal/pb"
	"github.com/mseptian/qbit/internal/repository"
	"github.com/mseptian/qbit/pkg/auth"
	"github.com/mseptian/qbit/pkg/hash"
)

type ServiceCart struct {
	Repository repository.Repositories
	Hash       hash.Hashing
	Token      auth.JWTManager
}

func NewServiceCart(auth repository.Repositories, hash hash.Hashing, token auth.JWTManager) *ServiceCart {
	return &ServiceCart{Repository: auth, Hash: hash, Token: token}
}

func (s *ServiceCart) GetCart(ctx context.Context, token string, input *pb.Empty) (*pb.MsgArrProductResp, error) {
	jwt, _ := s.Token.Verify(token)
	cart, _ := s.Repository.Cart.GetAll(ctx, jwt.Audience)

	var items []*pb.MsgProductResp
	total := 0
	for _, all := range cart {
		items = append(items, &pb.MsgProductResp{
			Uuid:  all.ProductId,
			Name:  all.Name,
			Qty:   uint32(all.CartQty),
			Price: uint32(all.Price),
		})
		total += all.Price * all.CartQty
	}

	return &pb.MsgArrProductResp{
		Items: items,
		Total: uint32(total),
	}, nil
}

func (s *ServiceCart) PostAddToCart(ctx context.Context, token string, input *pb.MsrReqPostAddToCart) (*pb.DataEmpty, error) {

	jwt, _ := s.Token.Verify(token)

	cart := &models.Cart{
		UserId:    jwt.Audience,
		ProductId: input.ProductId,
		Qty:       int(input.Qty),
	}

	_, rowsAffected := s.Repository.Cart.GetByUserIdAndProductId(ctx, jwt.Audience, input.ProductId)

	if rowsAffected < 1 {
		err := s.Repository.Cart.Save(ctx, cart)
		if err != nil {
			return nil, err
		}
	} else {
		err := s.Repository.Cart.Update(ctx, jwt.Audience, input.ProductId, cart)
		if err != nil {
			return nil, err
		}
	}
	return &pb.DataEmpty{}, nil
}

func (s *ServiceCart) GetCheckout(ctx context.Context, token string, input *pb.Empty) (*pb.DataEmpty, error) {
	jwt, _ := s.Token.Verify(token)
	cart, _ := s.Repository.Cart.GetAll(ctx, jwt.Audience)

	var items []models.Product
	for _, all := range cart {
		qtyFinal := all.ProductQty - all.CartQty
		if qtyFinal < 0 {
			return &pb.DataEmpty{}, errors.New("one of the quantity products is not correct")
		}
		items = append(items, models.Product{
			ProductId: all.ProductId,
			Qty:       qtyFinal,
			Name:      all.Name,
			Price:     all.Price,
		})
	}
	s.Repository.Product.UpdateMultiple(ctx, &items)
	s.Repository.Cart.RemoveByUserId(ctx, jwt.Audience)
	return &pb.DataEmpty{}, nil
}
