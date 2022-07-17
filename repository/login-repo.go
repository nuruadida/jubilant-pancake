package repository

import (
	"Nurul-trial/models"
	"gorm.io/gorm"
)

type RepoStruct struct {
	Db *gorm.DB
}

func NewImplRepo(db *gorm.DB) RepoLoginInterface {
	return &RepoStruct{Db: db}
}

func (a RepoStruct) LoginUser(username string, email string, password string) (models.GetItGuys, error) {

	var data models.GetItGuys
	//tx := a.Db.Begin()
	err := a.Db.Debug().Where("usrernmae =? and email = ? and password = ?", username, email, password).Find(data).Error
	if err != nil {
		//tx.Rollback()
		return models.GetItGuys{}, err
	}
	//tx.Commit()
	return data, nil
}
