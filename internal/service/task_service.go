package service

import (
	"context"
	"errors"

	"github.com/alikurb12/todo-app-go/internal/model"
	"github.com/alikurb12/todo-app-go/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTaks(ctx context.Context) ([]model.Task, error) {
	return s.repo.GetALL(ctx)
}

func (s *TaskService) GetTaskById(ctx context.Context, id int64) (*model.Task, error) {
	task, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) CreateTask(ctx context.Context, task *model.Task) error {
	if task.Title == "" {
		return errors.New("Task title cannot be empty")
	}
	return s.repo.Create(ctx, task)
}

func (s *TaskService) UpdateTask(ctx context.Context, id int64, updatedTask *model.Task) error {
	existingTask, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	existingTask.Title = updatedTask.Title
	existingTask.Description = updatedTask.Description
	existingTask.Completed = updatedTask.Completed

	return s.repo.Update(ctx, existingTask)
}

func (s *TaskService) DeleteTask(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
