package routes

import (
	"grouper/adapter/input/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(userController controller.UserControllerInterface, groupController controller.GroupControllerInterface, r *mux.Router) {

	var ApiV1 = r.PathPrefix("/v1").Subrouter()
	ApiV1.HandleFunc("/user", userController.CreateUser).Methods(http.MethodPost)

	ApiV1.HandleFunc("/group", groupController.CreateGroup).Methods(http.MethodPost)

}
