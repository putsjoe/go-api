package main

import (
	"fmt"
	"net/http"
)

func Handlr(w http.ResponseWriter, r *http.Request) {
	givenURL := r.URL.Path[len("/add/"):]
	fmt.Fprintf(w, "URL:  %s \n", givenURL)
	fmt.Fprintf(w, "\nBob")

}

func main() {
	http.HandleFunc("/add/", Handlr)
	http.ListenAndServe("localhost:1234", nil)
}
