package service

import (
	"github.com/mvgv/users-service/model"
	"github.com/mvgv/users-service/repository"
)

/*GetUsers servico de negocio para listar usuarios*/
func GetUsers() []model.User {
	return []model.User{}
}

/*GetUser servico de negocio para detalhar um usuario*/
func GetUser(userID uint64) (*model.User, error) {
	return repository.GetUserByID(userID)
}
