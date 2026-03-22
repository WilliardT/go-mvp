package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
	core_http_request "github.com/WilliardT/go-mvp/internal/core/transport/http/request"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"    validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("CreateUser called")

	var request CreateUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Fprintf(rw, "invalid request body: %v", err)
	}

	rw.WriteHeader(http.StatusOK)
}
