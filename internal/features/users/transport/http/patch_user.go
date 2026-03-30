package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_http_request "github.com/WilliardT/go-mvp/internal/core/transport/http/request"
	core_http_response "github.com/WilliardT/go-mvp/internal/core/transport/http/response"
	core_http_types "github.com/WilliardT/go-mvp/internal/core/transport/http/types"
)


type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"     swaggertype:"string"    example:"Василий Петрович"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"  swaggertype:"string"    example:"+79998887766"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("FullName cant`t be NULL")
		}

		fullNameLen := len([]rune(*r.FullName.Value))

		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("FullName length must be between 3 and 100 characters")
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value == nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("PhoneNumber length must be between 10 and 15 characters")
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("PhoneNumber must start with a '+' symbol")
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

// PatchUser    godoc
// @Summary     Изменение пользователя
// @Description Частично обновить информацию существующего пользователя по его ID с указанными данными
// @Description ### Логика обновление полей:
// @Description 1. **Поле не передано**: `phone_number` игнорируется, значение в БД не меняется
// @Description 2. **Явно передано значение**: `"phone_number": "+79998887766"` - устанавливается новый номер телефона в БД
// @Description 3. **Передан  null**: `"phone_number": null` - номер телефона удаляется из БД (устанавливается в `NULL`)
// @Description Ограничения: `full_name` не может быть выставлен как null
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       id path int true "ID изменяемого пользователя"
// @Param       request body PatchUserRequest true "Данные (body) для обновления пользователя"
// @Success     200 {object} PatchUserResponse "Пользователь успешно обновлён"
// @Failure     400 {object} core_http_response.ErrorResponse "Некорректный запрос"
// @Failure     404 {object} core_http_response.ErrorResponse "Пользователь не найден"
// @Failure     409 {object} core_http_response.ErrorResponse "Конфликт данных при обновлении пользователя"
// @Failure     422 {object} core_http_response.ErrorResponse "Невалидные данные для обновления пользователя"
// @Failure     500 {object} core_http_response.ErrorResponse "Внутренняя ошибка сервера"
// @Router      /users/{id} [patch]
func (h *UsersHTTPHandler) PatchUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"falied to get userID path value",
		)

		return
	}

	var request PatchUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)

		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)

	// log.Debug(
	// 	fmt.Sprintf(
	// 		"PatchUserRequest fields:\nFullName: '%v'\nPhoneNumber: '%v'",
	// 		request.FullName,
	// 		request.PhoneNumber,
	// 	),
	// )

	// rw.WriteHeader(http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}
