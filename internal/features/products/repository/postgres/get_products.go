package products_postgres_repository

import (
	"context"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

func (r *ProductsRepository) GetProducts(
	ctx context.Context,
	authorUserID *int,
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
			author_user_id
		FROM go_mvp_app.products
		WHERE ($1::int IS NULL OR author_user_id = $1)
		ORDER BY id ASC
		LIMIT $2
		OFFSET $3;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		authorUserID,
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
			&productModel.AuthorUserID,
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
