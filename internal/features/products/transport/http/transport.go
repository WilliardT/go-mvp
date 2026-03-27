package products_transport_http

import (
	"context"
	"net/http"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_http_server "github.com/WilliardT/go-mvp/internal/core/transport/http/server"
)

type ProductsHTTPHandler struct {
	productsService ProductsService
}

type ProductsService interface {
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
		patch domain.ProductPatch,
	) (domain.Product, error)
}

func NewProductsHTTPHandler(
	productsService ProductsService,
) *ProductsHTTPHandler {
	return &ProductsHTTPHandler{
		productsService: productsService,
	}
}

func (h *ProductsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/products",
			Handler: h.CreateProduct,
		},
		{
			Method:  http.MethodGet,
			Path:    "/products",
			Handler: h.GetProducts,
		},
		{
			Method:  http.MethodGet,
			Path:    "/products/{id}",
			Handler: h.GetProduct,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/products/{id}",
			Handler: h.DeleteProduct,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/products/{id}",
			Handler: h.PatchProduct,
		},
	}
}
