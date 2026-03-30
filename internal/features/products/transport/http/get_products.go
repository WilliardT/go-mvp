package products_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_http_request "github.com/WilliardT/go-mvp/internal/core/transport/http/request"
	core_http_response "github.com/WilliardT/go-mvp/internal/core/transport/http/response"
)

type GetProductsResponse []ProductDTOResponse

// GetProducts  godoc
// @Summary     Список продуктов
// @Description Получить список продуктов с опциональной фильтрацией по автору карточки продукта и пагинацией
// @Tags        products
// @Produce     json
// @Param       author_user_id query int false "ID автора карточки продукта"
// @Param       limit query int false "Количество продуктов (items) для получения"
// @Param       offset query int false "Количество продуктов для (items) пропуска от начала списка"
// @Success     200 {array} GetProductsResponse "Продукты успешно получены"
// @Failure     400 {object} core_http_response.ErrorResponse "Некорректный запрос"
// @Failure     500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router      /products [get]
func (h *ProductsHTTPHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, limit, offset, err := getUsedIdLimitOffsetQueryParams(r)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'author_user_id / limit / offset' query param",
		)

		return
	}

	productDomains, err := h.productsService.GetProducts(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get products",
		)

		return
	}

	response := GetProductsResponse(productsDTOFromDomains(productDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUsedIdLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		usedIdQueryParamKey = "author_user_id"
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	userID, err := core_http_request.GetIntQueryParam(r, usedIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'author_user_id' query param: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)

	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)

	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return userID, limit, offset, nil
}
