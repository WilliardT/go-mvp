package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/WilliardT/go-mvp/internal/core/logger"
)

type CreateUserRequest struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
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

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Fprintf(rw, "invalid request body: %v", err)
	}

	rw.WriteHeader(http.StatusOK)
}
