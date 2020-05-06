package service

import (
	"github.com/mvgv/users-service/model"
	"github.com/mvgv/users-service/validation"
)

/*UserService é a interface que expoe os contratos para as logicas de negocio sobre um usuario*/
type UserService interface {
	GetUser(userID uint64) (*model.User, *validation.ApplicationError)
	GetUsers() []model.User
}
