package biz

import (
	"context"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBeerUseCase)

//go:generate mockgen -source=biz.go -destination=../mocks/mbiz/catalog_biz.go -package=mbiz
type CatalogUseCase interface {
	Get(ctx context.Context, id int64) (*Beer, error) 
	Create(ctx context.Context, u *Beer) (*Beer, error)
}