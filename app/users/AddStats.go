package users

import (
	"aoweb-bot/app/models"
	"bytes"
	"net/http"
)

func AddStats(ApiUrl string, usersList []models.User, assortment bool) {
	for _, user := range usersList {
		if user.FreeSkillPoints > 0 {
			var mainStat string
			secondaryStat := "vitality"
			tertiaryStat := "luck"
			switch user.AClass {
			case "mage":
				mainStat = "intelligence"
			case "warrior":
				mainStat = "strength"
			case "archer":
				mainStat = "dexterity"
			}

			req, _ := http.NewRequest("POST", ApiUrl+"users/add-skill-points/"+mainStat, bytes.NewBuffer([]byte("")))
			req.Header.Set("Authorization", "Bearer "+user.Token)
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}

			if !assortment {
				for i := 0; i < user.FreeSkillPoints; i++ {
					resp, _ := client.Do(req)
					defer resp.Body.Close()
				}
			} else {
				for i := 0; i < user.FreeSkillPoints; i++ {
					if i%3 == 0 {
						req, _ = http.NewRequest("POST", ApiUrl+"users/add-skill-points/"+mainStat, bytes.NewBuffer([]byte("")))
						req.Header.Set("Authorization", "Bearer "+user.Token)
						req.Header.Set("Content-Type", "application/json")
					} else if i%3 == 1 {
						req, _ = http.NewRequest("POST", ApiUrl+"users/add-skill-points/"+secondaryStat, bytes.NewBuffer([]byte("")))
						req.Header.Set("Authorization", "Bearer "+user.Token)
						req.Header.Set("Content-Type", "application/json")
					} else {
						req, _ = http.NewRequest("POST", ApiUrl+"users/add-skill-points/"+tertiaryStat, bytes.NewBuffer([]byte("")))
						req.Header.Set("Authorization", "Bearer "+user.Token)
						req.Header.Set("Content-Type", "application/json")
					}
					resp, _ := client.Do(req)
					defer resp.Body.Close()
				}
			}
		}
	}
}
