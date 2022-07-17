package repository

import "Nurul-trial/models"

type RepoLoginInterface interface {
	LoginUser(username string, email string, password string) (models.GetItGuys, error)
}

type RepoUserInterface interface {
	AddUser(v models.User) (models.User, error)
}
