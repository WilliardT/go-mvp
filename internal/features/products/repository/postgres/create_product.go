package products_postgres_repository

import (
	"context"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
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
			user_id
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
			user_id;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		product.Title,
		product.Description,
		product.Price,
		product.UserID,
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
		&productModel.UserID,
	)
	if err != nil {
		return domain.Product{}, fmt.Errorf("scan error: %w", err)
	}

	return productDomainFromModel(productModel), nil
}
