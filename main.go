package main

import (
	"database/sql"
	"grouper/adapter/input/controller"
	"grouper/adapter/input/controller/routes"
	"grouper/adapter/output/repository"
	"grouper/application/services"
	"grouper/config"
	"grouper/config/database/postgres"
	"grouper/config/logger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {

	cfg := config.NewConfig()

	logger.Init(cfg)

	database, err := postgres.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := mux.NewRouter()

	routes.InitRoutes(userController, router)

	logger.Info("Init server",
		zap.String("journey", "main"),
	)

	http.ListenAndServe(":8080", router)
}

func initDependencies(db *sql.DB) controller.UserControllerInterface {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserServices(userRepo)
	return controller.NewUserControllerInterface(userService)
}