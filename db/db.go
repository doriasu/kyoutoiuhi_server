package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // here
)

type POST struct {
	PostID     string `json:"post_id"`
	CreatedAT  string `json:"created_at"`
	UserID     string `json:"user_id"`
	Evaluation int    `json:"evaluation"`
	Comment    string `json:"comment"`
}

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func init() {
	Env_load()
	db, err := sql.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=kyoutoiuhi", os.Getenv("USER_NAME"), os.Getenv("PASSWORD")))
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		fmt.Println(rows)
	}

}
func Accept_post(w http.ResponseWriter, r *http.Request) {
	// request bodyの読み取り
	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println("io error")
	// 	return
	// }

	// // jsonのdecode
	// jsonBytes := ([]byte)(b)
	// data := new(POST)
	// if err := json.Unmarshal(jsonBytes, &data); err != nil {
	// 	fmt.Println("JSON Unmarshal error:", err)
	// 	fmt.Fprintln(w, "うまく行きませんでした")
	// 	return
	// }
	// fmt.Println(r.Body)
	var data POST
	json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data.Comment)
	insertPost(data)
	fmt.Fprintln(w, "書き込みました")
}
func insertPost(post POST) {
	Env_load()
	db, err := sql.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=kyoutoiuhi", os.Getenv("USER_NAME"), os.Getenv("PASSWORD")))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(fmt.Sprintf("insert into post (post_id,created_at,user_id,evaluation,comment) values (%s,%s,%s,%d,%s);", post.postID, post.createdAT, post.userID, post.evaluation, post.comment))
	_, err = db.Query(fmt.Sprintf("insert into post (post_id,created_at,user_id,evaluation,comment) values ('%s','%s','%s',%d,'%s');", post.PostID, post.CreatedAT, post.UserID, post.Evaluation, post.Comment))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return

}
