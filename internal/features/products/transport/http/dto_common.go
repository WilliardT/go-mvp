package products_transport_http

import (
	"time"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

type ProductDTOResponse struct {
	ID           int       `json:"id"              example:"10"`
	Version      int       `json:"version"         example:"1"`
	Title        string    `json:"title"           example:"MacBook Pro 14"`
	Description  *string   `json:"description"     example:"Ноутбук в отличном состоянии"`
	Price        float64   `json:"price"           example:"149990.50"`
	CreatedAt    time.Time `json:"created_at"      example:"2026-03-30T12:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at"      example:"2026-03-30T12:00:00Z"`
	AuthorUserID int       `json:"author_user_id"  example:"1"`
}

func productDTOFromDomain(product domain.Product) ProductDTOResponse {
	return ProductDTOResponse{
		ID:           product.ID,
		Version:      product.Version,
		Title:        product.Title,
		Description:  product.Description,
		Price:        product.Price,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
		AuthorUserID: product.AuthorUserID,
	}
}

func productsDTOFromDomains(products []domain.Product) []ProductDTOResponse {
	productsDTO := make([]ProductDTOResponse, len(products))

	for i, product := range products {
		productsDTO[i] = productDTOFromDomain(product)
	}

	return productsDTO
}
