package main

import (
	"fmt"
	"strconv"
)

func pokemon_insert_make(info pokemon_info) string{
	str:="insert into pokemon values("
	str+="'" + info.Id + "', "
	str+="'" + info.Name + "', "
	str+="'" + info.Property1 + "', "
	str+=strconv.Itoa(info.Attack)+", "
	str+=strconv.Itoa(info.Special_Attack)+", "
	str+=strconv.Itoa(info.Defense)+", "
	str+=strconv.Itoa(info.Special_Defense)+", "
	str+=strconv.Itoa(info.Health)+", "
	str+=strconv.Itoa(info.Speed)
	if info.Property2 != "0" {
		str+=", '" + info.Property2 + "'"
	}
	str+=");"
	return str
}

func pokemon_insert(info pokemon_info) int{
	str:=pokemon_insert_make(info)
	rows,err:=db.Query(str)
	checkErr(err)
	if rows == nil {} //防止rows没用到报错
	if err==nil {
		return 1
	}else{
		return 0
	}
}

func user_pokemon_insert(username string, id string){
	rows, err := db.Query("select pokemon_name from pokemon where pokemon_id = '" + id + "';")
	checkErr(err)
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
	}
	str:="insert into " + username + "_warehouse "
	str+="values('" + id + "', '" + name + "');"
	fmt.Println(str)
	rows, err = db.Query(str)
	checkErr(err)
	if err==nil {
		fmt.Println("insert successfully!")
	}else{
		fmt.Println("fail to insert")
	}
	if rows==nil {}
}
