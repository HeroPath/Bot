package main

import (
	"aoweb-bot/app/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var users []models.User

	userRegister := models.UserRegister{
		Username:  "test",
		Password:  "test",
		Email:     "test@test.com",
		ClassName: "mage",
	}

	data, _ := json.Marshal(userRegister)
	resp, _ := http.Post("http://localhost:8000/api/v1/auth/register", "application/json", bytes.NewBuffer(data))
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var user models.User
	decoder.Decode(&user)
	users = append(users, user)

	fmt.Printf("%+v\n", users)
}
