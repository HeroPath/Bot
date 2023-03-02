package main

import (
	"Bot/app/models"
	"Bot/app/npc"
	"Bot/app/users"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ApiUrl := os.Getenv("API_URL")
	amountUsers, err := strconv.Atoi(os.Getenv("AMOUNT_USERS_TO_REGISTER"))
	addInitialStatsRandom, err := strconv.ParseBool(os.Getenv("ADD_INITIAL_STATS_RANDOM"))
	attackNpcs, err := strconv.ParseBool(os.Getenv("ATTACK_NPCS"))

	var usersList []models.User
	usersList = users.Register(ApiUrl, usersList, amountUsers)
	time.Sleep(3 * time.Second)
	usersList = users.Login(ApiUrl, usersList)
	users.AddStats(ApiUrl, usersList, addInitialStatsRandom)

	userActing(ApiUrl, usersList, attackNpcs, 500)

}

func userActing(ApiUrl string, usersList []models.User, attackNpc bool, repetitions int) {
	if attackNpc {
		npc.AttackNPC(ApiUrl, usersList, true)
		usersList = users.GetProfile(ApiUrl, usersList)
		users.AddStats(ApiUrl, usersList, false)

		for i := 0; i < repetitions; i++ {
			npc.AttackNPC(ApiUrl, usersList, false)
			usersList = users.GetProfile(ApiUrl, usersList)
		}
	}

}
