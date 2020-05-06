package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mvgv/users-service/controller"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func urlMapping() {
	router.GET("/users/:user_id", controller.GetUser)
}

/*StartApplication inicializa as rotas do webserver*/
func StartApplication() {

	urlMapping()
	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}
