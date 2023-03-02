package users

import (
	"Bot/app/models"
	"bytes"
	"encoding/json"
	"net/http"
)

func GetProfile(ApiUrl string, users []models.User) []models.User {
	for i := 0; i < len(users); i++ {
		token := users[i].Token
		req, _ := http.NewRequest("GET", ApiUrl+"users/profile", bytes.NewBuffer([]byte("")))
		req.Header.Set("Authorization", "Bearer "+users[i].Token)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var user models.User
		decoder.Decode(&user)
		users[i] = user
		users[i].Token = token
	}
	return users
}
