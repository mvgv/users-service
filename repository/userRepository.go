package repository

import (
	"time"
)

var (
	user = map[string]*User{
		"123": &User{1, "Marcus", "Grilo", 0, time.Now(), time.Now(),
			time.Now(), "https://www.avatar.com.br/avat.png",
			"ADMIN", "Gundam Pilot", "Assinatura dahora", "mvgv1989@gmail.com", "ATIVO", "aweqweq1"},
	}
)

/*UserRepository implementa UserDAO para recuperar e manipular usuarios*/

func getUserByID(userID uint64) (User, error) {
	return User
}
