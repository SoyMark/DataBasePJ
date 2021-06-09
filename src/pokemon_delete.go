package main

import "fmt"


func pokemon_delete(id string) int{
	str:="delete from pokemon where pokemon_id = '" + id +"';"
	fmt.Println(str)
	rows, err := db.Query(str)
	if rows==nil {}
	checkErr(err)
	if err==nil {
		return 1
	}else{
		return 0
	}
}

func user_pokemon_delete(username string, id string){
	str:="delete from " + username + "_warehouse " + "where pokemon_id = '" + id + "';"
	rows, err := db.Query(str)
	checkErr(err)
	if err==nil {
		fmt.Println("delete successfully!")
	}else{
		fmt.Println("fail to delete")
	}
	if rows==nil {}
}
