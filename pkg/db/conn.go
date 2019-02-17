package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "online-store"
)

// DBdriver will be used for all database interactions
var DBdriver *sql.DB

func init() {
	var er error
	connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	DBdriver, er = sql.Open("postgres", connString)
	if er != nil {
		log.Print("Error in connecting to the database: ", er)
		return
	}
	fmt.Println("Connected to the database")
}
