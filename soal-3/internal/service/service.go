package service

import (
	"github.com/mseptian/qbit/internal/repository"
	"github.com/mseptian/qbit/pkg/auth"
	"github.com/mseptian/qbit/pkg/hash"
)

type Service struct {
	Auth    AuthService
	User    UserService
	Product ProductService
	Cart    CartService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    hash.Hashing
	Token      auth.JWTManager
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:    NewServiceAuth(*deps.Repository, deps.Hashing, deps.Token),
		User:    NewServiceUser(*deps.Repository, deps.Hashing, deps.Token),
		Product: NewServiceProduct(*deps.Repository, deps.Hashing, deps.Token),
		Cart:    NewServiceCart(*deps.Repository, deps.Hashing, deps.Token),
	}
}
