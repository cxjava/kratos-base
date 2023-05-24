package biz

import (
	"context"
	"kratos-base/api/user/service/v1"
	"kratos-base/app/user/service/internal/data/ent"
	"kratos-base/app/user/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
)

//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
type UserRepo interface {
	CreateUser(context.Context, *v1.UserCreateReq) (*ent.User, error)
	UpdateUser(context.Context, *v1.UserUpdateReq) (*ent.User, error)
	GetUser(context.Context, *v1.UserReq) (*ent.User, error)
}

type userUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) service.UserUseCase {
	return &userUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "base-service/biz")),
	}
}

func (uc *userUseCase) Create(ctx context.Context, req *v1.UserCreateReq) (*ent.User, error) {
	return uc.repo.CreateUser(ctx, req)
}
func (uc *userUseCase) Update(ctx context.Context, req *v1.UserUpdateReq) (*ent.User, error) {
	return uc.repo.UpdateUser(ctx, req)
}
func (uc *userUseCase) Get(ctx context.Context, req *v1.UserReq) (*ent.User, error) {
	return uc.repo.GetUser(ctx, req)
}
