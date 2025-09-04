package app

import (
	"log"

	"github.com/NickVasky/docstorage/internal/config"
)

type serviceProvider struct {
	pgConfig         config.PgConfig
	httpServerConfig config.HttpServerConfig
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
