package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func activateRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", r)
}
