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

	routesController := initDependencies(database)

	router := mux.NewRouter()
	routes.InitRoutes(routesController, router)

	logger.Info("Init server", zap.String("journey", "Bootstrap"))
	http.ListenAndServe(":8080", router)
}

func initDependencies(db *gorm.DB) *routes.Routes {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	topicRepo := repository.NewTopicRepository(db)
	topicService := services.NewTopicService(topicRepo)

	subtopicRepo := repository.NewSubtopicRepository(db)
	subtopicService := services.NewSubtopicService(subtopicRepo)

	groupRepo := repository.NewGroupRepository(db)
	repos := services.GroupService{
		RepoGroup:    groupRepo,
		RepoTopic:    topicRepo,
		RepoUser:     userRepo,
		RepoSubtopic: subtopicRepo,
	}
	groupService := services.NewGroupService(repos)

	routesController := routes.Routes{
		UserController:     controller.NewUserController(userService),
		GroupController:    controller.NewGroupController(groupService),
		TopicController:    controller.NewTopicController(topicService),
		SubtopicController: controller.NewSubtopicController(subtopicService),
	}

	return &routesController
}
