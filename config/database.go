package config

import (
	"log"
	"github.com/Zmohamed6991/JWT-Project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectingDB(){	

	db, err := gorm.Open(postgres.Open(`postgres://postgres:password@localhost:5432/JWT`)) //username, password, db name
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil{
		log.Fatal("error migrating to db")
	}

	DB = db

}


