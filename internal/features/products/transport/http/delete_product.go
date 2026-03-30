package products_transport_http

import (
	"net/http"

	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_http_request "github.com/WilliardT/go-mvp/internal/core/transport/http/request"
	core_http_response "github.com/WilliardT/go-mvp/internal/core/transport/http/response"
)

// DeleteProduct godoc
// @Summary      Удалить продукт (item)
// @Description  Удалить существующий продукт (item) по его ID
// @Tags         products
// @Param        id path int true "ID продукта"
// @Success      204 "Продукт успешно удалён"
// @Failure      400 {object} core_http_response.ErrorResponse "Некорректный запрос"
// @Failure      404 {object} core_http_response.ErrorResponse "Продукт (item) не найден"
// @Failure      500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /products/{id} [delete]
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
