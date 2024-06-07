package main
import (
	"database/sql"
	"fmt"
	"log"
	"sql_package"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Connecting into the database")
	sql_package.ConnectDatabase()
	fmt.Println("Creating the table in the database in the sql server")
	new_db ,err := sql.Open("mysql","root:Harsha@2004@tcp(127.0.0.1:3306)/Vardhan")
	if err!=nil {
		log.Fatal(err)
		return 
	}
	//sql_package.CreateTable(new_db)
	fmt.Println("Created the table successfully")
	sql_package.InsertIntoTable(new_db,1,"Harsha","harsha10012004@gmail.com")
	sql_package.InsertIntoTable(new_db,2,"Vardhan","vardhan123@gmail.com")
	sql_package.InsertIntoTable(new_db,3,"Ishaa","naim@123.gmail.com")
	fmt.Println("Inserted into the table succesfully")
	sql_package.UpdateDatabase(new_db,"Vardhan",1)
	fmt.Println("Updated the table in the database succesfully")
	sql_package.DeleteTable(new_db,1)
	fmt.Println("Deleted the table successfully")
	fmt.Println("Selecting the rows from the table in the database")
	sql_package.SelectData(new_db)

}
