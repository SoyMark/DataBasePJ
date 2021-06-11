package main

import "fmt"

func create_role(user User_info) error {
	str := "CREATE ROLE \"" + user.Id + "\" with\n"
	if user.Login != 0 {
		str += "login\n"
	} else {
		str += "nologin\n"
	}
	if user.Superuser != 0 {
		str += "superuser\n"
	} else {
		str += "nosuperuser\n"
	}
	if user.Createdb != 0 {
		str += "createdb\n"
	} else {
		str += "nocreatedb\n"
	}
	if user.Createrole != 0 {
		str += "createrole\n"
	} else {
		str += "nocreaterole\n"
	}
	if user.Inherit != 0 {
		str += "inherit\n"
	} else {
		str += "noinherit\n"
	}
	if user.Replication != 0 {
		str += "replication\n"
	} else {
		str += "noreplication\n"
	}
	str += "connection limit -1\n"
	str += "password '" + user.Password + "';"
	rows, err := db.Query(str)
	if rows == nil {
	} // rows没用到会报错，所以加一行
	checkErr(err)
	if err == nil {
		//在创建用户的时候为用户新建一个仓库
		sentence := "create table " + user.Id + "_warehouse (\n"
		sentence += "pokemon_id character varying(20),\n"
		sentence += "pokemon_name character varying(20),\n"
		sentence += "catch_time timestamp,\n"
		sentence += "primary key(pokemon_id,catch_time),\n"
		sentence += "foreign key(pokemon_id) references pokemon(pokemon_id)\n );"
		_, err = db.Query(sentence)
		//创建用户的backpack
		sentence3 := "create table " + user.Id + "_backpack (\n"
		sentence3 += "pokemon_id varchar(20),\n catch_time timestamp,\n " +
			"primary key(pokemon_id,catch_time)," +
			"\n\tforeign key(pokemon_id,catch_time) references " + user.Id + "_warehouse(pokemon_id,catch_time));"
		_, err = db.Query(sentence3)
		//创建用户的ballpack
		sentence4 := "create table " + user.Id + "_ballpack (\n"
		sentence4 += "ball_type int,\n\tball_num int,\n\tprimary key(ball_type));"
		_, err = db.Query(sentence4)
		//在ballpack里加上四种精灵球条目
		sentence5 := "insert into " + user.Id + "_ballpack values(1, 0);\n"
		sentence5 += "insert into " + user.Id + "_ballpack values(2, 0);\n"
		sentence5 += "insert into " + user.Id + "_ballpack values(3, 0);\n"
		sentence5 += "insert into " + user.Id + "_ballpack values(4, 0);"
		_, err = db.Query(sentence5)
		//在用户信息中加入新用户
		sentence1 := "insert into User_info values('" + user.Name + "', now(), 1000, '" + user.Id + "');\n"
		_, err = db.Query(sentence1)
		//设置两个背包和仓库表的所有者为用户
		sentence2 := "ALTER TABLE " + user.Id + "_warehouse\n	OWNER to " + user.Id + ";\n"
		sentence2 += "alter table " + user.Id + "_backpack\n owner to " + user.Id + ";\n"
		sentence2 += "alter table " + user.Id + "_ballpack\n owner to " + user.Id + ";"
		_, err := db.Query(sentence2)

		fmt.Println(" create a role successfully")
		return err
	} else {
		fmt.Println("fail to create a role")
		return err
	}

}

func make_user_info(id string, name string, password string, login int, superuser int, createdb int, createrole int, inherit int, replication int) User_info {
	var a User_info
	a.Id = id
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
