package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func Handlr(w http.ResponseWriter, r *http.Request) {
	givenURL := r.URL.Path[len("/add/"):]
	fmt.Fprintf(w, "URL:  %s \n", givenURL)
	resp, err := http.Get("http://www.google.co.uk")
	fmt.Println("Given Error is: ", err)

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(w, "Body Read error is: ", err)

	fmt.Fprintf(w, "\n", string(body))
//	resp.Body.Close();


}

func main() {
	http.HandleFunc("/add/", Handlr)
	http.ListenAndServe("212.18.232.164:1234", nil)
}
