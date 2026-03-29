package statistics_postgres_repository

import (
	"github.com/WilliardT/go-mvp/internal/core/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

type StatisticsModel struct {
	ProductsCount      int
	ProductPriceMedian pgtype.Float8
}

func statisticsDomainFromModel(model StatisticsModel) domain.Statistics {
	var productPriceMedian *float64

	if model.ProductPriceMedian.Valid {
		productPriceMedian = &model.ProductPriceMedian.Float64
	}

	return domain.NewStatistics(
		model.ProductsCount,
		productPriceMedian,
		domain.ProductPriceRanges{},
	)
}
