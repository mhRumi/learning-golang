package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_NAME")
	connectionString := "postgres://" + dbUser + ":@" + dbHost + "/" + dbName + "?sslmode=disable"
	// "postgres://postgres:@192.168.43.198/swesociety?sslmode=disable"
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Database connection successful...")
	}
}
