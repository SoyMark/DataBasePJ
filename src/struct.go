package main

type struct_login struct{
	Id string `json:"id"`
	Password string `json:"password"`
}

type pokemon_info struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
	Attack int `json:"attack"`
	Special_Attack int `json:"special_attack"`
	Defense int `json:"defense"`
	Special_Defense int `json:"special_defense""`
	Health int `json:"health"`
	Speed int `json:"speed"`
}

type user_info struct{
	Id string `jspn:"id"`
	Name string	`json:"name" binging:"required"`
	Password string	`json:"password"`
	Login int	`json:"login"`
	Superuser int	`json:"superuser"`
	Createdb int	`json:"createdb"`
	Createrole int	`json:"createrole"`
	Inherit int	`json:"inherit"`
	Replication int `json:"replication"`
}

type skill_info struct{
	skill_name string
	attribute string
	skill_type int
	attack_type string
	attack_power int
	effect_id string
	effect_result int
	self_damage int
}

type user_free struct{
	User_name string `json:"user_name"`
	Pokemon_id string `json:"pokemon_id"`
}
