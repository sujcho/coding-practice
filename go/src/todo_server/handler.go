package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//function callback for url endpoint
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcom")
}
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Feed Mungee"},
		Todo{Name: "No Study"},
	}
	//if request to todos, show static values
	//NewEncoder: writes to w
	//Encode: endoe todos to json and writes to stream
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
