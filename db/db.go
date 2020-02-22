// Pacakge provides the db object to be used by the other packages (fetchrss,searchress,serverss)
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Conn() *sql.DB {
	conn, err := sql.Open("mysql", "root:admin@/rssfeeds")
	if err != nil {
		fmt.Println("Error in db opening :>", err)
		log.Fatal(err)
	}
	return conn
}
