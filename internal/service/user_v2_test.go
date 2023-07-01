package service_test

import (
	"context"
	"errors"
	pb "kratos-base/api/user/v1"
	v1 "kratos-base/api/user/v1"
	"kratos-base/internal/convert"
	"kratos-base/internal/data/ent"
	"kratos-base/internal/mocks/mbiz"
	"kratos-base/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var (
	errUserNotFound = errors.New("user not found in database")
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func TestUserService_CreateUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mUserBiz := mbiz.NewMockUserUseCase(ctrl)
	userService := service.NewUserService(mUserBiz, nil)

	infoReq := &v1.UserCreateReq{
		Id:   1,
		Name: "admin",
	}

	tests := []test{
		{
			name: "user create successfully",
			mock: func() {
				info := &ent.User{
					ID:   1,
					Name: "admin",
				}
				mUserBiz.EXPECT().Create(gomock.Any(), gomock.Any()).Return(info, nil)
			},
			res: &pb.UserCreateReply{
				Code:    200,
				Message: "success",
				Result: convert.UserData2Reply(&ent.User{
					ID:   1,
					Name: "admin",
				}),
			},
			err: nil,
		},
		{
			name: "user not found",
			mock: func() {
				mUserBiz.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errUserNotFound)
			},
			res: nil,
			err: errUserNotFound,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			res, err := userService.CreateUser(context.Background(), infoReq)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			}
			if tc.res != nil {
				require.EqualValues(t, res, tc.res)
			}
		})
	}
}
