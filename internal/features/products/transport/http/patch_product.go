package products_transport_http

import (
	"fmt"
	"net/http"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_http_request "github.com/WilliardT/go-mvp/internal/core/transport/http/request"
	core_http_response "github.com/WilliardT/go-mvp/internal/core/transport/http/response"
	core_http_types "github.com/WilliardT/go-mvp/internal/core/transport/http/types"
)

type PatchProductRequest struct {
	Title       core_http_types.Nullable[string]  `json:"title"`
	Description core_http_types.Nullable[string]  `json:"description"`
	Price       core_http_types.Nullable[float64] `json:"price"`
	UserID      core_http_types.Nullable[int]     `json:"user_id"`
}

func (r *PatchProductRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("Title cant`t be NULL")
		}

		titleLength := len([]rune(*r.Title.Value))

		if titleLength < 1 || titleLength > 100 {
			return fmt.Errorf("Title length must be between 1 and 100 characters")
		}
	}

	if r.Description.Set && r.Description.Value != nil {
		descriptionLength := len([]rune(*r.Description.Value))

		if descriptionLength < 1 || descriptionLength > 100 {
			return fmt.Errorf("Description length must be between 1 and 100 characters")
		}
	}

	if r.Price.Set {
		if r.Price.Value == nil {
			return fmt.Errorf("Price cant`t be NULL")
		}

		if *r.Price.Value <= 0 {
			return fmt.Errorf("Price must be greater than 0")
		}
	}

	if r.UserID.Set {
		if r.UserID.Value == nil {
			return fmt.Errorf("UserID cant`t be NULL")
		}

		if *r.UserID.Value <= 0 {
			return fmt.Errorf("UserID must be greater than 0")
		}
	}

	return nil
}

type PatchProductResponse ProductDTOResponse

func (h *ProductsHTTPHandler) PatchProduct(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	productID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get productID path value",
		)

		return
	}

	var request PatchProductRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	productPatch := productPatchFromRequest(request)

	productDomain, err := h.productsService.PatchProduct(ctx, productID, productPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch product",
		)

		return
	}

	response := PatchProductResponse(productDTOFromDomain(productDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func productPatchFromRequest(request PatchProductRequest) domain.ProductPatch {
	return domain.NewProductPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Price.ToDomain(),
		request.UserID.ToDomain(),
	)
}
