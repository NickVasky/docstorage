package config

import (
	"fmt"
	"net"
	"os"
)

const (
	httpServerHostEnvName = "APP_HOST"
	httpServerPortEnvName = "APP_PORT"
)

type httpServerConfig struct {
	host string
	port string
}

type HttpServerConfig interface {
	Address() string
}

func (cfg *httpServerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func NewHttpServerConfig() (HttpServerConfig, error) {
	envs := map[string]string{}
	envs[httpServerHostEnvName] = os.Getenv(httpServerHostEnvName)
	envs[httpServerPortEnvName] = os.Getenv(httpServerPortEnvName)

	for k, v := range envs {
		if len(v) == 0 {
			return nil, fmt.Errorf("App env: %v is not set", k)
		}
	}

	return &httpServerConfig{
		host: envs[httpServerHostEnvName],
		port: envs[httpServerPortEnvName],
	}, nil
}
