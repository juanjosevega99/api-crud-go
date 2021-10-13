package main

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID: 1,
		Name: "Task one",
		Content: "Some content"
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my GO API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
}
