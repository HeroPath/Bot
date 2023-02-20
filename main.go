package main

import (
	"aoweb-bot/app/models"
	"aoweb-bot/app/users"
	"time"
)

const ApiUrl = "http://localhost:8000/api/v1/"

func main() {
	var usersList []models.User
	usersList = users.Register(ApiUrl, usersList, 10)
	time.Sleep(3 * time.Second)
	usersList = users.Login(ApiUrl, usersList)
	users.AddStats(ApiUrl, usersList, false)
}
