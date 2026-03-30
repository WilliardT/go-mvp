package users_transport_http

import "github.com/WilliardT/go-mvp/internal/core/domain"


type UserDTOResponse struct {
	ID          int     `json:"id".            example:"10"`
	Version     int     `json:"version"        example:"10"`
	FullName    string  `json:"full_name"      example:"Василий Петрович"`
	PhoneNumber *string `json:"phone_number"   example:"+79998887766"`
}


func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}