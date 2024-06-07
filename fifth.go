package sql_package
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func DeleteTable(db *sql.DB, id int){
	fmt.Println("Deleting the value in the database")
	another_one_id, err := db.Exec("DELETE FROM Students WHERE id = 1")
	if err!=nil {
		log.Fatal(err)
		return
	}
	new_id_table , err:= another_one_id.RowsAffected()
	if err!=nil {
		log.Fatal(err)
		return
	}
	fmt.Println("The rows affected in the table are ",new_id_table)
	
}