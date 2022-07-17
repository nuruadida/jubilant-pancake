package usecase

import "Nurul-trial/models"

type UcaseLoginInterface interface {
	LoginUser(username string, password string, email string) (models.GetItGuys, error)
}

type UcaseUserInterface interface {
	AddUser(v models.User) (models.User, error)
}
