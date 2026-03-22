package users_service

import (
	"context"
	"fmt"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}
	
}