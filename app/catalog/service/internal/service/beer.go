package service

import (
	"context"

	v1 "kratos-base/api/catalog/service/v1"
	"kratos-base/app/catalog/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type CatalogService struct {
	v1.UnimplementedCatalogServer

	bc  biz.CatalogUseCase
	log *log.Helper
}

func NewCatalogService(bc biz.CatalogUseCase, logger log.Logger) *CatalogService {
	return &CatalogService{
		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}

func (s *CatalogService) CreateBeer(ctx context.Context, req *v1.CreateBeerReq) (*v1.CreateBeerReply, error) {
	b := &biz.Beer{
		Name:        req.Name,
	}
	x, err := s.bc.Create(ctx, b)
	return &v1.CreateBeerReply{
		Id:          x.Id,
		Name:        x.Name,
	}, err
}

func (s *CatalogService) GetBeer(ctx context.Context, req *v1.GetBeerReq) (*v1.GetBeerReply, error) {
	x, err := s.bc.Get(ctx, req.Id)
	return &v1.GetBeerReply{
		Id:          x.Id,
		Name:        x.Name,
	}, err
}
