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
	stmt,err := db.Prepare("delete  from info where id = ?")
	res,err := stmt.Exec(5)
	id,err := res.RowsAffected()
	if err != nil{
		panic(err)
	}
	fmt.Println(id)

}
