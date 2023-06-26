package biz_test

import (
	v1 "kratos-base/api/user/service/v1"
	"kratos-base/app/user/service/internal/biz"
	"kratos-base/app/user/service/internal/data/ent"
	"kratos-base/app/user/service/internal/mocks/mrepo"
	"kratos-base/app/user/service/internal/service"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserUsecase", func() {
	var userCase service.UserUseCase
	var mUserRepo *mrepo.MockUserRepo

	BeforeEach(func() {
		mUserRepo = mrepo.NewMockUserRepo(ctl)
		userCase = biz.NewUserUseCase(mUserRepo, nil)
	})

	It("Create", func() {
		info := &ent.User{
			ID:   1,
			Name: "admin",
		}
		infoReq := &v1.UserCreateReq{
			Id:   1,
			Name: "admin",
		}
		mUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(info, nil)
		l, err := userCase.Create(ctx, infoReq)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l.ID).To(Equal(int64(1)))
		Ω(l.Name).To(Equal("admin"))
	})

})
