package handler

import (
	"context"
	"github.com/mseptian/qbit/internal/pb"
	"github.com/mseptian/qbit/internal/service"
	"google.golang.org/grpc/metadata"
	"strings"
)

type UserHandlerGrpc struct {
	pb.UnimplementedUserServer
	user service.UserService
}

func NewUserHandlerGrpcHandler(user service.UserService) *UserHandlerGrpc {
	return &UserHandlerGrpc{
		user: user,
	}
}

func (s *UserHandlerGrpc) GetUser(ctx context.Context, request *pb.Empty) (*pb.MsgResGetUser, error) {

	md, _ := metadata.FromIncomingContext(ctx)
	values := md["authorization"]
	accessToken := values[0]

	response, err := s.user.Profile(ctx, strings.TrimPrefix(accessToken, "Bearer "))
	return response, err
}
