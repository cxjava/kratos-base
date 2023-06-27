package biz

import (
	"context"

	"golang.org/x/sync/singleflight"
	"kratos-base/pkg/page_token"
	"github.com/go-kratos/kratos/v2/log"
)

type Image struct {
	URL string
}

type Beer struct {
	Id          int64
	Name        string
}

//go:generate mockgen -destination=../mocks/mrepo/catalog.go -package=mrepo . BeerRepo
type BeerRepo interface {
	CreateBeer(ctx context.Context, c *Beer) (*Beer, error)
	GetBeer(ctx context.Context, id int64) (*Beer, error)
}

type BeerUseCase struct {
	repo      BeerRepo
	log       *log.Helper
	pageToken page_token.ProcessPageTokens
	sg        singleflight.Group
}

func NewBeerUseCase(repo BeerRepo, logger log.Logger) CatalogUseCase {
	return &BeerUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/beer")), pageToken: page_token.NewTokenGenerate()}
}

func (uc *BeerUseCase) Create(ctx context.Context, u *Beer) (*Beer, error) {
	return uc.repo.CreateBeer(ctx, u)
}

func (uc *BeerUseCase) Get(ctx context.Context, id int64) (*Beer, error) {
	return uc.repo.GetBeer(ctx, id)
}
