package main


func main(){
	name:="test7"
	sentence2:="ALTER TABLE " + name + "_warehouse\n	OWNER to " + name + ";"
	println(sentence2);
}
