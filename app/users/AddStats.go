package users

import (
	"aoweb-bot/app/models"
	"bytes"
	"net/http"
)

func AddStats(ApiUrl string, usersList []models.User) {
	for _, user := range usersList {
		if user.FreeSkillPoints > 0 {
			var urlAclass string
			switch user.AClass {
			case "mage":
				urlAclass = "intelligence"
			case "warrior":
				urlAclass = "strength"
			case "archer":
				urlAclass = "dexterity"
			}

			req, _ := http.NewRequest("POST", ApiUrl+"users/add-skill-points/"+urlAclass, bytes.NewBuffer([]byte("")))
			req.Header.Set("Authorization", "Bearer "+user.Token)
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			for i := 0; i < user.FreeSkillPoints; i++ {
				resp, _ := client.Do(req)
				defer resp.Body.Close()
			}
		}
	}
}
