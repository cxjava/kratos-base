package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-base/api/user/v1"
	"kratos-base/internal/biz"
	"kratos-base/internal/data/ent"
)

type userRepo struct {
	commonRepo *CommonRepo
	data       *Data
	log        *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, commonRepo *CommonRepo, logger log.Logger) biz.UserRepo {
	return &userRepo{
		commonRepo: commonRepo,
		data:       data,
		log:        log.NewHelper(log.With(logger, "module", "base-service/data")),
	}
}

// CreateUser 创建
func (r *userRepo) CreateUser(ctx context.Context, req *v1.UserCreateReq) (*ent.User, error) {
	return r.data.db.User.Create().
		SetAge(req.Age).
		SetName(req.Name).
		Save(ctx)

}

// UpdateUser 更新
func (r *userRepo) UpdateUser(ctx context.Context, req *v1.UserUpdateReq) (*ent.User, error) {
	return r.data.db.User.UpdateOneID(req.Id).
		SetAge(req.Age).
		SetName(req.Name).
		Save(ctx)
}

// GetUser 根据Id查询
func (r *userRepo) GetUser(ctx context.Context, req *v1.UserReq) (*ent.User, error) {
	return r.data.db.User.Get(ctx, req.Id)
}

