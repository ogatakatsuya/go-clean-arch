package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetByEmail(user *model.User, email string) error
	Create(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Create(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
