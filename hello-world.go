package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{name}", func(w http.ResponseWriter, r1 *http.Request) {
		vars := mux.Vars(r1)
		title := strings.ToUpper(vars["name"])
		fmt.Fprintf(w, "Your name is %s", title)
	})

	http.ListenAndServe(":80", r)
}
