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

type ServiceAuth struct {
	Repository repository.Repositories
	Hash       hash.Hashing
	Token      auth.JWTManager
}

func NewServiceAuth(auth repository.Repositories, hash hash.Hashing, token auth.JWTManager) *ServiceAuth {
	return &ServiceAuth{Repository: auth, Hash: hash, Token: token}
}

func (s *ServiceAuth) Login(ctx context.Context, input *pb.MsgReqPostLogin) (*pb.MsgRespPostLogin, error) {

	user, RowsAffected := s.Repository.User.GetByEmail(ctx, input.Email)
	if RowsAffected < 1 {
		return nil, errors.New("user not found")
	}

	if s.Hash.ComparePassword(user.Password, input.Password) != nil {
		return nil, errors.New("wrong password")
	}

	token, _ := s.Token.Generate(user)

	return &pb.MsgRespPostLogin{
		Error: false,
		Msg:   "",
		Data:  &pb.MsgRespPostLogin_DataPostLogin{Token: token},
	}, nil
}

func (s *ServiceAuth) Register(ctx context.Context, input *pb.MsgReqPostRegister) (*pb.DataEmpty, error) {

	_, RowsAffected := s.Repository.User.GetByEmail(ctx, input.Email)
	if RowsAffected > 0 {
		return nil, errors.New("user already register")
	}

	user := &models.User{
		Email:    input.Email,
		Password: s.Hash.HashPassword(input.Password),
		Name:     input.Name,
	}
	err := s.Repository.User.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.DataEmpty{
		Error: false,
		Msg:   "Register Success",
		Data:  nil,
	}, nil
}
