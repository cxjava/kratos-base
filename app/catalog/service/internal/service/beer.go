package service

import (
	"context"

	v1 "kratos-base/api/catalog/service/v1"
	"kratos-base/app/catalog/service/internal/biz"
)

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
