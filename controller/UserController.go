package controller

import "net/http"

/*GetUsers é reponsavel por recuperar a lista de usuários do sistema*/
func GetUsers(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Usuario X"))
}
