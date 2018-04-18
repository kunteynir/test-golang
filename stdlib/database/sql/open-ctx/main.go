package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	envDBName = "APP_DB_NAME"
	envDBUser = "APP_DB_USER"
	envDBPWD  = "APP_DB_PWD"
)

func main() {

	connStr := fmt.Sprintf("user=%v password=%v dbname=%v",
		os.Getenv(envDBUser), os.Getenv(envDBPWD), os.Getenv(envDBName))

	// simple
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
	rows, err := db.QueryContext(ctx, "SELECT config FROM synthetic_monitor")
	_ = cancel
	// cancel()
	// defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var conf string
		if err := rows.Scan(&conf); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("conf is %v\n", conf)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
