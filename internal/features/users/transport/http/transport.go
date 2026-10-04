package users_transport_http

type UsersAHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
}

func NewUsersHTTPHandler(usersService UsersService) *UsersAHTTPHandler {
	return &UsersAHTTPHandler{
		usersService: usersService,
	}
}
