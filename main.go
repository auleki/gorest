package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endoint Hit: Homepage!")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/all", returnAllArticles)
	router.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":4000", router))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode((Articles))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
}

var Articles []Article

func main() {
	fmt.Println("REST API v2.0 - Mux Router")
	Articles = []Article{
		Article{
			Id:      "1",
			Title:   "Free Nigeria",
			Desc:    "Legal Approaches to making this happen",
			Content: "We discuss legal approaches to breaking the chains shackled to our motherlan ",
		},
		Article{
			Id:      "2",
			Title:   "Suns or Bucks",
			Desc:    "WHo wins the playoffs?",
			Content: "We go into details and history on who might emerge victorious tonight",
		},
	}
	handleRequests()
}
