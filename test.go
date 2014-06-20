package main

import (
	"database/sql"
	_ "github.com/jbarham/gopgsqldriver"
	"os"
)

func main() {
	url = os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	var greeting string
	err := db.QueryRow("SELECT 'hello world'::text").Scan(&greeting)
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)
}
