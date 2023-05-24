package service_test

import (
	v1 "kratos-base/api/catalog/service/v1"
	"kratos-base/app/catalog/service/internal/biz"
	"kratos-base/app/catalog/service/internal/mocks/mbiz"
	"kratos-base/app/catalog/service/internal/service"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CatalogService", func() {
	var catalogService *service.CatalogService
	var mCatalogBiz *mbiz.MockCatalogUseCase

	BeforeEach(func() {
		mCatalogBiz = mbiz.NewMockCatalogUseCase(ctl)
		catalogService = service.NewCatalogService(mCatalogBiz, nil)
	})

	It("Create Catalog", func() {
		info := &biz.Beer{
			Id:   1,
			Name: "admin",
		}
		infoReq := &v1.CreateBeerReq{
			Name: "admin",
		}
		mCatalogBiz.EXPECT().Create(gomock.Any(), gomock.Any()).Return(info, nil)
		l, err := catalogService.CreateBeer(ctx, infoReq)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l.Id).To(Equal(int64(1)))
		Ω(l.Name).To(Equal("admin"))
	})

})
