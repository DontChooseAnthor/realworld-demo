package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"realworld/internal/biz"
	"realworld/internal/data/ent/user"
)

type realWorldRepo struct {
	user userRepo
}

func NewRealWorldRepo(data *Data, logger log.Logger) {
	NewUserRepo(data, logger)
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	_, err := ur.data.client.User.
		Create().
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetUsername(user.Username).
		SetBio(user.Bio).
		SetImage(user.Image).
		Save(ctx)
	return err
}

func (ur *userRepo) GetUser(ctx context.Context, email string) (*biz.User, error) {
	u, err := ur.data.client.User.Query().Where(user.Email(email)).First(ctx)
	if err != nil {
		ur.log.Error("query user error:", err)
		return nil, errors.New("query user error")
	}

	return &biz.User{
		// 转token
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Bio:      u.Bio,
		Image:    u.Image,
	}, nil
}
