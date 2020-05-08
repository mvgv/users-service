package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mvgv/users-service/validation"

	"github.com/mvgv/users-service/service"
)

/*GetUser é reponsavel por detalhar um usuario da lista de usuários do sistema*/
func GetUser(context *gin.Context) {
	userID, err := strconv.ParseUint(context.Param("user_id"), 10, 64)
	if err != nil {
		userErr := &validation.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
		}
		context.JSON(userErr.StatusCode, userErr)
		return
	}

	user, userErr := service.UserServiceImpl.GetUser(userID)
	if userErr != nil {
		context.JSON(userErr.StatusCode, userErr)
		return
	}

	context.JSON(http.StatusOK, user)
}

/*GetUsers é reponsavel por detalhar um usuario da lista de usuários do sistema*/
func GetUsers(context *gin.Context) {

	user, userErr := service.UserServiceImpl.GetUsers()
	if userErr != nil {
		context.JSON(userErr.StatusCode, userErr)
		return
	}

	context.JSON(http.StatusOK, user)
}
