package services

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) ListUsers() []string {
	// Simulação de listagem de usuários

}
