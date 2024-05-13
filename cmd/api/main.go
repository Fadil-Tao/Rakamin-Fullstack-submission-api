package main

import "github.com/rakamins-pbi/final-task-pbi-rakamin-fullstack-HadadFadilah/Internals/models"

func main() {
	router := Router()
	models.ConnectDatabase()
	router.Run(":8080")
}