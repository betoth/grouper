package routes

import (
	"grouper/adapter/input/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(userController controller.UserControllerInterface, r *mux.Router) {

	r.HandleFunc("/user", userController.CreateUser).Methods(http.MethodPost)

}
