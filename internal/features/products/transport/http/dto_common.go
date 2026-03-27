package products_transport_http

import (
	"time"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

type ProductDTOResponse struct {
	ID          int       `json:"id"`
	Version     int       `json:"version"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
}

func productDTOFromDomain(product domain.Product) ProductDTOResponse {
	return ProductDTOResponse{
		ID:          product.ID,
		Version:     product.Version,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		UserID:      product.UserID,
	}
}

func productsDTOFromDomains(products []domain.Product) []ProductDTOResponse {
	productsDTO := make([]ProductDTOResponse, len(products))

	for i, product := range products {
		productsDTO[i] = productDTOFromDomain(product)
	}

	return productsDTO
}
