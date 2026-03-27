package products_service

import (
	"context"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

func (s *ProductsService) GetProduct(
	ctx context.Context,
	id int,
) (domain.Product, error) {
	product, err := s.productsRepository.GetProduct(ctx, id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("get product from repository: %w", err)
	}

	return product, nil
}
