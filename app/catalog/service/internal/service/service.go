package service

import (
	"context"
	v1 "kratos-base/api/catalog/service/v1"
	"kratos-base/app/catalog/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCatalogService)

type CatalogService struct {
	v1.UnimplementedCatalogServer

	bc CatalogUseCase
	log *log.Helper
}

func NewCatalogService(bc CatalogUseCase, logger log.Logger) *CatalogService {
	return &CatalogService{
		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}

//go:generate mockgen -source=service.go -destination=../mocks/mbiz/catalog_biz.go -package=mbiz
type CatalogUseCase interface {
	Get(ctx context.Context, id int64) (*biz.Beer, error) 
	Create(ctx context.Context, u *biz.Beer) (*biz.Beer, error)
}