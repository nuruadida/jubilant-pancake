package models

import "gorm.io/gorm"

type GetItGuys struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
}
