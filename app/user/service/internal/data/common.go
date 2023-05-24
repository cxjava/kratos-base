package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-base/app/user/service/internal/data/ent"
)

type CommonRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommonRepo(data *Data, logger log.Logger) *CommonRepo {
	return &CommonRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "ent/transaction-data")),
	}
}

func (r *CommonRepo) Transaction(ctx context.Context, f func(tx *ent.Tx) error) error {
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		r.log.Error("Transaction failed:", err)
		return err
	}
	err = f(tx)
	if err != nil {
		e := tx.Rollback()
		if e != nil {
			r.log.Error("Rollback failed:", err)
			return e
		}
		return err
	}
	return tx.Commit()
}
