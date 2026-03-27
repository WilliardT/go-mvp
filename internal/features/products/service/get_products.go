package products_service

import (
	"context"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_errors "github.com/WilliardT/go-mvp/internal/core/errors"
)

func (s *ProductsService) GetProducts(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Product, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	products, err := s.productsRepository.GetProducts(
		ctx,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("get products from repository: %w", err)
	}

	return products, nil
}
