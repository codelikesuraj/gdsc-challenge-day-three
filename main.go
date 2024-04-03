package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	PORT = "3000"

	DB_USER     = "your_database_username"
	DB_PASS     = "your_database_password"
	DB_DATABASE = "your_database_name"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		dsn := fmt.Sprintf("%s:%s@/%s", DB_USER, DB_PASS, DB_DATABASE)

		sql, err := sql.Open("mysql", dsn)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Error connecting to database - %s", err.Error())))
			return
		}

		if err = sql.Ping(); err != nil {
			w.Write([]byte(fmt.Sprintf("Error pinging database - %s", err.Error())))
			return
		}

		w.Write([]byte("Database connected successfully!"))
	})

	fmt.Printf("Server listening on %q\n", PORT)
	fmt.Printf("Visit the URL http://127.0.0.1:%s to check if the connection was successful\n", PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
