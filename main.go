package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"./db"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/post",db.Acceptpost)
	http.HandleFunc("/get",db.Getrequest)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
  articles:=Articles{}
  for i:=0;i<10;i++{
    title:="Hello_%d"
    articles=append(articles, Article{Title:fmt.Sprintf(title,i),Desc:"Article Description",Content: "Article Content"})
  }
  fmt.Println("Endpoint Hit: returnAllArticles")
  json.NewEncoder(w).Encode(articles)

}
func main() {
	handleRequests()
}
