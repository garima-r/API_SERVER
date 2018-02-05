package dbconnection

import(
    "database/sql"
    _"database/sql/driver/mysql"
    
)

var Db *sql.DB 
var err error

func Init(){

	Db, err = sql.Open("mysql","root:system@tcp(127.0.0.1:3306)/godatabase")                //database connection
    
    if err != nil {
        panic(err.Error()) 
        defer Db.Close() 
    }
    
    err = Db.Ping()
    if err != nil {
        panic(err.Error())
    }
}