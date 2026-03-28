package products_service

import (
	"context"
	"fmt"
)

func (s *ProductsService) DeleteProduct(
	ctx context.Context,
	id int,
) error {
	if err := s.productsRepository.DeleteProduct(ctx, id); err != nil {
		return fmt.Errorf("delete product: %w", err)
	}

	return nil
}
