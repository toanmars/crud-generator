package repositories

import (
	"github.com/toanmars/crud-generator/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(model *models.User) error {
	return r.db.Create(model).Error
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var model models.User
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *UserRepository) Update(model *models.User) error {
	return r.db.Save(model).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
