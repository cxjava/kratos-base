package data

import (
	"context"
	"fmt"
	catalogv1 "kratos-base/api/catalog/service/v1"
	"kratos-base/app/user/service/internal/conf"
	"kratos-base/app/user/service/internal/data/ent"
	"kratos-base/app/user/service/internal/data/ent/migrate"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewUserRepo, NewCommonRepo,NewCatalogServiceClient)

// generate the mocked CatalogClient
//go:generate mockgen -destination=../mocks/mcatalog/catalog_data.go -package=mcatalog kratos-base/api/catalog/service/v1 CatalogClient

// Data .
type Data struct {
	db  *ent.Client
	log *log.Helper
	bc  catalogv1.CatalogClient // use mockgen when run the unit test
}

func NewEntClient(conf *conf.Data, logger log.Logger) (*ent.Client, error) {
	log := log.NewHelper(log.With(logger, "module", "ent/data"))
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(25 * time.Second)

	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		log.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		log.Errorf("failed opening connection to postgresql: %v", err)
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, err
	}
	return client, nil
}

// NewData .
func NewData(conf *conf.Data, db *ent.Client, logger log.Logger, bc catalogv1.CatalogClient) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "base-service/data"))

	d := &Data{
		db:  db,
		log: log,
		bc:  bc,
	}
	cleanup := func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}

func NewCatalogServiceClient(conf *conf.Data) catalogv1.CatalogClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.GetCatalogAddress()),
		// grpc.WithEndpoint("discovery:///beer.catalog.service"),
		// grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return catalogv1.NewCatalogClient(conn)
}