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
	Title        core_http_types.Nullable[string]  `json:"title"            swaggertype:"string"  example:"MacBook Pro 14"`
	Description  core_http_types.Nullable[string]  `json:"description"      swaggertype:"string"  example:"Ноутбук в отличном состоянии"`
	Price        core_http_types.Nullable[float64] `json:"price"            swaggertype:"number"  example:"149990.50"`
	AuthorUserID core_http_types.Nullable[int]     `json:"author_user_id"   swaggertype:"integer" example:"1"`
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

	if r.AuthorUserID.Set {
		if r.AuthorUserID.Value == nil {
			return fmt.Errorf("AuthorUserID cant`t be NULL")
		}

		if *r.AuthorUserID.Value <= 0 {
			return fmt.Errorf("AuthorUserID must be greater than 0")
		}
	}

	return nil
}

type PatchProductResponse ProductDTOResponse

// PatchProduct godoc
// @Summary     Изменение продукта
// @Description Частично обновить информацию существующего продукта по его ID
// @Description ### Логика обновления полей:
// @Description 1. **Поле не передано**: значение в БД не меняется
// @Description 2. **Явно передано значение**: устанавливается новое значение поля
// @Description 3. **Передан null**: допустимо только для `description`, поле будет очищено
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id path int true "ID изменяемого продукта"
// @Param       request body PatchProductRequest true "Данные (body) для обновления продукта"
// @Success     200 {object} PatchProductResponse "Продукт успешно обновлён"
// @Failure     400 {object} core_http_response.ErrorResponse "Некорректный запрос"
// @Failure     404 {object} core_http_response.ErrorResponse "Продукт не найден"
// @Failure     409 {object} core_http_response.ErrorResponse "Конфликт данных при обновлении продукта"
// @Failure     500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router      /products/{id} [patch]
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
		request.AuthorUserID.ToDomain(),
	)
}
