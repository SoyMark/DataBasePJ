package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	// "strings"
	// "net/http"
)

var db *sql.DB = nil
var err error

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func temp_select(id string) string {
	rows, err := db.Query("select pokemon_name from pokemon where pokemon_id = '" + id + "';")
	checkErr(err)
	var name *string
	name = new(string)
	for rows.Next() {
		err = rows.Scan(&name)
	}
	defer rows.Close()
	return *name
}

func login(user Struct_login, dbname string) error{
	var str string
	str = "user=" + user.Id + " password=" + user.Password + " dbname=" + dbname + " sslmode=disable"
	db, err = sql.Open("postgres", str)
	err = db.Ping()
	return err
	/*if err==nil{
		return 1
	}else {
		return 0
	}*/
}

func handler(c *gin.Context) {
	id := c.PostForm("pokemon_id")
	fmt.Println(id)
	name := temp_select(id)
	c.JSON(200, gin.H{
		"pokemon_name": name,
	})
}

func logout(db *sql.DB) int {
	db.Close()
	return 0
}

	
func main() {
	/*var user struct_login
	user.Name = "postgres"
	user.Password = "1908"
	db = login(user, "DataBasePJ")*/
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/select", Select_handler)
	router.POST("/insert", Insert_handler)
	router.POST("/delete", Delete_handler)
	router.POST("/update", Update_handler)
	router.POST("/create_role", Create_role_handler)
	router.POST("/login", Login_handler)
	router.POST("/user_free",User_free_handler)
	router.Run()
	/*var user2 struct_login
	user2.name = "syl"
	user2.password = "123"
	db2:=login(user2, "DataBasePJ")
	rows,err:=db1.Query("insert into store values('7',1,1);")
	if rows==nil {}
	checkErr(err)
	rows,err=db2.Query("insert into store values('8',2,2);")
	checkErr(err)*/
	//pokemon_search("select* from pokemon where pokemon_id = '123'")
	//info := make_pokemon_info("1215", "胡地", "超能力","1", 1,1,1,1,1,1)
	//new2:= make_pokemon_info("1234", " ", " ", " ",1,1,1,1,1,1)
	//pokemon_update(info)
	//a:=make_user_info("test5", "111", 1,0,1,1,0,0)
	/*info:=make_pokemon_info("0", "0", "0", "0", 0, 0, 0,0, 60, 0)
	sentense:=select_make(info)
	pokemon_search(sentense)*/
	//create_role(a)
	//skills:=get_skill("123")
	//info1:=get_skill_info(skills[0])
	//fmt.Println(info1)
	//db.Close()
}
