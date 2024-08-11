package routes

import (
	"grouper/adapter/input/controller"
	"grouper/adapter/input/middleware"
	"grouper/adapter/input/util"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	UserController  controller.UserController
	GroupController controller.GroupController
}

func InitRoutes(routes Routes, r *mux.Router) {

	var ApiV1 = r.PathPrefix("/v1").Subrouter()
	ApiV1.Use(middleware.RequestLogger)

	ApiV1.HandleFunc("/users", routes.UserController.Create).Methods(http.MethodPost)
	ApiV1.HandleFunc("/users/login", routes.UserController.Login).Methods(http.MethodPost)
	ApiV1.HandleFunc("/users/{userId:"+util.UUIDPattern+"}/groups", middleware.Auth(routes.UserController.GetGroups)).Methods(http.MethodGet)

	ApiV1.HandleFunc("/groups", middleware.Auth(routes.GroupController.Create)).Methods(http.MethodPost)
	ApiV1.HandleFunc("/groups/{groupId:"+util.UUIDPattern+"}/join", middleware.Auth(routes.GroupController.Join)).Methods(http.MethodPost)
	ApiV1.HandleFunc("/groups/{groupId:"+util.UUIDPattern+"}/leave", middleware.Auth(routes.GroupController.Leave)).Methods(http.MethodPost)
	ApiV1.HandleFunc("/groups", middleware.Auth(routes.GroupController.GetGroups)).Methods(http.MethodGet)
}
