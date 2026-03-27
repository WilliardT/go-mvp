package products_service

import (
	"context"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

type ProductsService struct {
	productsRepository ProductsRepository
}

type ProductsRepository interface {
	CreateProduct(
		ctx context.Context,
		product domain.Product,
	) (domain.Product, error)

	GetProducts(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Product, error)

	GetProduct(
		ctx context.Context,
		id int,
	) (domain.Product, error)

	DeleteProduct(
		ctx context.Context,
		id int,
	) error

	PatchProduct(
		ctx context.Context,
		id int,
		product domain.Product,
	) (domain.Product, error)
}

func NewProductsService(
	productsRepository ProductsRepository,
) *ProductsService {
	return &ProductsService{
		productsRepository: productsRepository,
	}
}
