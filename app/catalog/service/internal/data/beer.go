package data

import (
	"context"
	"kratos-base/app/catalog/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.BeerRepo = (*beerRepo)(nil)

type beerRepo struct {
	data *Data
	log  *log.Helper
}

func NewBeerRepo(data *Data, logger log.Logger) biz.BeerRepo {
	return &beerRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/beer")),
	}
}

func (r *beerRepo) CreateBeer(ctx context.Context, b *biz.Beer) (*biz.Beer, error) {
	po, err := r.data.db.Beer.
		Create().
		SetName(b.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Beer{
		Id:   po.ID,
		Name: po.Name,
	}, nil
}

func (r *beerRepo) GetBeer(ctx context.Context, id int64) (*biz.Beer, error) {
	po, err := r.data.db.Beer.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Beer{
		Id:   po.ID,
		Name: po.Name,
	}, nil
}
