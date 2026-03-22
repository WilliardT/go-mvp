package users_transport_http

import (
	"net/http"

	core_http_server "github.com/WilliardT/go-mvp/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService usersService
}

type usersService interface{}

func NewUsersHTTPHandler(
	usersService usersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
