package models

type Npc struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Hp      int    `json:"hp"`
	MaxHp   int    `json:"maxHp"`
	MinDmg  int    `json:"minDmg"`
	MaxDmg  int    `json:"maxDmg"`
	Defense int    `json:"defense"`
	Zone    string `json:"zone"`
}
