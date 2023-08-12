package main

import (
	"fmt"
	"log"
	"net/http"
)

func httphandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, my name is ayeni clement , I love you all %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", httphandler)

	log.Fatal(http.ListenAndServe(":2020", nil))
}