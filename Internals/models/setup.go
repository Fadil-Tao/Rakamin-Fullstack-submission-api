package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB 

func ConnectDatabase(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("DBPORT")
	DBHOST := os.Getenv("DBHOST")
	USER := os.Getenv("USER")	
	PASSWORD := os.Getenv("PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DBHOST, USER, PASSWORD, DB_NAME, PORT)


	database , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        log.Fatalln(err,"connect failed")
    }
	
	database.AutoMigrate(&User{},&Photos{})
	
	DB = database
}