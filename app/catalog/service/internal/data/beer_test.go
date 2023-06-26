package data_test

import (
	"kratos-base/app/catalog/service/internal/biz"
	"kratos-base/app/catalog/service/internal/data"

	"github.com/go-kratos/kratos/v2/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Beer", func() {
	var ro biz.BeerRepo
	var uD *biz.Beer
	BeforeEach(func() {
		ro = data.NewBeerRepo(Db, log.DefaultLogger)
		// 这里你可以不引入外部组装好的数据，可以在这里直接写
		uD = &biz.Beer{
			Id: 11,
			Name: "beer",
		}
	})
	// 设置 It 块来添加单个规格
	It("CreateBeer", func() {
		u, err := ro.CreateBeer(ctx, uD)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(u.Id).Should(Equal(int64(1)))
		Ω(u.Name).Should(Equal("beer"))
	})
	// 设置 It 块来添加单个规格
	It("ListBeer", func() {
		user, err := ro.GetBeer(ctx, 1)
		Ω(err).ShouldNot(HaveOccurred())   // 获取列表不应该出现错误
		Ω(user).ShouldNot(BeNil())         // 结果不应该为空
		Ω(user.Id).Should(Equal(int64(1))) // 总数应该为 1，因为上面只创建了一条
		Ω(user.Name).Should(Equal("beer"))
	})
})
