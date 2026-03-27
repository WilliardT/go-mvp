package products_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_errors "github.com/WilliardT/go-mvp/internal/core/errors"
	core_postgres_pool "github.com/WilliardT/go-mvp/internal/core/repository/postgres/pool"
)

func (r *ProductsRepository) GetProduct(
	ctx context.Context,
	id int,
) (domain.Product, error) {
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
		WHERE id = $1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var productModel ProductModel

	err := row.Scan(
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
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Product{}, fmt.Errorf(
				"product with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.Product{}, fmt.Errorf("scan product error: %w", err)
	}

	return productDomainFromModel(productModel), nil
}
