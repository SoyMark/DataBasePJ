package main


//用户在仓库里点击查看宝可梦的完整信息
func pokemon_details(id string) pokemon_info{
	var info pokemon_info
	rows, err := db.Query("select* from pokemon where pokemon_id = '" + id + "';")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&info.Id, &info.Name, &info.Property1,
			&info.Attack, &info.Special_Attack, &info.Defense,
			&info.Special_Defense, &info.Health,&info.Speed, &info.Property2)
	}
	return info
}