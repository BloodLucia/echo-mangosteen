package service

type userService struct {
}

// UserLogin implements UserService.
func (*userService) UserLogin() error {
	panic("unimplemented")
}

type UserService interface {
	UserLogin() error
}

func NewUserService() UserService {
	return &userService{}
}
