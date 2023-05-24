package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"kratos-base/internal/convert"

	pb "kratos-base/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  UserUseCase
	log *log.Helper
}

func NewUserService(uc UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(log.With(logger, "module", "base-service/service"))}
}
func (s *UserService) GetUser(ctx context.Context, req *pb.UserReq) (*pb.UserReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUser")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserData2Reply(data),
	}
	return reply, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UserUpdateReq) (*pb.UserUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateUser")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserData2Reply(data),
	}
	return reply, nil
}
func (s *UserService) CreateUser(ctx context.Context, req *pb.UserCreateReq) (*pb.UserCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateUser")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserData2Reply(data),
	}
	return reply, err
}
