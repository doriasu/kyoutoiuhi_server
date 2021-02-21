package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"os"
  "github.com/joho/godotenv"
)

type POST struct {
	postID     string
	createdAT  pq.NullTime
	userID     string
	evaluation int
	comment    string
}
func Env_load() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
	Env_load()
	db, err := sql.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=kyoutoiuhi",os.Getenv("USER_NAME"),os.Getenv("PASSWORD")))
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT * FROM post")
  if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		fmt.Println(rows)
	}

}
