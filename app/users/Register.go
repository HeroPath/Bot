package users

import (
	"aoweb-bot/app/models"
	"bytes"
	"encoding/json"
	"github.com/bxcodec/faker/v4"
	"math/rand"
	"net/http"
	"time"
)

func Register(ApiUrl string, users []models.User, amountUsers int) []models.User {
	for i := 0; i < amountUsers; i++ {
		var password = randomString(8)
		var userRegister models.UserRegister
		faker.FakeData(&userRegister)
		userRegister.Password = password
		data, _ := json.Marshal(userRegister)

		resp, _ := http.Post(ApiUrl+"auth/register", "application/json", bytes.NewBuffer(data))
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var user models.User
		decoder.Decode(&user)
		user.Token = "none"
		user.Password = password
		users = append(users, user)
	}
	return users
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var result string
	for i := 0; i < length; i++ {
		result += string(letters[rand.Intn(len(letters))])
	}

	return result
}
