package usecase

import (
	"Nurul-trial/models"
	"Nurul-trial/repository"
)

type UcaseStruct struct {
	repo repository.RepoLoginInterface
}

func NewImplUcase(repo repository.RepoLoginInterface) UcaseLoginInterface {
	return UcaseStruct{repo: repo}
}

func (a UcaseStruct) LoginUser(username string, password string, email string) (models.GetItGuys, error) {
	resp, err := a.repo.LoginUser(username, password, email)
	if err != nil {
		return resp, err
	}
	return models.GetItGuys{}, nil
}
