package main

import "fmt"

func edit_role(user user_info) string{
	str:="alter role \"" + user.name + "\"\n"
	if user.login!=0{
		str+="login\n"
	} else{
		str+="nologin\n"
	}
	if user.superuser!=0{
		str+="superuser\n"
	} else{
		str+="nosuperuser\n"
	}
	if user.createdb!=0{
		str+="createdb\n"
	} else{
		str+="nocreatedb\n"
	}
	if user.createrole!=0{
		str+="createrole\n"
	} else{
		str+="nocreaterole\n"
	}
	if user.inherit!=0{
		str+="inherit\n"
	} else{
		str+="noinherit\n"
	}
	if user.replication!=0{
		str+="replication\n"
	} else{
		str+="noreplication\n"
	}
	str+="connection limit -1\n"
	str+="password '" + user.password + "';"
	rows, err:= db.Query(str)
	if rows == nil {} // rows没用到会报错，所以加一行
	checkErr(err)
	if err==nil {
		return "edit a role successfully"
	}
	return "fail to edit a role"
}

//删除一个用户，先删除他的宝可梦库（表），再删除用户
func delete_role(name string){
	str_table:= "drop table " + name + "_warehouse;\n"
	str_user:= "drop user " + name + ";"
	str_table+=str_user
	rows, err:=db.Query(str_table)
	checkErr(err)
	if rows==nil {}
	if err == nil {
		fmt.Println("delete successfully!")
	}else{
		fmt.Println("fail to delete a user")
	}
}