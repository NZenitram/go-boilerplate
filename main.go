package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", indexPageHandler)

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
