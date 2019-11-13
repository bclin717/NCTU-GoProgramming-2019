package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func hacked(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Oh... No! Your website is hacked!")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
	r.ParseForm() // Parse the parameters of URL(GET) or Request body(POST)
	fmt.Fprintf(w, "The method is: %s\n", r.Method)
	fmt.Fprintf(w, "The form is: %s\n", r.Form)
	fmt.Fprintf(w, "The path is: %s\n", r.URL.Path)
	fmt.Fprintf(w, "The p3 parameter is: %s\n\n", r.Form["p3"])
	for k, v := range r.Form {
		fmt.Fprintf(w, "Parameter: %s\n", k)
		fmt.Fprintf(w, "Value: %s\n", strings.Join(v, ""))
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/hacked", hacked)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
