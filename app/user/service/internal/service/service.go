package service

import (
	"context"
	v1 "kratos-base/api/user/service/v1"
	"kratos-base/app/user/service/internal/data/ent"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

//go:generate mockgen -source=service.go -destination=../mocks/mbiz/user_biz.go -package=mbiz
type UserUseCase interface {
	Create(ctx context.Context, req *v1.UserCreateReq) (*ent.User, error)
	Update(ctx context.Context, req *v1.UserUpdateReq) (*ent.User, error)
	Get(ctx context.Context, req *v1.UserReq) (*ent.User, error)
}
