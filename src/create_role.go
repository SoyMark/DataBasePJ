package main

import "fmt"


func create_role(user user_info) int {
	str:="CREATE ROLE \"" + user.Name + "\" with\n"
	if user.Login!=0{
		str+="login\n"
	} else{
		str+="nologin\n"
	}
	if user.Superuser!=0{
		str+="superuser\n"
	} else{
		str+="nosuperuser\n"
	}
	if user.Createdb!=0{
		str+="createdb\n"
	} else{
		str+="nocreatedb\n"
	}
	if user.Createrole!=0{
		str+="createrole\n"
	} else{
		str+="nocreaterole\n"
	}
	if user.Inherit!=0{
		str+="inherit\n"
	} else{
		str+="noinherit\n"
	}
	if user.Replication!=0{
		str+="replication\n"
	} else{
		str+="noreplication\n"
	}
	str+="connection limit -1\n"
	str+="password '" + user.Password + "';"
	rows, err:= db.Query(str)
	if rows == nil {} // rows没用到会报错，所以加一行
	checkErr(err)
	if err==nil {
		//在创建用户的时候为用户新建一个仓库
		sentence:= "create table " + user.Name + "_warehouse (\n"
		sentence += "pokemon_id character varying(20),\n"
		sentence += "pokemon_name character varying(20),"
		sentence += "primary key(pokemon_id),\n"
		sentence += "foreign key(pokemon_id) references pokemon(pokemon_id)\n );"
		rows, err = db.Query(sentence)
		checkErr(err)
		//在用户信息中加入新用户
		sentence1:= "insert into user_info values('" + user.Name + "', now(), 0, '"+user.Id +"');"
		rows, err = db.Query(sentence1)
		checkErr(err)
		//设置表的所有者为用户
		sentence2:="ALTER TABLE " + user.Name + "_warehouse\n	OWNER to " + user.Name + ";"
		rows, err := db.Query(sentence2)
		if rows==nil {}
		checkErr(err)
		fmt.Println(" create a role successfully")
		return 1
	} else{
		fmt.Println("fail to create a role")
		return 0
	}
}

func make_user_info(name string, password string, login int, superuser int, createdb int, createrole int, inherit int, replication int) user_info {
	var a user_info
	a.Name = name
	a.Login = login
	a.Superuser = superuser
	a.Createdb = createdb
	a.Createrole = createrole
	a.Inherit = inherit
	a.Replication = replication
	a.Password = password
	return a
}

/*func main(){
	a:=make_user_info("test2", "111", 1,0,0,0,0,0)
	str:=create_role(a)
	fmt.Println(str)
}*/
