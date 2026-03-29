package statistics_postgres_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *StatisticsRepository) GetProductsPrices(
	ctx context.Context,
	authorUserID *int,
	createdFrom *time.Time,
	createdTo *time.Time,
) ([]float64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT price::float8
		FROM go_mvp_app.products
		WHERE ($1::int IS NULL OR author_user_id = $1)
			AND ($2::timestamptz IS NULL OR created_at >= $2)
			AND ($3::timestamptz IS NULL OR created_at < ($3::timestamptz + INTERVAL '1 day'))
		ORDER BY id ASC;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		authorUserID,
		createdFrom,
		createdTo,
	)
	if err != nil {
		return nil, fmt.Errorf("select product prices: %w", err)
	}
	defer rows.Close()

	prices := make([]float64, 0)

	for rows.Next() {
		var price float64

		if err := rows.Scan(&price); err != nil {
			return nil, fmt.Errorf("scan product price: %w", err)
		}

		prices = append(prices, price)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next product price rows: %w", err)
	}

	return prices, nil
}
