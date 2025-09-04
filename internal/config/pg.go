package config

import (
	"fmt"
	"os"
)

const (
	pgHostEnvName     = "PG_HOST"
	pgPortEnvName     = "PG_PORT"
	pgUserEnvName     = "PG_USER"
	pgPasswordEnvName = "PG_PASSWORD"
	pgDBEnvName       = "PG_DB"
	pgSSLModeEnvName  = "PG_SSLMODE"
)

type pgConfig struct {
	dsn string
}

type PgConfig interface {
	DSN() string
}

func NewPgConfig() (PgConfig, error) {
	envs := map[string]string{}
	envs[pgHostEnvName] = os.Getenv(pgHostEnvName)
	envs[pgPortEnvName] = os.Getenv(pgPortEnvName)
	envs[pgUserEnvName] = os.Getenv(pgUserEnvName)
	envs[pgPasswordEnvName] = os.Getenv(pgPasswordEnvName)
	envs[pgDBEnvName] = os.Getenv(pgDBEnvName)
	envs[pgSSLModeEnvName] = os.Getenv(pgSSLModeEnvName)

	for k, v := range envs {
		if len(v) == 0 {
			return nil, fmt.Errorf("PG env: %v is not set", k)
		}
	}

	DSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envs[pgHostEnvName], envs[pgPortEnvName], envs[pgUserEnvName], envs[pgPasswordEnvName], envs[pgDBEnvName], envs[pgSSLModeEnvName])

	return &pgConfig{
		dsn: DSN,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
