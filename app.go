package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "local-db"
	port     = 5432
	user     = "test"
	password = "testtset"
	dbname   = "test-db"
)

func main() {
	fmt.Println("test")
	connectDB()

	http.HandleFunc("/ping", ping)

	http.ListenAndServe(":8080", nil)
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Println("pong")
}

func connectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB OK")
}
