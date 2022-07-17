package repository

import (
	"Nurul-trial/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewImplUserRepo(db *gorm.DB) RepoUserInterface {
	return UserRepo{DB: db}
}

func (u UserRepo) AddUser(v models.User) (models.User, error) {
	err := u.DB.Debug().Create(&v).Error
	if err != nil {
		return models.User{}, err
	}
	return v, nil
}
