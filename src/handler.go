package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Select_handler(c *gin.Context) {
	var info Pokemon_info
	c.ShouldBindJSON(&info)
	fmt.Println(info)
	if info.Id == ""{info.Id = "0"}
	if info.Name == ""{info.Name = "0"}
	if info.Property1 == ""{info.Property1 = "0"}
	if info.Property2 == ""{info.Property2 = "0"}
	infos, i :=pokemon_select(info)
	var my_slice []Pokemon_info
	for j:=0;j<i;j++{
		my_slice = append(my_slice, infos[j])
	}
	c.IndentedJSON(200, my_slice)
}

func Insert_handler(c *gin.Context) {
	var info Pokemon_info
	c.ShouldBindJSON(&info)
	if info.Id == ""{info.Id = "0"}
	if info.Name == ""{info.Name = "0"}
	if info.Property1 == ""{info.Property1 = "0"}
	if info.Property2 == ""{info.Property2 = "0"}
	result := pokemon_insert(info)
	if result == 1{
		c.String(http.StatusOK,"insert successfully!")
	}else{
		c.String(http.StatusOK,"fail to delete")
	}
}

func Update_handler(c *gin.Context) {
	var info Pokemon_info
	c.ShouldBindJSON(&info)
	result := pokemon_delete(info.Id)
	if result == 1{
		c.String(http.StatusOK,"delete successfully!")
	}else{
		c.String(http.StatusNotAcceptable,"fail to delete")
	}
}

func Delete_handler(c *gin.Context) {
	var info Pokemon_info
	c.ShouldBindJSON(&info)
	result := pokemon_delete(info.Id)
	if result == 1 {
		c.String(http.StatusOK, "delete successfully!")
	} else {
		c.String(http.StatusOK, "fail to delete")
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

func Login_handler(c *gin.Context) {
	var user Struct_login
	c.ShouldBindJSON(&user)
	// c.Request.Response.Header.Set("Access-Control-Allow-Origin", "*")
	// c.Request.Response.Header.Add("Access-Control-Allow-Headers", "Context-Type")
	// c.Request.Response.Header.Set("content-type", "application/json")

	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Add("Access-Control-Allow-Headers", "Context-Type")
	// c.Writer.Header().Set("content-type", "application/json")
	fmt.Println(user)
	err = login(user, "DataBasePJ")
	if err!=nil{
		c.String(http.StatusNotAcceptable,"登录失败!\n 错误信息：%s", err)
	}else {
		c.String(http.StatusOK, "登录成功")
	}//错误检查有点问题，但是登录是能用的
}

func User_free_handler(c *gin.Context){
	var free User_free
	c.ShouldBindJSON(&free)
	user_pokemon_delete(free.User_name, free.Pokemon_id)
}

func load_handler(c *gin.Context){

}
