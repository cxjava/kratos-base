package service_test

import (
	v1 "kratos-base/api/user/service/v1"
	"kratos-base/app/user/service/internal/data/ent"
	"kratos-base/app/user/service/internal/mocks/mbiz"
	"kratos-base/app/user/service/internal/service"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserService", func() {
	var userService *service.UserService
	var mUserBiz *mbiz.MockUserUseCase

	BeforeEach(func() {
		mUserBiz = mbiz.NewMockUserUseCase(ctl)
		userService = service.NewUserService(mUserBiz, nil)
	})

	It("Create User", func() {
		info := &ent.User{
			ID:   1,
			Name: "admin",
		}
		infoReq := &v1.UserCreateReq{
			Id:   1,
			Name: "admin",
		}
		mUserBiz.EXPECT().Create(gomock.Any(), gomock.Any()).Return(info, nil)
		l, err := userService.CreateUser(ctx, infoReq)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l.Result.Id).To(Equal(int64(1)))
		Ω(l.Result.Name).To(Equal("admin"))
	})

})
