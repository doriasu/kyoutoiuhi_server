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
type GET struct {
	UserID string `json:"user_id"`
	Year   int    `json:year`
	Month  int    `json:month`
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

//Getrequest カレンダー取得リクエストを受ける
func Getrequest(w http.ResponseWriter, r *http.Request) {
	var getRequest GET
	json.NewDecoder(r.Body).Decode(&getRequest)
	getFromdb(getRequest)

}
func getFromdb(get GET) {
	Env_load()
	db, err := sql.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=kyoutoiuhi", os.Getenv("USER_NAME"), os.Getenv("PASSWORD")))
	if err != nil {
		fmt.Println(err)
	}
	result,err:=db.Query(fmt.Sprintf("select * from post where DATE(created_at) >= '%d-%d-1' and DATE(created_at)<'%d-%d-1'",get.Year,get.Month,get.Year,get.Month+1))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(fmt.Sprintf("select * from post where DATE(created_at) >= '2021-%d-1' and DATE(created_at)<'2021-%d-1';",get.Month,get.Month+1))
	var posts []POST
	for result.Next() {
		var post POST
		result.Scan(&post.PostID, &post.CreatedAT, &post.UserID, &post.Evaluation, &post.Comment)
		posts = append(posts, post)
	}
	fmt.Println(posts)

}

//Acceptpost 書き込むためのpostリクエストを受ける
func Acceptpost(w http.ResponseWriter, r *http.Request) {
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
