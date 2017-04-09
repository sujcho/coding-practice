package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/mux"
)

//set attribute to be lower case in Json
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func main() {

	//new router instance
	router := mux.NewRouter()
	router.HandleFunc("/index", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//function callback for url endpoint
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Feed Mungee"},
		Todo{Name: "No Study"},
	}

	//if request to todos, show static values
	//NewEncoder: writes to w
	//Encode: endoe todos to json and writes to stream
	json.NewEncoder(w).Encode(todos)

	//fmt.Fprintf(w, "Todo Index!")
}
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
