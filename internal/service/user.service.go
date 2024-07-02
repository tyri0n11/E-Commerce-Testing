package service

import "github.com/tyri0n11/Muffin/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user service us
func (us *UserService) GetUserInfo() string {
	return us.userRepo.GetInfoUser()
}
