package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//function callback for url endpoint
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcom")
}
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	//specify the content type
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//if request to todos, show static values
	//NewEncoder: writes to w
	//Encode: endoe todos to json and writes to stream
	//Output
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	//ioutil: read from r, return (data,error)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	//parse json
	if err := json.Unmarshal(body, &todo); err != nil {
		//if error occurs, send error back in json to client
		w.Header().Set("Content-Type", "applicatoin/json;charset=UTF-8")
		w.WriteHeader(422) //uprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//create id
	t := RepoCreateTodo(todo)

	//send back success to client
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
