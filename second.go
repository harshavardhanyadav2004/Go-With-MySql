package sql_package 
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
 func CreateTable(db *sql.DB){
	 const file string = `
	 CREATE TABLE Students (
	       id INTEGER ,
		   name TEXT ,
		   email TEXT 
	 
	 )
	`
	if _,err:= db.Exec(file) ; err!= nil {
		 log.Fatal(err)
		 return 
	}
	fmt.Println("Created table successfully")

 }