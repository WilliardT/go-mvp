package users_transport_http


type UsersHTTPHandler struct {
	usersService usersService
}

type usersService interface {}


func NewUsersHTTPHandler(
	usersService usersService
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}