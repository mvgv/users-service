package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mvgv/users-service/validation"

	"github.com/mvgv/users-service/service"
)

/*GetUsers é reponsavel por recuperar a lista de usuários do sistema*/
func GetUsers(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	userID, err := strconv.ParseUint(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		userErr := &validation.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
		}
		json, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write(json)
		return
	}

	user, userErr := service.GetUser(userID)
	if userErr != nil {
		json, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write(json)
		return
	}

	json, _ := json.Marshal(user)
	resp.Write(json)
}
