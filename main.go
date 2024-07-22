package main

import (
	"context"
	"database/sql"
	"fmt"
	"grouper/adapter/input/controller"
	"grouper/adapter/input/controller/routes"
	"grouper/adapter/output/repository"
	services "grouper/application/services/user"
	"grouper/configuration/database/postgres"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database, err := postgres.NewPostgresConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}
	userController := initDependencies(database)
	router := mux.NewRouter()

	routes.InitRoutes(userController, router)
	fmt.Println("Iniciando servidor")
	http.ListenAndServe(":8080", router)
}

func initDependencies(db *sql.DB) controller.UserControllerInterface {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserServices(userRepo)
	return controller.NewUserControllerInterface(userService)
}
