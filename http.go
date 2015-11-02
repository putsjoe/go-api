package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):] //get everything after the /hello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.ListenAndServe("localhost:9999", nil)
}
