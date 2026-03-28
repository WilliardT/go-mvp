package products_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_errors "github.com/WilliardT/go-mvp/internal/core/errors"
	core_postgres_pool "github.com/WilliardT/go-mvp/internal/core/repository/postgres/pool"
)

func (r *ProductsRepository) CreateProduct(
	ctx context.Context,
	product domain.Product,
) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO go_mvp_app.products (
			title,
			description,
			price,
			created_at,
			updated_at,
			author_user_id
		)
		VALUES ($1, $2, $3, NOW(), NOW(), $4)
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
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Product{}, fmt.Errorf(
				"%v: user with id='%d': %w",
				err,
				product.AuthorUserID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Product{}, fmt.Errorf("scan error: %w", err)
	}

	productDomain := domain.NewProduct(
		productModel.ID,
		productModel.Version,
		productModel.Title,
		productModel.Description,
		productModel.Price,
		productModel.CreatedAt,
		productModel.UpdatedAt,
		productModel.AuthorUserID,
	)

	return productDomain, nil
}
