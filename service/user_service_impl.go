package service

import (
	"github.com/mvgv/users-service/model"
	"github.com/mvgv/users-service/repository"
	"github.com/mvgv/users-service/validation"
)

var (
	/*UserServiceImpl implementa a interface definida para UserService*/
	UserServiceImpl UserService
)

type userServiceImpl struct{}

func init() {
	UserServiceImpl = &userServiceImpl{}
}

/*GetUsers servico de negocio para listar usuarios*/
func (userServiceImpl *userServiceImpl) GetUsers() []model.User {
	return []model.User{}
}

/*GetUser servico de negocio para detalhar um usuario*/
func (userServiceImpl *userServiceImpl) GetUser(userID uint64) (*model.User, *validation.ApplicationError) {
	return repository.UserRepositoryImpl.GetUserByID(userID)
}
