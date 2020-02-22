// Pacakge provides the db object to be used by the other packages (fetchrss,searchress,serverss)
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	DbName     = "rssfeeds"
	DbUser     = "root"
	DbPassword = "admin"
)

func Conn() *sql.DB {
	creds := fmt.Sprintf("%v:%v@/%v", DbUser, DbPassword, DbName)
	conn, err := sql.Open("mysql", creds)
	if err != nil {
		fmt.Println("Error in db opening :>", err)
		log.Fatal(err)
	}
	return conn
}
