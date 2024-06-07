package sql_package

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func UpdateDatabase(db *sql.DB, name string , id int){
	 fmt.Println("Updating the value in the database")
	 another_one , err:= db.Exec("UPDATE Students SET name = ? WHERE id = ?",name,id)
	 if err!=nil {
		 log.Fatal(err)
		 return 
	 }
	 effect,err:=another_one.RowsAffected()
	if err!=nil {
		log.Fatal(err)
		return
	}
	fmt.Println("The Affected rows are ",effect)
	


}