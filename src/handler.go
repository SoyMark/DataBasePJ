package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//宝可梦图鉴，返回宝可梦所有信息，前端根据需要显示
func Select_handler(c *gin.Context) {
	var info Pokemon_info
	c.ShouldBindJSON(&info)
	fmt.Println(info)
	if info.Id == "" {
		info.Id = "0"
	}
	if info.Name == "" {
		info.Name = "0"
	}
	if info.Property1 == "" {
		info.Property1 = "0"
	}
	if info.Property2 == "" {
		info.Property2 = "0"
	}
	infos, i := pokemon_select(info)
	var my_slice []Pokemon_info
	for j := 0; j < i; j++ {
		my_slice = append(my_slice, infos[j])
	}
	c.IndentedJSON(200, my_slice)
}

func Insert_handler(c *gin.Context) {
	var info Pokemon_info
	c.ShouldBindJSON(&info)
	if info.Id == "" {
		info.Id = "0"
	}
	if info.Name == "" {
		info.Name = "0"
	}
	if info.Property1 == "" {
		info.Property1 = "0"
	}
	if info.Property2 == "" {
		info.Property2 = "0"
	}
	result := pokemon_insert(info)
	if result == 1 {
		c.String(http.StatusOK, "insert successfully!")
	} else {
		c.String(http.StatusNotAcceptable, "fail to delete")
	}
}

func Update_handler(c *gin.Context) {
	var info Pokemon_info
	c.ShouldBindJSON(&info)
	result := pokemon_delete(info.Id)
	if result == 1 {
		c.String(http.StatusOK, "delete successfully!")
	} else {
		c.String(http.StatusNotAcceptable, "fail to delete")
	}
}

func Delete_handler(c *gin.Context) {
	var info Pokemon_info
	c.ShouldBindJSON(&info)
	result := pokemon_delete(info.Id)
	if result == 1 {
		c.String(http.StatusOK, "delete successfully!")
	} else {
		c.String(http.StatusNotAcceptable, "fail to delete")
	}
}

func Create_role_handler(c *gin.Context) {
	var user User_info
	c.ShouldBindJSON(&user)
	err2 := create_role(user)
	if err2 == nil {
		c.String(http.StatusOK, "create a role successfully")
	} else {
		c.String(http.StatusNotAcceptable, "fail to create a role")
	}
}

//登录
func Login_handler(c *gin.Context) {
	c.ShouldBindJSON(&user)
	fmt.Println(user)
	_, err := login(user, "DataBasePJ")
	if err != nil {
		c.String(http.StatusNotAcceptable, "登录失败")
	} else {
		rows, err := db.Query("select user_name, user_money from user_info where user_id = '" + user.Id + "';")
		var name string
		var money int
		for rows.Next() {
			rows.Scan(&name, &money)
		}
		rows2, err := db.Query("select count(*) from " + user.Id + "_warehouse;")
		var num int
		for rows2.Next() {
			rows2.Scan(&num)
		}
		var poke_id [6]string
		i := 0
		rows3, err := db.Query("select pokemon_id from " + user.Id + "_backpack;")
		for rows3.Next() {
			rows3.Scan(&poke_id[i])
			i++
		}
		checkErr(err)
		var result Result
		var poke_info Pokemon_info_withskill
		for j := 0; j < i; j++ {
			poke_info.info = pokemon_details(poke_id[j])
			poke_info.skill = get_skills_info(poke_id[j])
			result.pokemons = append(result.pokemons, poke_info)
		}
		var s string
		//fmt.Sprintf(s,result)
		fmt.Println(result)
		//data, _ := json.Marshal(result)
		c.IndentedJSON(http.StatusOK, result)
		c.String(http.StatusOK, s)
		c.JSON(http.StatusOK, gin.H{
			"id":          user.Id,
			"name":        name,
			"pokemon_num": num,
			"money":       money,
		})
	}
}

//放生
func Poke_free_handler(c *gin.Context) {
	var free Id_time
	c.ShouldBindJSON(&free)
	err := user_pokemon_delete(user.Id, free.Id, free.Catch_time)
	if err == nil {
		c.String(http.StatusOK, "放生成功")
	} else {
		c.String(http.StatusNotAcceptable, "放生失败")
	}
}

//返回用户仓库里的信息
type warehouse_struct struct {
	Pokemon_name string
	Property     string
	Catch_time   string
}

