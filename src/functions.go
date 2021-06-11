package main

//用户在仓库里点击查看宝可梦的完整信息
func pokemon_details(id string) Pokemon_info {
	var info Pokemon_info
	rows, err := db.Query("select* from pokemon where pokemon_id = '" + id + "';")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&info.Id, &info.Name, &info.Property1,
			&info.Attack, &info.Special_Attack, &info.Defense,
			&info.Special_Defense, &info.Health, &info.Speed, &info.Property2)
	}
	return info
}

func get_skills_info(id string) [4]Skill_info {
	var s1 string
	var s2 string
	var s3 string
	var s4 string
	str := "select* from pokemon_skill where pokemon_id = '" + id + "';"
	rows, _ := db.Query(str)
	var useless string
	for rows.Next() {
		rows.Scan(&useless, &s1, &s2, &s3, &s4)
	}
	var ret [4]Skill_info
	ret[0] = get_skills_details(s1)
	ret[1] = get_skills_details(s2)
	ret[2] = get_skills_details(s3)
	ret[3] = get_skills_details(s4)
	return ret
}

func get_skills_details(name string) Skill_info {
	rows, _ := db.Query("select* from skill where skill_name = '" + name + "';")
	var s Skill_info
	s.Effect_id = make([]int, 5)
	s.Effect_result = make([]int, 5)
	for rows.Next() {
		rows.Scan(&s.Skill_name, &s.Skill_type, &s.Attack_type, &s.Attack_power,
			&s.Self_damage, &s.Attribute, &s.Effect_id, &s.Effect_result)
	}
	return s
}
