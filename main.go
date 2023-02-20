package main

import (
	"aoweb-bot/app/models"
	"aoweb-bot/app/users"
	"time"
)

func main() {
	var usersList []models.User
	usersList = users.Register(usersList, 10)
	time.Sleep(3 * time.Second)

}
