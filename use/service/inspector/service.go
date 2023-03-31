package inspector

import (
	"context"
	"strings"
	"time"

	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/binance"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/bybit"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/huobi"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/mexc"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

type Service interface {
	Inspect(ctx context.Context)
	Bundles() chan domain.Bundle
	Symbols() ([]string, error)
}

type (
	service struct {
		log *log.Logger

		exchangers      []exchangers.Client
		primaryExchange exchangers.Client

		bundles chan domain.Bundle
	}

	config struct {
		PrimaryName string `config:"PRIMARY_EXCHANGE_NAME"`
	}
)

func New(log *log.Logger, exchangers []exchangers.Client) Service {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cfg := &config{}
	loader := confita.NewLoader(env.NewBackend())
	if err := loader.Load(ctx, cfg); err != nil {
		return nil
	}

	primaryExchange := exchangers[0]
	for _, e := range exchangers {
		if strings.EqualFold(e.Name(), cfg.PrimaryName) {
			primaryExchange = e
		}
	}

	return &service{log: log, exchangers: exchangers, primaryExchange: primaryExchange, bundles: make(chan domain.Bundle, 1000)}
}
func (s *service) Bundles() chan domain.Bundle {
	return s.bundles
}

func DefaultExchangers(log *log.Logger) []exchangers.Client {
	return []exchangers.Client{binance.New(log), bybit.New(log), mexc.New(log), huobi.New(log)}
}
