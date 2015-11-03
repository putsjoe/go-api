package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"reflect"
)

func Handlr(w http.ResponseWriter, r *http.Request) {
	givenURL := r.URL.Path[len("/add/"):]
	fmt.Fprintf(w, "URL:  %s \n", givenURL)
	resp, err := http.Get("http://www.google.co.uk")
	fmt.Println("Given Error is: ", err)

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Body Read error is: ", err)
	fmt.Println("Body type - ", reflect.TypeOf(body))

	//Parse html
	prs, err := html.Parse(resp.Body)
	fmt.Println("Body type after Parse - ", reflect.TypeOf(prs))

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			fmt.Println(n.Data)
		}
		fmt.Println(n.Data)
	}
	f(prs)
	fmt.Println(prs.Data)

	//	for ab := range prs {
	//		fmt.Println(ab[
	//	}

	//fmt.Fprintf(w, "\n", string(body))
	//	resp.Body.Close();

}

func main() {
	http.HandleFunc("/add/", Handlr)
	http.ListenAndServe("212.18.232.164:1234", nil)
}
