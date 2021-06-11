package main

func get_skill(id string) [4]string {
	str := "select skill1, skill2, skill3, skill4 from pokemon_skill where pokemon_id= '" + id + "';"
	rows := db.QueryRow(str)
	var skills [4]string
	err = rows.Scan(&skills[0], &skills[1], &skills[2], &skills[3])
	checkErr(err)
	return skills
}

/*
func get_skill_info(skill_name string) Skill_info {
	str := "select skill_name, skill_type from skill where skill_name = '" + skill_name + "';"
	var info Skill_info
	var nil sql.NullString
	rows := db.QueryRow(str)
	err := rows.Scan(&info.skill_name, &info.skill_type)
	checkErr(err)
	str = "select* from skill where skill_name = '" + skill_name + "';"
	db.QueryRow(str)
	if info.skill_type == 0 { // 说明是攻击型金技能
		err = rows.Scan(&info.skill_name, &info.skill_type, &info.attack_type,
			&info.attack_power, info.self_damage, &info.attribute, &nil, &nil)
	} else if info.skill_type == 1 {
		err = rows.Scan(&info.skill_name, &info.skill_type, &nil, &nil,
			&nil, &nil, &info.effect_id, &info.effect_result)
	}
	return info
}*/
