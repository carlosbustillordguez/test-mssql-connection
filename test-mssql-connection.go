package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/xo/dburl"
)

var db *sql.DB
var databaseUrl = os.Getenv("DATABASE_URL")

func main() {
	// Check if the DATABASE_URL is already defined
	if databaseUrl == "" {
		msg1 := "The environment variable 'DATABASE_URL' is not defined.\n"
		msg2 := "Use the format: 'sqlserver://username:password@host:port?database=dbName&param1=value'"
		log.Fatal(msg1 + msg2)
	}

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", databaseUrl)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
	u, err := dburl.Parse(databaseUrl)
	if err == nil {
		fmt.Printf("Connection details:\n")
		fmt.Printf(" * Host: %s \n", u.Hostname())
		fmt.Printf(" * Port: %s \n", u.Port())
		fmt.Printf(" * Username: %s \n", u.User.Username())
		fmt.Printf(" * Password: **** \n")
		fmt.Printf(" * Options: %s", u.Query().Encode())

	}
}
