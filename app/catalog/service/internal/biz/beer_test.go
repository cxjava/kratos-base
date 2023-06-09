package biz_test

import (
	"kratos-base/app/catalog/service/internal/biz"
	"kratos-base/app/catalog/service/internal/mocks/mrepo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("CatalogUsecase", func() {
	var userCase biz.CatalogUseCase
	var mUserRepo *mrepo.MockBeerRepo

	BeforeEach(func() {
		mUserRepo = mrepo.NewMockBeerRepo(ctl)
		userCase = biz.NewBeerUseCase(mUserRepo, nil)
	})

	It("Create", func() {
		info := &biz.Beer{
			Name: "admin",
		}
		infoResp := &biz.Beer{
			Id:   11,
			Name: "admin",
		}
		mUserRepo.EXPECT().CreateBeer(ctx, gomock.Any()).Return(infoResp, nil)
		l, err := userCase.Create(ctx, info)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l.Id).To(Equal(int64(11)))
		Ω(l.Name).To(Equal("admin"))
	})

})
