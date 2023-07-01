package data_test

import (
	v1 "kratos-base/api/user/service/v1"
	"kratos-base/app/user/service/internal/biz"
	"kratos-base/app/user/service/internal/data"
	"kratos-base/app/user/service/internal/testdata"

	"github.com/go-kratos/kratos/v2/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("User", func() {
	var ro biz.UserRepo
	var uD *v1.UserCreateReq
	BeforeEach(func() {
		ro = data.NewUserRepo(Db, nil, log.DefaultLogger)
		// 这里你可以不引入外部组装好的数据，可以在这里直接写
		uD = testdata.User()
	})
	// 设置 It 块来添加单个规格
	It("CreateUser", func() {
		mCatalog.EXPECT().CreateBeer(gomock.Any(), gomock.Any()).Return(testdata.Catalog(), nil)
		u, err := ro.CreateUser(ctx, uD)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(u.Age).Should(Equal(int32(11)))
		// just ignore it, but in the real case, need do more validation base on the CreateBeer interface return value
		Ω(u.Name).Should(Equal("admin_catalog_test"))
	})
	// 设置 It 块来添加单个规格
	It("ListUser", func() {
		req := &v1.UserReq{Id: 1}
		user, err := ro.GetUser(ctx, req)
		Ω(err).ShouldNot(HaveOccurred())   // 获取列表不应该出现错误
		Ω(user).ShouldNot(BeNil())         // 结果不应该为空
		Ω(user.ID).Should(Equal(int64(1))) // 总数应该为 1，因为上面只创建了一条
		Ω(user.Age).Should(Equal(int32(11)))
		Ω(user.Name).Should(Equal("admin_catalog_test"))
	})

	// 设置 It 块来添加单个规格
	It("UpdateUser", func() {
		req := &v1.UserUpdateReq{Id: 1, Name: "updated", Age: 32}
		user, err := ro.UpdateUser(ctx, req)
		Ω(err).ShouldNot(HaveOccurred()) // 更新不应该出现错误
		Ω(user).ShouldNot(BeNil())       // 结果不应该为空
		Ω(user.Age).Should(Equal(int32(32)))
		Ω(user.Name).Should(Equal("updated"))
	})

})
