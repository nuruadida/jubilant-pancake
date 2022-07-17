package db

import (
	"Nurul-trial/config/env"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	var configDb = env.Config

	if err != nil {
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", configDb.Host, configDb.User, configDb.Pass, configDb.Name, configDb.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to DB")
	}

	DB = db
	fmt.Println("success connect to DB")
	Automigrate(db)

}
