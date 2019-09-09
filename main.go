package main

import (
	"fmt"
	"net/http"

	C "github.com/Tech4GoodPH/go-sample-api.git/cmd"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, C.HomeMessage())
}

func main() {
	C.StartServer()
}
