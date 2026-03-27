package products_postgres_repository

import (
	"time"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

type ProductModel struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      int
}

func productDomainFromModel(product ProductModel) domain.Product {
	return domain.NewProduct(
		product.ID,
		product.Version,
		product.Title,
		product.Description,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
		product.UserID,
	)
}

func productDomainsFromModels(products []ProductModel) []domain.Product {
	productDomains := make([]domain.Product, len(products))

	for i, product := range products {
		productDomains[i] = productDomainFromModel(product)
	}

	return productDomains
}
