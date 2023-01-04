package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "realworld/api/user/service/v1"
	"realworld/internal/biz"
	"realworld/internal/middleware"
)

var (
	ParamsValidErr = "params valid error"
	PwdValidErr    = "password valid failed"
)

type UserService struct {
	pb.UnimplementedRealWorldServer
	log  *log.Helper
	user *biz.UserUseCase
}

func NewUserService(user *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		user: user,
		log:  log.NewHelper(logger),
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	s.log.Infof("input data %v", req)
	user, err := s.user.Get(ctx, req.User.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != req.User.Password {
		return nil, errors.New(PwdValidErr)
	}
	// 生成token,存redis
	token, err := middleware.JWTGenerate(req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		User: &pb.LoginReply_User{
			Email:    user.Email,
			Username: user.Username,
			Bio:      user.Bio,
			Image:    user.Image,
		},
		Token: token,
	}, nil
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	s.log.Infof("input data %v", req)
	if req.User.Email == "" || req.User.Password == "" || req.User.Username == "" {
		return nil, errors.New(ParamsValidErr)
	}
	err := s.user.Create(ctx, &biz.User{
		Email:    req.User.Email,
		Username: req.User.Password,
		Password: req.User.Username,
	})

	if err != nil {
		return nil, err
	}

	// 生成token,存redis
	token, err := middleware.JWTGenerate(req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterReply{
		User: &pb.RegisterReply_User{
			Email:    req.User.Email,
			Username: req.User.Username,
			Bio:      "",
			Image:    "",
		},
		Token: token,
	}, nil
}
