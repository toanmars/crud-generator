package services

import (
    "github.com/toanmars/crud-generator/models"
    "github.com/toanmars/crud-generator/repositories"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Create(model *models.User) error {
    return s.repo.Create(model)
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
    return s.repo.FindByID(id)
}

func (s *UserService) Update(model *models.User) error {
    return s.repo.Update(model)
}

func (s *UserService) Delete(id uint) error {
    return s.repo.Delete(id)
}