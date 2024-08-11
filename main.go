package main

import (
	"grouper/adapter/input/controller"
	"grouper/adapter/input/routes"
	"grouper/adapter/output/repository"
	"grouper/application/services"
	"grouper/config"
	"grouper/config/database/postgres"
	"grouper/config/logger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()

	logger.Init(cfg)

	database, err := postgres.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	routesController := routes.Routes{}
	routesController.UserController, routesController.GroupController = initDependencies(database)
	router := mux.NewRouter()
	routes.InitRoutes(routesController, router)

	logger.Info("Init server", zap.String("journey", "Bootstrap"))
	http.ListenAndServe(":8080", router)
}

func initDependencies(db *gorm.DB) (controller.UserController, controller.GroupController) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	groupRepo := repository.NewGroupRepository(db)
	groupService := services.NewGroupService(groupRepo)

	return controller.NewUserController(userService), controller.NewGroupController(groupService)
}
