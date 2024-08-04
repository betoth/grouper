package routes

import (
	"grouper/adapter/input/controller"
	"grouper/adapter/input/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	UserController  controller.UserControllerInterface
	GroupController controller.GroupControllerInterface
}

func InitRoutes(routes Routes, r *mux.Router) {

	var ApiV1 = r.PathPrefix("/v1").Subrouter()
	ApiV1.HandleFunc("/user", routes.UserController.CreateUser).Methods(http.MethodPost)
	ApiV1.HandleFunc("/user/login", routes.UserController.Login).Methods(http.MethodPost)

	ApiV1.HandleFunc("/group", middleware.Auth(routes.GroupController.CreateGroup)).Methods(http.MethodPost)
	ApiV1.HandleFunc("/group/{groupId}/join", middleware.Auth(routes.GroupController.Join)).Methods(http.MethodPost)
	ApiV1.HandleFunc("/group/{groupId:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}/leave", middleware.Auth(routes.GroupController.Leave)).Methods(http.MethodPost)

}
