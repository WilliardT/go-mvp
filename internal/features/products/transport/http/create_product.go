package products_transport_http

import (
	"net/http"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_http_request "github.com/WilliardT/go-mvp/internal/core/transport/http/request"
	core_http_response "github.com/WilliardT/go-mvp/internal/core/transport/http/response"
)

type CreateProductRequest struct {
	Title        string  `json:"title"            validate:"required,min=1,max=100"   example:"MacBook Pro 14"`
	Description  *string `json:"description"      validate:"omitempty,min=1,max=100"  example:"Ноутбук в отличном состоянии"`
	Price        float64 `json:"price"            validate:"required,gt=0"            example:"149990.50"`
	AuthorUserID int     `json:"author_user_id"   validate:"required,gt=0"            example:"1"`
}

type CreateProductResponse ProductDTOResponse

// CreateProduct   godoc
// @Summary        Создать продукт
// @Description    Создать новый продукт (item) с указанными данными
// @Tags           products
// @Accept         json
// @Produce        json
// @Param          request body CreateProductRequest true "Данные для создания продукта"
// @Success        201 {object} CreateProductResponse "Продукт успешно создан"
// @Failure        400 {object} core_http_response.ErrorResponse "Некорректный запрос"
// @Failure        404 {object} core_http_response.ErrorResponse "Пользователь (автор карточки продукта) не найден"
// @Failure        500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router         /products [post]
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
		dto.AuthorUserID,
	)
}
