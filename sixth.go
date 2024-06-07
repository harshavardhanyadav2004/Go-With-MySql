package sql_package 
import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)
func SelectData(db *sql.DB){
	fmt.Println("Connecting into the database in the sql")
    rows , err := db.Query("SELECT * FROM Students")
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Extracting the rows from the table")
	for rows.Next(){
		var id int 
		var name string 
		var email string 
		rows.Scan(&id,&name,&email)
		fmt.Println("Name :",name,"Name:",name,"Email :",email)
	}
	if  err:= rows.Err() ; err!=nil {
		log.Fatal(err)
	}

}