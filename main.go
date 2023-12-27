package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "title: %s\n", vars["title"])
	fmt.Fprintf(w, "page: %s\n", vars["page"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", getPage)
	log.Fatal(http.ListenAndServe(":8080", r))
}
