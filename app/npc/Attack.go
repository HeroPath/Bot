package npc

import (
	"aoweb-bot/app/models"
	"bytes"
	"encoding/json"
	"net/http"
)

func getListNpcs(ApiUrl string, token string) []models.Npc {
	var npcList []models.Npc
	
	req, _ := http.NewRequest("GET", ApiUrl+"npcs/zone/forest", bytes.NewBuffer([]byte("")))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&npcList)
	return npcList
}

func AttackNPC(ApiUrl string, users []models.User, initialNpc bool) {
	for i := 0; i < len(users); i++ {
		var npcList = getListNpcs(ApiUrl, users[i].Token)
		var data = map[string]interface{}{
			"name": "",
		}
		if initialNpc {
			data["name"] = npcList[0].Name
		} else {
			userLevel := users[i].Level
			var npc models.Npc
			for _, val := range npcList {
				if userLevel >= val.Level {
					npc = val
				}
			}
			data["name"] = npc.Name
		}
		jsonData, _ := json.Marshal(data)
		req, _ := http.NewRequest("POST", ApiUrl+"users/attack-npc", bytes.NewBuffer(jsonData))
		req.Header.Set("Authorization", "Bearer "+users[i].Token)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, _ := client.Do(req)
		defer resp.Body.Close()
	}
}
