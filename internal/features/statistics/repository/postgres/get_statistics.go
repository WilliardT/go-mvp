package statistics_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

func (r *StatisticsRepository) GetProductsStatistics(
	ctx context.Context,
	authorUserID *int,
	createdFrom *time.Time,
	createdTo *time.Time,
) (domain.Statistics, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT
			COUNT(*)::int AS products_count,
			percentile_cont(0.5) WITHIN GROUP (ORDER BY price::float8) AS product_price_median
		FROM go_mvp_app.products
		WHERE ($1::int IS NULL OR author_user_id = $1)
			AND ($2::timestamptz IS NULL OR created_at >= $2)
			AND ($3::timestamptz IS NULL OR created_at < ($3::timestamptz + INTERVAL '1 day'));
	`

	// можно сделать через append к строке

	var statisticsModel StatisticsModel

	err := r.pool.QueryRow(
		ctx,
		query,
		authorUserID,
		createdFrom,
		createdTo,
	).Scan(
		&statisticsModel.ProductsCount,
		&statisticsModel.ProductPriceMedian,
	)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			"select products statistics: %w",
			err,
		)
	}

	return statisticsDomainFromModel(statisticsModel), nil
}
