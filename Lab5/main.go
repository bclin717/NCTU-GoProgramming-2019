package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.SplitN(r.URL.Path, "/", -1)

	operation := pathParts[1]
	a, _ := strconv.Atoi(pathParts[2])
	b, _ := strconv.Atoi(pathParts[3])

	if operation == "add" {

	} else if operation == "sub" {

	} else if operation == "mul" {

	} else if operation == "div" {

	} else {
		fmt.Fprintf(w, "hello world!")
	}
}

func add(a, b int) int {

}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
