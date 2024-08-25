package controller

import (
	"grouper/adapter/input/converter"
	"grouper/adapter/input/httperror"
	"grouper/adapter/input/response"
	"grouper/application/port/input"
	"grouper/config/logger"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewTopicController(serviceInterface input.TopicService) TopicController {

	return &topicController{
		service: serviceInterface,
	}
}

type TopicController interface {
	FindByID(w http.ResponseWriter, r *http.Request)
}

type topicController struct {
	service input.TopicService
}

func (ctrl *topicController) FindByID(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init FindByID controller", zap.String("journey", "FindByID"))
	parameter := mux.Vars(r)
	groupID := parameter["topicId"]

	domainResult, err := ctrl.service.FindByID(groupID)

	if err != nil {
		logger.Error("Error trying to call FindByID service", err, zap.String("journey", "FindByID"))
		httperror.MapAndRespond(w, err, "FindByID")
		return
	}
	groupResponse := converter.ConvertTopicDomainToResponse(domainResult)
	logger.Debug("Finish FindByID controller", zap.String("journey", "FindByID"))
	response.JSON(w, http.StatusOK, groupResponse)
}
