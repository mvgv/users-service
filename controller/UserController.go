package controller

import (
	"net/http"

	"github.com/mvgv/users-service/service"
)

/*GetUsers é reponsavel por recuperar a lista de usuários do sistema*/
func GetUsers(resp http.ResponseWriter, req *http.Request) {

	users := service.GetUsers()
	resp.Write(users)
}
