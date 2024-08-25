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

func NewSubtopicController(serviceInterface input.SubtopicService) SubtopicController {

	return &subtopicController{
		service: serviceInterface,
	}
}

type SubtopicController interface {
	FindByID(w http.ResponseWriter, r *http.Request)
}

type subtopicController struct {
	service input.SubtopicService
}

func (ctrl *subtopicController) FindByID(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init FindByID in subtopicController", zap.String("journey", "FindSubtopicByID"))
	parameter := mux.Vars(r)
	subtopicID := parameter["subtopicId"]

	SubtopicDomain, err := ctrl.service.FindByID(subtopicID)

	if err != nil {
		httperror.ErrorToErrorResponse(w, err, "FindByID")
		return
	}

	SubtopicResponse := converter.ConvertSubtopicDomainToResponse(SubtopicDomain)

	logger.Debug("Finish FindByID in subtopicController", zap.String("journey", "FindSubtopicByID"))
	response.JSON(w, http.StatusOK, SubtopicResponse)
}
