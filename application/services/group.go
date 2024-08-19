package services

import (
	"fmt"
	"grouper/application/domain"
	"grouper/application/dto"
	"grouper/application/errors"
	"grouper/application/port/input"
	"grouper/application/port/output"
	"grouper/config/logger"
	"grouper/config/rest_errors"

	"go.uber.org/zap"
)

func NewGroupService(repos GroupService) input.GroupService {
	return &GroupService{
		repos.RepoGroup,
		repos.RepoTopic,
		repos.RepoUser,
	}
}

type GroupService struct {
	RepoGroup output.GroupPort
	RepoTopic output.TopicPort
	RepoUser  output.UserPort
}

func (service *GroupService) CreateGroup(groupDomain domain.Group) (*dto.Group, error) {
	logger.Debug("Init CreateGroup service", zap.String("journey", "CreateGroup"))
	groupRepo, err := service.RepoGroup.CreateGroup(groupDomain)

	if err != nil {
		return nil, errors.HandleServiceError(err, "Group", "CreateGroup")
	}

	groupDto := dto.Group{
		ID:       groupRepo.ID,
		Name:     groupRepo.Name,
		UserName: "UserNameDTO",
		Topic: dto.GroupTopic{
			ID:   "TopicDto ID",
			Name: "TopicDto Name",
			Subtopic: dto.GroupSubtopic{
				ID:   "SubtopicDto ID",
				Name: "SubtopicDTO Name",
			},
		},
		CreatedAt: groupRepo.CreatedAt,
	}

	logger.Debug("Finish CreateGroup service", zap.String("journey", "CreateGroup"))
	return &groupDto, nil
}

func (service *GroupService) JoinGroup(userID, groupID string) error {
	logger.Debug("Init Join service", zap.String("journey", "JoinGroup"))
	err := service.RepoGroup.JoinGroup(userID, groupID)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "JoinGroup"))
		return err
	}
	logger.Debug("Finish Join service", zap.String("journey", "JoinGroup"))
	return nil
}

func (service *GroupService) LeaveGroup(userID, groupID string) *rest_errors.RestErr {
	logger.Debug("Init Leave service", zap.String("journey", "LeaveGroup"))
	err := service.RepoGroup.LeaveGroup(userID, groupID)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "LeaveGroup"))
		return err
	}
	logger.Debug("Init Leave service", zap.String("journey", "LeaveGroup"))
	return nil
}

func (service *GroupService) GetGroups(parameter dto.GetGroupsParameter) (*[]dto.Group, *rest_errors.RestErr) {
	logger.Debug("Init GetGroups service", zap.String("journey", "GetGroups"))
	queryParameter := dto.GetGroupsParameter{
		Name: parameter.Name,
	}

	groupsRepo, err := service.RepoGroup.GetGroups(queryParameter)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "GetGroups"))
		return nil, err
	}
	if groupsRepo == nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "GetGroups"))
		return nil, err
	}

	var groupsDto []dto.Group
	for _, groupRepo := range *groupsRepo {

		groupDto := dto.Group{
			ID:       groupRepo.ID,
			Name:     groupRepo.Name,
			UserName: "UserNameDTO",
			Topic: dto.GroupTopic{
				ID:   "TopicDto ID",
				Name: "TopicDto Name",
				Subtopic: dto.GroupSubtopic{
					ID:   "SubtopicDto ID",
					Name: "SubtopicDTO Name",
				},
			},
			CreatedAt: groupRepo.CreatedAt,
		}
		groupsDto = append(groupsDto, groupDto)

	}

	logger.Debug("Finish GetGroups service", zap.String("journey", "GetGroups"))
	return &groupsDto, nil
}

func (service *GroupService) FindByID(groupID string) (*dto.Group, error) {
	logger.Debug("Init FindByID service", zap.String("journey", "FindByID"))

	groupRepo, err := service.RepoGroup.FindByID(groupID)
	if err != nil {
		return nil, errors.HandleServiceError(err, "Group", "FindByID")
	}

	topicRepo, err := service.RepoTopic.FindByID(groupRepo.TopicID)
	if err != nil {
		return nil, errors.HandleServiceError(err, "Group", "FindByID")
	}

	userRepo, err := service.RepoUser.FindByID(groupRepo.UserID)
	fmt.Println(err)

	groupDto := dto.Group{
		ID:       groupRepo.ID,
		Name:     groupRepo.Name,
		UserName: userRepo.Name,
		Topic: dto.GroupTopic{
			ID:   topicRepo.ID,
			Name: topicRepo.Name,
			Subtopic: dto.GroupSubtopic{
				ID:   "SubtopicDto ID",
				Name: "SubtopicDTO Name",
			},
		},
		CreatedAt: groupRepo.CreatedAt,
	}

	return &groupDto, nil
}
