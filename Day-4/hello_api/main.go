package main

import "net/http"

func main() {
	http.HandleFunc("/hello", helloHeader)

	http.ListenAndServe("localhost:4000", nil)
}

func helloHeader(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API"))
}