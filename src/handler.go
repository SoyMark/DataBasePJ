package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func select_handler(c *gin.Context){
	var info pokemon_info
	c.ShouldBindJSON(&info)
	fmt.Println(info)
	if info.Id == ""{info.Id = "0"}
	if info.Name == ""{info.Name = "0"}
	if info.Property1 == ""{info.Property1 = "0"}
	if info.Property2 == ""{info.Property2 = "0"}
	infos, i :=pokemon_select(info)
	var my_slice []pokemon_info
	for j:=0;j<i;j++{
		my_slice = append(my_slice, infos[j])
	}
	c.IndentedJSON(200, my_slice)
}

func insert_handler(c *gin.Context){
	var info pokemon_info
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

func update_handler(c *gin.Context){
	var info pokemon_info
	c.ShouldBindJSON(&info)
	result := pokemon_update(info)
	if result == 1{
		c.String(http.StatusOK,"update or insert successfully!")
	}else{
		c.String(http.StatusOK,"fail to update or insert")
	}
}

func delete_handler(c *gin.Context){
	var info pokemon_info
	c.ShouldBindJSON(&info)
	result := pokemon_delete(info.Id)
	if result == 1{
		c.String(http.StatusOK,"delete successfully!")
	}else{
		c.String(http.StatusOK,"fail to delete")
	}
}

func create_role_handler(c *gin.Context){
	var user user_info
	c.ShouldBindJSON(&user)
	result :=create_role(user)
	if result == 1{
		c.String(http.StatusOK,"create a role successfully")
	}else{
		c.String(http.StatusOK,"fail to create a role")
	}
}

func login_handler(c *gin.Context){
	var user struct_login
	c.ShouldBindJSON(&user)
	fmt.Println(user)
	err = login(user, "DataBasePJ")
	if err!=nil{
		c.String(http.StatusOK,"登录失败!\n 错误信息：%s", err)
	}else {
		c.String(http.StatusOK, "登录成功")
	}//错误检查有点问题，但是登录是能用的
}

func user_free_handler(c *gin.Context){
	var free user_free
	c.ShouldBindJSON(&free)
	user_pokemon_delete(free.User_name, free.Pokemon_id)
}
