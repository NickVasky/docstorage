package app

import (
	"context"
	"log"

	"github.com/NickVasky/docstorage/internal/api/service"
	"github.com/NickVasky/docstorage/internal/closer"
	"github.com/NickVasky/docstorage/internal/config"
	"github.com/NickVasky/docstorage/internal/repository"
	"github.com/NickVasky/docstorage/internal/repository/documents"
	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	pgConfig         config.PgConfig
	httpServerConfig config.HttpServerConfig

	documentsRepo repository.DocumentsRepo
	pgConn        *pgxpool.Pool

	apiService service.ServiceInterface
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PgConfig() config.PgConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPgConfig()
		if err != nil {
			log.Fatalf("Failed to get PG config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) HttpConfig() config.HttpServerConfig {
	if s.httpServerConfig == nil {
		cfg, err := config.NewHttpServerConfig()
		if err != nil {
			log.Fatalf("Failed to get http server config: %s", err.Error())
		}

		s.httpServerConfig = cfg
	}

	return s.httpServerConfig
}

func (s *serviceProvider) DocumentsRepo(ctx context.Context) repository.DocumentsRepo {
	if s.documentsRepo == nil {
		s.documentsRepo = documents.NewRepo(s.PgConn(ctx))
	}
	return s.documentsRepo
}

func (s *serviceProvider) PgConn(ctx context.Context) *pgxpool.Pool {
	if s.pgConn == nil {
		pgx, err := pgxpool.New(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to postgres: %s", err.Error())
		}

		err = pgx.Ping(ctx)
		if err != nil {
			log.Fatalf("postgres ping error: %s", err.Error())
		}
		closer.Add(
			func() error {
				pgx.Close()
				return nil
			})

		s.pgConn = pgx
	}
	return s.pgConn
}

func (s *serviceProvider) APIService(ctx context.Context) service.ServiceInterface {
	if s.apiService == nil {
		s.apiService = service.NewServiceImpl(s.DocumentsRepo(ctx))
	}
	return s.apiService
}
