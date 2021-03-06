package repository

import (
	"net/http"
	"time"

	"github.com/mvgv/users-service/model"
	"github.com/mvgv/users-service/validation"
)

var (
	selectedUser = map[uint64]*model.User{
		123: &model.User{123, "Marcus", "Grilo", 0, time.Now(), time.Now(),
			time.Now(), "https://www.avatar.com.br/avat.png",
			"ADMIN", "Gundam Pilot", "Assinatura dahora", "mvgv1989@gmail.com", "ATIVO", "aweqweq1"},
	}

	selectedUsers = []*model.User{selectedUser[123]}
	/*UserRepositoryImpl expoe a implementacao do repository para outros pacotes*/
	UserRepositoryImpl UserRepository
)

func init() {
	UserRepositoryImpl = &userRepositoryImpl{}
}

type userRepositoryImpl struct{}

/*GetUserByID implementa busca um usuario pelo seu id*/
func (userRepository *userRepositoryImpl) GetUserByID(userID uint64) (*model.User, *validation.ApplicationError) {
	user := selectedUser[userID]
	if user == nil {
		return nil, &validation.ApplicationError{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
		}
	}
	return user, nil
}

func (userRepository *userRepositoryImpl) GetUsers() (*[]*model.User, *validation.ApplicationError) {
	return &selectedUsers, nil
}
