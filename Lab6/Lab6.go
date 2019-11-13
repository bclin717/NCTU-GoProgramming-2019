package main

import (
	"log"
	"net/http"
)

func lab6_2(w http.ResponseWriter, r *http.Request) {

}

func lab6_1(w http.ResponseWriter, r *http.Request) {

}

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
