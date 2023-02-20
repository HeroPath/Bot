package main

import (
	"aoweb-bot/app/models"
	"aoweb-bot/app/users"
	"time"
)

func main() {
	var usersList []models.User
	usersList = users.Register(usersList, 3)
	time.Sleep(3 * time.Second)
	usersList = users.Login(usersList)

}
