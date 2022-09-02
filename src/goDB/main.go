package main

import (
	"dbhandler/mysql"
	"fmt"
)

func main() {
	fmt.Println("Hello Mongo")
	mysql.Mysql_Start()
}
