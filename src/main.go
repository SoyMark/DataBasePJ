package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB = nil
var err error

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func login(user struct_login, dbname string) error{
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


func logout(db *sql.DB) int{
	db.Close()
	return 0
}

func main(){
	/*var user struct_login
	user.Name = "postgres"
	user.Password = "1907"
	login(user, "DataBasePJ")*/
	router := gin.Default()
	router.POST("/select",select_handler)
	router.POST("/insert",insert_handler)
	router.POST("/delete",delete_handler)
	router.POST("/update",update_handler)
	router.POST("/create_role",create_role_handler)
	router.POST("/login",login_handler)
	router.POST("/user_free",user_free_handler)
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