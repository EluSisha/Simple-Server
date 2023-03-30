package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, i am mini server!\n"))

}

func main() {
	http.HandleFunc("/", Handler)

	log.Println("Start http server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
