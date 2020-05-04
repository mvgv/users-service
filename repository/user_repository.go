package repository

import (
	"github.com/mvgv/users-service/model"
	"github.com/mvgv/users-service/validation"
)

/*UserRepository define a interface de acesso a base de clientes*/
type UserRepository interface {
	GetUserByID(userID uint64) (*model.User, *validation.ApplicationError)
}
