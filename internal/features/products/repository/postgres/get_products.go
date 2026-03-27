package products_postgres_repository

import (
	"context"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

func (r *ProductsRepository) GetProducts(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT
			id,
			version,
			title,
			description,
			price::float8,
			created_at,
			updated_at,
			user_id
		FROM go_mvp_app.products
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select products: %w", err)
	}

	defer rows.Close()

	var productModels []ProductModel

	for rows.Next() {
		var productModel ProductModel

		err := rows.Scan(
			&productModel.ID,
			&productModel.Version,
			&productModel.Title,
			&productModel.Description,
			&productModel.Price,
			&productModel.CreatedAt,
			&productModel.UpdatedAt,
			&productModel.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan products: %w", err)
		}

		productModels = append(productModels, productModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return productDomainsFromModels(productModels), nil
}