func Warehouse_handler(c *gin.Context) {
	user_warehosue := user.Id + "_warehouse"
	str := "select pokemon_name, catch_time, property from " + user_warehosue + " natural join pokemon"
	rows, err := db.Query(str)
	i := 0
	var temp warehouse_struct
	var my_slice []warehouse_struct
	for rows.Next() {
		rows.Scan(&temp.Pokemon_name, &temp.Catch_time, &temp.Property)
		i++
		my_slice = append(my_slice, temp)
	}
	if err == nil {
		c.IndentedJSON(http.StatusOK, my_slice)
	} else {
		c.JSON(http.StatusNotAcceptable, "{}")
	}
}

//放入背包
type Id_time struct {
	Id         string `json:"id"`
	Catch_time string `json:"catch_time"`
}

func Takein_handler(c *gin.Context) {
	var info Id_time
	c.ShouldBindJSON(&info)
	user_backpack := user.Id + "_backpack"
	str := "insert into " + user_backpack + " values('" + info.Id + "','" + info.Catch_time + "');"
	_, err := db.Query(str)
	if err == nil {
		c.String(http.StatusOK, "放入成功")
	} else {
		c.String(http.StatusNotAcceptable, "放入失败")
	}
}

//从背包取出
func Takeout_handler(c *gin.Context) {
	var info Id_time
	c.ShouldBindJSON(&info)
	user_backpack := user.Id + "_backpack"
	str := "delete from " + user_backpack + " where pokemon_id = '" + info.Id +
		"' and catch_time = '" + info.Catch_time + "';"
	_, err := db.Query(str)
	if err == nil {
		c.String(http.StatusOK, "取出成功")
	} else {
		c.String(http.StatusNotAcceptable, "取出失败")
	}
}

type Buy struct {
	Money int `json:"money"`
	Type  int `json:"type"`
}

func Buy_pokeball_handler(c *gin.Context) {
	var info Buy
	c.ShouldBindJSON(&info)
	str := "update user_info set user_money = user_money - " + strconv.Itoa(info.Money) +
		" where user_id = '" + user.Id + "';"
	fmt.Println(str)
	_, err := db.Query(str)
	if err == nil {
		user_pack := user.Id + "_ballpack"
		str = "update " + user_pack + " set ball_num = ball_num + 1 where ball_type = " + strconv.Itoa(info.Type) + ";"
		fmt.Println(str)
		db.Query(str)
		c.String(http.StatusOK, "购买成功")
	} else {
		c.String(http.StatusNotAcceptable, "无法购买！")
	}
}

//抓宝可梦成功放入仓库
func Catch_handler(c *gin.Context) {
	var info Id_time
	var name string
	c.ShouldBindJSON(&info)
	user_warehouse := user.Id + "_warehouse"
	rows := db.QueryRow("select pokemon_name from pokemon where pokemon_id = '" + info.Id + "';")
	rows.Scan(&name)
	str := "insert into " + user_warehouse + " values('" + info.Id + "', '" + name + "', '" +
		info.Catch_time + "');"
	_, err := db.Query(str)
	if err == nil {
		c.String(http.StatusOK, "捕捉成功")
	} else {
		c.String(http.StatusNotAcceptable, "出现问题，无法放入仓库")
	}
}

//奖励
type Award struct {
	Money int `json:"money"`
}

func Award_handler(c *gin.Context) {
	var info Award
	c.ShouldBindJSON(&info)
	str := "update user_info set user_money = user_money + " + strconv.Itoa(info.Money) +
		" where user_id = '" + user.Id + "';"
	_, err := db.Query(str)
	if err == nil {
		c.String(http.StatusOK, "")
	} else {
		c.String(http.StatusNotAcceptable, "")
	}
}

type A_poke_info struct {
	Pokemon_id string `json:"pokemon_id"`
}

func Fight_handler(c *gin.Context) {
	var info A_poke_info
	c.BindJSON(&info)
	fmt.Println(info.Pokemon_id)
	var poke_info Pokemon_info_withskill
	poke_info.info = pokemon_details(info.Pokemon_id)
	poke_info.skill = get_skills_info(info.Pokemon_id)
	fmt.Println(poke_info)
	c.IndentedJSON(http.StatusOK, poke_info)
}
