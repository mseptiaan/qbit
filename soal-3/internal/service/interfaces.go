package service

import (
	"context"
	"github.com/mseptian/qbit/internal/pb"
)

type AuthService interface {
	Login(ctx context.Context, input *pb.MsgReqPostLogin) (*pb.MsgRespPostLogin, error)
	Register(ctx context.Context, input *pb.MsgReqPostRegister) (*pb.DataEmpty, error)
}

type UserService interface {
	Profile(ctx context.Context, token string) (*pb.MsgResGetUser, error)
}

type ProductService interface {
	GetProduct(ctx context.Context, input *pb.Empty) (*pb.MsgArrProduct, error)
	GetProductSearch(ctx context.Context, input *pb.MsgReqGetProductSearch) (*pb.MsgArrProduct, error)
	GetProductId(ctx context.Context, input *pb.MsgReqGetProductId) (*pb.MsgProduct, error)
}

type CartService interface {
	GetCart(ctx context.Context, token string, input *pb.Empty) (*pb.MsgArrProductResp, error)
	PostAddToCart(ctx context.Context, token string, input *pb.MsrReqPostAddToCart) (*pb.DataEmpty, error)
	GetCheckout(ctx context.Context, token string, input *pb.Empty) (*pb.DataEmpty, error)
}
