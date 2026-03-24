package users_transport_http

import (
	"net/http"

	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
)


type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	re
} 