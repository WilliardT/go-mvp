package products_transport_http

import (
	"net/http"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_http_request "github.com/WilliardT/go-mvp/internal/core/transport/http/request"
	core_http_response "github.com/WilliardT/go-mvp/internal/core/transport/http/response"
)

type CreateProductRequest struct {
	Title       string  `json:"title" validate:"required,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1,max=100"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	UserID      int     `json:"user_id" validate:"required,gt=0"`
}

type CreateProductResponse ProductDTOResponse

func (h *ProductsHTTPHandler) CreateProduct(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("CreateProduct called")

	var request CreateProductRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	productDomain := domainFromDTO(request)

	productDomain, err := h.productsService.CreateProduct(ctx, productDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create product")

		return
	}

	response := CreateProductResponse(productDTOFromDomain(productDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateProductRequest) domain.Product {
	return domain.NewProductUninitialized(
		dto.Title,
		dto.Description,
		dto.Price,
		dto.UserID,
	)
}
