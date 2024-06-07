package sql_package

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func InsertIntoTable(db *sql.DB,new_id int , name string , email string ){
	fmt.Println("Inserting Into the table in the database")
	var query string = "INSERT INTO Students (id,name,email) VALUES (?,?,?)"
	id,err:=db.Exec(query,new_id,name,email)
	if err!=nil {
		log.Fatal(err)
		return 
	}
	more_id , err:=id.LastInsertId() 
	if err!=nil {
		log.Fatal(err)
		return 
	}
	fmt.Println("The Last inserted id is",more_id)

}