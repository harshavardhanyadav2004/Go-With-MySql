package sql_package

import(
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func ConnectDatabase(){
	 database ,err:=sql.Open("mysql","root:Harsha@2004@tcp(127.0.0.1:3306)/Vardhan")
	 if err!=nil {
		log.Fatal(err)
		return 
	 }
	 if err:=database.Ping() ; err!=nil {
		 log.Fatal(err)
		 return
	 }
	 fmt.Println("Connected to the database succesfully")
	 defer database.Close()
	 fmt.Println("The DataBase was closed successfully")
	 
	 
}