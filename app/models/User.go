package models

type User struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Inventory struct {
		Items []interface{} `json:"items"`
	} `json:"inventory"`
	Equipment struct {
		Items []interface{} `json:"items"`
	} `json:"equipment"`
	Level                 int     `json:"level"`
	Experience            int     `json:"experience"`
	ExperienceToNextLevel int     `json:"experienceToNextLevel"`
	Gold                  int     `json:"gold"`
	Diamond               int     `json:"diamond"`
	MaxDmg                int     `json:"maxDmg"`
	MinDmg                int     `json:"minDmg"`
	MaxHp                 int     `json:"maxHp"`
	Hp                    int     `json:"hp"`
	Defense               int     `json:"defense"`
	Evasion               int     `json:"evasion"`
	CriticalChance        float64 `json:"criticalChance"`
	Strength              int     `json:"strength"`
	Dexterity             int     `json:"dexterity"`
	Intelligence          int     `json:"intelligence"`
	Vitality              int     `json:"vitality"`
	Luck                  int     `json:"luck"`
	FreeSkillPoints       int     `json:"freeSkillPoints"`
	NpcKills              int     `json:"npcKills"`
	PvpWins               int     `json:"pvpWins"`
	PvpLosses             int     `json:"pvpLosses"`
	TitleName             string  `json:"titleName"`
	TitlePoints           int     `json:"titlePoints"`
	GuildName             string  `json:"guildName"`
	AClass                string  `json:"aclass"`
}

type UserRegister struct {
	Username  string `json:"username" faker:"username" faker_len:"10"`
	Password  string `json:"password" faker:"password" faker_len:"10"`
	Email     string `json:"email" faker:"email"`
	ClassName string `json:"className" faker:"oneof:mage,warrior,archer"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
