package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db,err := sql.Open("mysql","root:123123@tcp(127.0.0.1:3306)/demo?charset=utf8")
	if err !=nil {
		panic(err)
	}
	row,err := db.Query("select * from info where id = 3")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%v",row)
	for row.Next(){
		var id int
		var username string
		var address string
		var phone string
		err = row.Scan(&id,&username,&address,&phone)
		fmt.Println(id,username,address,phone)
	}

}
