package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUsernameInvalid = errors.New("username invalid")
	ErrEmailInvalid    = errors.New("email invalid")
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUserNotFound    = errors.New("user not found")
)

type User struct {
	Email    string
	Username string
	Password string
	Bio      string
	Image    string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, email string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) Get(ctx context.Context, email string) (*User, error) {
	return uc.repo.GetUser(ctx, email)
}
