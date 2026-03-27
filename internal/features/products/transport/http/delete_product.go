package products_transport_http

import (
	"net/http"

	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_http_request "github.com/WilliardT/go-mvp/internal/core/transport/http/request"
	core_http_response "github.com/WilliardT/go-mvp/internal/core/transport/http/response"
)

func (h *ProductsHTTPHandler) DeleteProduct(
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

	err = h.productsService.DeleteProduct(ctx, productID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete product",
		)

		return
	}

	responseHandler.NoContentResponse()
}
