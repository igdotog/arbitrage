package mexc

import (
	"context"
	"net/http"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/igdotog/core/logger"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

const ExchangeName = "mexc"

type (
	client struct {
		httpClient *http.Client
		hosts      []string

		cfg    *config
		logger *logger.Logger
	}

	config struct {
		ApiKey    string `config:"MEXC_API_KEY"`
		SecretKey string `config:"MEXC_SECRET_KEY"`
	}
)

func New(logger *logger.Logger) exchangers.Client {
	cfg := &config{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := confita.NewLoader(env.NewBackend()).Load(ctx, cfg); err != nil {
		logger.Error(err)
	}

	httpClient := &http.Client{
		Timeout: exchangers.ProvTimeoutSec * time.Second,
	}

	hosts := []string{
		"https://api.mexc.com",
	}

	return &client{httpClient: httpClient, cfg: cfg, hosts: hosts, logger: logger}
}

func (c *client) Name() string {
	return ExchangeName
}
