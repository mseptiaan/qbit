package handler

import (
	"context"
	"encoding/json"
	"github.com/mseptian/qbit/internal/pb"
	"github.com/mseptian/qbit/internal/service"
	"github.com/mseptian/qbit/pkg/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type CartHandlerGrpc struct {
	pb.UnimplementedCartServer
	cart service.CartService
}

func NewCartHandlerGrpcHandler(cart service.CartService) *CartHandlerGrpc {
	return &CartHandlerGrpc{
		cart: cart,
	}
}

func (s *CartHandlerGrpc) GetCart(ctx context.Context, input *pb.Empty) (*pb.MsgArrProductResp, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	values := md["authorization"]
	accessToken := values[0]
	createdUser, err := s.cart.GetCart(ctx, strings.TrimPrefix(accessToken, "Bearer "), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return createdUser, nil
}

func (s *CartHandlerGrpc) PostAddToCart(ctx context.Context, input *pb.MsrReqPostAddToCart) (*pb.DataEmpty, error) {
	validate := validator.NewValidator()
	if err := validate.Struct(input); err != nil {
		errorData, _ := json.Marshal(validator.ValidatorErrors(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", string(errorData))
	}

	md, _ := metadata.FromIncomingContext(ctx)
	values := md["authorization"]
	accessToken := values[0]
	createdUser, err := s.cart.PostAddToCart(ctx, strings.TrimPrefix(accessToken, "Bearer "), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return createdUser, nil
}

func (s *CartHandlerGrpc) GetCheckout(ctx context.Context, input *pb.Empty) (*pb.DataEmpty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	values := md["authorization"]
	accessToken := values[0]

	createdUser, err := s.cart.GetCheckout(ctx, strings.TrimPrefix(accessToken, "Bearer "), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return createdUser, nil
}
