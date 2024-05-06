package service

import (
	"context"
	"errors"
	"github.com/mseptian/qbit/internal/pb"
	"github.com/mseptian/qbit/internal/repository"
	"github.com/mseptian/qbit/pkg/auth"
	"github.com/mseptian/qbit/pkg/hash"
)

type ServiceUser struct {
	Repository repository.Repositories
	Hash       hash.Hashing
	Token      auth.JWTManager
}

func NewServiceUser(auth repository.Repositories, hash hash.Hashing, token auth.JWTManager) *ServiceUser {
	return &ServiceUser{Repository: auth, Hash: hash, Token: token}
}

func (s *ServiceUser) Profile(ctx context.Context, token string) (*pb.MsgResGetUser, error) {
	jwt, _ := s.Token.Verify(token)

	user, RowsAffected := s.Repository.User.GetById(ctx, jwt.Audience)
	if RowsAffected < 1 {
		return nil, errors.New("user not found")
	}
	return &pb.MsgResGetUser{
		UserId: user.UserID,
		Email:  user.Email,
		Name:   user.Name,
	}, nil
}
