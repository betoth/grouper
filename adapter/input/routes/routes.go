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

const UUIDPattern = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"

func InitRoutes(routes Routes, r *mux.Router) {

	var ApiV1 = r.PathPrefix("/v1").Subrouter()
	ApiV1.Use(middleware.LogRequest)

	ApiV1.HandleFunc("/users", routes.UserController.CreateUser).Methods(http.MethodPost)
	ApiV1.HandleFunc("/users/login", routes.UserController.Login).Methods(http.MethodPost)
	ApiV1.HandleFunc("/users/{userId:"+UUIDPattern+"}/groups", middleware.Auth(routes.UserController.GetUserGroups)).Methods(http.MethodGet)

	ApiV1.HandleFunc("/groups", middleware.Auth(routes.GroupController.CreateGroup)).Methods(http.MethodPost)
	ApiV1.HandleFunc("/groups/{groupId:"+UUIDPattern+"}/join", middleware.Auth(routes.GroupController.Join)).Methods(http.MethodPost)
	ApiV1.HandleFunc("/groups/{groupId:"+UUIDPattern+"}/leave", middleware.Auth(routes.GroupController.Leave)).Methods(http.MethodPost)
	ApiV1.HandleFunc("/groups", middleware.Auth(routes.GroupController.GetGroups)).Methods(http.MethodGet)
}
