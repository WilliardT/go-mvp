package products_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_errors "github.com/WilliardT/go-mvp/internal/core/errors"
	core_postgres_pool "github.com/WilliardT/go-mvp/internal/core/repository/postgres/pool"
)

func (r *ProductsRepository) PatchProduct(
	ctx context.Context,
	id int,
	product domain.Product,
) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE go_mvp_app.products
		SET
			title = $1,
			description = $2,
			price = $3,
			updated_at = NOW(),
			author_user_id = $4,
			version = version + 1
		WHERE id = $5 AND version = $6
		RETURNING
			id,
			version,
			title,
			description,
			price::float8,
			created_at,
			updated_at,
			author_user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		product.Title,
		product.Description,
		product.Price,
		product.AuthorUserID,
		id,
		product.Version,
	)

	var productModel ProductModel

	err := row.Scan(
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
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Product{}, fmt.Errorf(
				"product with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.Product{}, fmt.Errorf("scan error: %w", err)
	}

	return productDomainFromModel(productModel), nil
}
