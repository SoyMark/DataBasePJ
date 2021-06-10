package main

import (
	"fmt"
	"strconv"
)

func buy(info User_price) error {
	str:="update user_info set user_money = user_money - " + strconv.Itoa(info.Price)
	str+=" where user_id = '" + info.User_id + "';"
	fmt.Println(str)
	rows,err2 := db.Query(str)
	if rows == nil {}
	return err2
}
