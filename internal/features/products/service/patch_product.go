package products_service

import (
	"context"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

func (s *ProductsService) PatchProduct(
	ctx context.Context,
	id int,
	patch domain.ProductPatch,
) (domain.Product, error) {
	product, err := s.productsRepository.GetProduct(ctx, id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("get product: %w", err)
	}

	if err := product.ApplyPatch(patch); err != nil {
		return domain.Product{}, fmt.Errorf("apply product patch: %w", err)
	}

	patchedProduct, err := s.productsRepository.PatchProduct(ctx, id, product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("patch product: %w", err)
	}

	return patchedProduct, nil
}
