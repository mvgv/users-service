package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mvgv/users-service/service"
)

/*GetUsers é reponsavel por recuperar a lista de usuários do sistema*/
func GetUsers(resp http.ResponseWriter, req *http.Request) {
	userID, _ := strconv.ParseUint(req.URL.Query().Get("user_id"), 10, 64)
	user, _ := service.GetUser(userID)
	json, _ := json.Marshal(user)
	resp.Write(json)
}
