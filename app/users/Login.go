package users

import (
	"aoweb-bot/app/models"
	"bytes"
	"encoding/json"
	"net/http"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(ApiUrl string, users []models.User) []models.User {
	for i, user := range users {
		if user.Token == "none" {
			var userLogin models.UserLogin
			userLogin.Username = user.Username
			userLogin.Password = user.Password
			data, _ := json.Marshal(userLogin)

			resp, _ := http.Post(ApiUrl+"auth/login", "application/json", bytes.NewBuffer(data))
			defer resp.Body.Close()

			decoder := json.NewDecoder(resp.Body)
			var loginResponse LoginResponse
			decoder.Decode(&loginResponse)

			users[i].Token = loginResponse.Token
		}
	}
	return users
}
