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

// Create and return a database connection for our DB
func Conn() *sql.DB {
	creds := fmt.Sprintf("%v:%v@tcp(RSSProcessorDockerDB)/%v", DbUser, DbPassword, DbName) // Include name of container for tcp i.e @tcp(container name) for Docker to work
	conn, err := sql.Open("mysql", creds)
	if err != nil {
		fmt.Println("Error can't open DB connection :>", err)
		log.Fatal(err)
	}
	return conn
}
