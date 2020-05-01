package app

import (
	"net/http"

	"github.com/mvgv/users-service/controller"
)

/*StartApplication inicializa as rotas do webserver*/
func StartApplication() {

	http.HandleFunc("/users", controller.GetUsers)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
