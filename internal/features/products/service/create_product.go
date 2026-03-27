package products_service

import (
	"context"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

func (s *ProductsService) CreateProduct(
	ctx context.Context,
	product domain.Product,
) (domain.Product, error) {
	if err := product.Validate(); err != nil {
		return domain.Product{}, fmt.Errorf("validate product domain: %w", err)
	}

	product, err := s.productsRepository.CreateProduct(ctx, product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("create product: %w", err)
	}

	return product, nil
}
