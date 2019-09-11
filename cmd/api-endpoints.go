package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeMessage ...
func HomeMessage() string {
	return "Tech 4 Good !!!"
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, HomeMessage())
}

// StartServer ...
func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/message", homeLink)
	fmt.Println("Exposing API endpoints on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
