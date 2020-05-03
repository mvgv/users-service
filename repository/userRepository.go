package repository

import (
	"errors"
	"time"

	"github.com/mvgv/users-service/model"
)

var (
	selectedUser = map[uint64]*model.User{
		123: &model.User{1, "Marcus", "Grilo", 0, time.Now(), time.Now(),
			time.Now(), "https://www.avatar.com.br/avat.png",
			"ADMIN", "Gundam Pilot", "Assinatura dahora", "mvgv1989@gmail.com", "ATIVO", "aweqweq1"},
	}
)

/*GetUserByID implementa busca um usuario pelo seu id*/
func GetUserByID(userID uint64) (*model.User, error) {
	user := selectedUser[userID]
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
