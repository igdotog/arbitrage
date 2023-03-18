package bundle

import (
	"context"

	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/internal/dbo"
)

func (s *service) Save(ctx context.Context, bundle *domain.Bundle) error {
	return s.bundle.Save(ctx, &dbo.Bundle{
		Symbol:               bundle.Symbol,
		ExchangeFrom:         bundle.ExchangeFrom,
		ExchangeTo:           bundle.ExchangeTo,
		PercentageDifference: bundle.PercentageDifference,
	})
}