package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// user repository ur
func (ur *UserRepo) GetInfoUser() string {
	return "Tyr1on"
}
