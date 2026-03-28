package products_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/WilliardT/go-mvp/internal/core/errors"
)

func (r *ProductsRepository) DeleteProduct(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM go_mvp_app.products
		WHERE id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"product with id='%d': %w",
			id,
			core_errors.ErrNotFound,
		)
	}

	return nil
}
