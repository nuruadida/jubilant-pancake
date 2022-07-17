package usecase

import (
	"Nurul-trial/models"
	"Nurul-trial/repository"
)

type UserStruct struct {
	UserRepo repository.RepoUserInterface
}

func NewImplUserUcase(UserRepo repository.RepoUserInterface) UcaseUserInterface {
	return &UserStruct{UserRepo: UserRepo}
}

func (u UserStruct) AddUser(v models.User) (models.User, error) {
	user, err := u.UserRepo.AddUser(v)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
