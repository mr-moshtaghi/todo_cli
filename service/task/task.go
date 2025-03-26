package task

import (
	"fmt"
	"todo_cli/entity"
)

type ServiceRepository interface {
	DoesThisUserHaveThisCategoryID(userID, categoryID int) bool
	CreateNewTask(t entity.Task) (entity.Task, error)
	ListUserTask(userID int) ([]entity.Task, error)
}
type Service struct {
	repository ServiceRepository
}

func NewService(repo ServiceRepository) Service {
	return Service{repository: repo}
}

type CreateRequest struct {
	Title              string
	DueDate            string
	CategoryId         int
	AuthenticateUserId int
}

type CreateResponse struct {
	Task entity.Task
}

type ListRequest struct {
	UserID int
}

type ListResponse struct {
	Tasks []entity.Task
}

func (s Service) CreatTask(req CreateRequest) (CreateResponse, error) {

	//ok := s.repository.DoesThisUserHaveThisCategoryID(req.AuthenticateUserId, req.CategoryId)
	//if !ok {
	//	return CreateResponse{}, fmt.Errorf("user doesn't have this category: %d", req.CategoryId)
	//}

	task := entity.Task{
		Title:      req.Title,
		DueDate:    req.DueDate,
		CategoryID: req.CategoryId,
		IsDone:     false,
		UserID:     req.AuthenticateUserId,
	}

	createdTask, cErr := s.repository.CreateNewTask(task)
	if cErr != nil {
		return CreateResponse{}, fmt.Errorf("create new task failed: %v", cErr)
	}
	return CreateResponse{Task: createdTask}, nil
}

func (s Service) List(req ListRequest) (ListResponse, error) {
	tasks, err := s.repository.ListUserTask(req.UserID)
	if err != nil {
		return ListResponse{}, fmt.Errorf("list user tasks failed: %v", err)
	}
	return ListResponse{Tasks: tasks}, nil
}
