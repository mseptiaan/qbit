package handler

import (
	"context"
	"encoding/json"
	"github.com/mseptian/qbit/internal/pb"
	"github.com/mseptian/qbit/internal/service"
	"github.com/mseptian/qbit/pkg/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandlerGrpc struct {
	pb.UnimplementedAuthServer
	auth service.AuthService
}

func NewAuthHandlerGrpcHandler(auth service.AuthService) *AuthHandlerGrpc {
	return &AuthHandlerGrpc{
		auth: auth,
	}
}

func (s *AuthHandlerGrpc) PostLogin(ctx context.Context, request *pb.MsgReqPostLogin) (*pb.MsgRespPostLogin, error) {
	validate := validator.NewValidator()
	if err := validate.Struct(request); err != nil {
		errorData, _ := json.Marshal(validator.ValidatorErrors(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", string(errorData))
	}
	createdUser, err := s.auth.Login(ctx, request)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return createdUser, nil
}

func (s *AuthHandlerGrpc) PostRegister(ctx context.Context, request *pb.MsgReqPostRegister) (*pb.DataEmpty, error) {
	validate := validator.NewValidator()
	if err := validate.Struct(request); err != nil {
		errorData, _ := json.Marshal(validator.ValidatorErrors(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", string(errorData))
	}
	createdUser, err := s.auth.Register(ctx, request)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return createdUser, nil
}
