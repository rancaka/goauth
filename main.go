package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	http.ListenAndServe(":8080", handler())
}

func handler() http.Handler {
	r := httprouter.New()

	r.GET("/", indexHandler)

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello World")
}
