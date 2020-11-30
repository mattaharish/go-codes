package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello! Matta!!</h1>")
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	name := strings.TrimPrefix(r.URL.Path, "/greet/")
	fmt.Fprintf(w, "<h1>Hello!"+name+"!!</h1>")
}

func hello() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/greet/", greet)
	fmt.Println("Server is starting..")
	http.ListenAndServe(":3000", nil)
	fmt.Println("Served is running on port 3000")
}
