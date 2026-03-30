package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_http_request "github.com/WilliardT/go-mvp/internal/core/transport/http/request"
	core_http_response "github.com/WilliardT/go-mvp/internal/core/transport/http/response"
)

type ProductPriceRangesResponse struct {
	Cheap     int `json:"cheap"      example:"2"`
	Medium    int `json:"medium"     example:"5"`
	Expensive int `json:"expensive"  example:"1"`
}

type GetStatisticsResponse struct {
	ProductsCount      int                        `json:"products_count"       example:"8"`
	ProductPriceMedian *float64                   `json:"product_price_median" example:"499.99"`
	ProductPriceRanges ProductPriceRangesResponse `json:"price_ranges"`
}

// GetProductsStatistics godoc
// @Summary     Статистика по продуктам
// @Description Получить агрегированную статистику по продуктам с опциональной фильтрацией по автору и диапазону дат создания
// @Tags        statistics
// @Produce     json
// @Param       author_user_id query int false "ID автора продукта"
// @Param       created_from query string false "Начальная дата фильтра по created_at в формате YYYY-MM-DD"
// @Param       created_to query string false "Конечная дата фильтра по created_at в формате YYYY-MM-DD"
// @Success     200 {object} GetStatisticsResponse "Статистика успешно получена"
// @Failure     400 {object} core_http_response.ErrorResponse "Некорректный запрос"
// @Failure     500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router      /statistics/products [get]
func (h *StatisticsHTTPHandler) GetProductsStatistics(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, from, to, err := getUserIDFromToQueryParams(r)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID/from/to query params",
		)

		return
	}

	statistics, err := h.statisticsService.GetProductsStatistics(
		ctx,
		userID,
		from,
		to,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get statistics",
		)

		return
	}

	responseHandler.JSONResponse(
		toDTOFromDomain(statistics),
		http.StatusOK,
	)
}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	return GetStatisticsResponse{
		ProductsCount:      statistics.ProductsCount,
		ProductPriceMedian: statistics.ProductPriceMedian,
		ProductPriceRanges: ProductPriceRangesResponse{
			Cheap:     statistics.ProductPriceRanges.Cheap,
			Medium:    statistics.ProductPriceRanges.Medium,
			Expensive: statistics.ProductPriceRanges.Expensive,
		},
	}
}

func getUserIDFromToQueryParams(r *http.Request) (
	*int,
	*time.Time,
	*time.Time,
	error,
) {
	const (
		userIDQueryParamKey = "author_user_id"
		fromQueryParamKey   = "created_from"
		toQueryParamKey     = "created_to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)

	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"get 'author_user_id' query param: %w",
			err,
		)
	}

	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamKey)

	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"get 'created_from' query param: %w",
			err,
		)
	}

	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)

	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"get 'created_to' query param: %w",
			err,
		)
	}

	return userID, from, to, nil
}
