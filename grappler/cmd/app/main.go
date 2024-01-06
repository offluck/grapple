package main

import "net/http"

func hello(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("heya"))
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(hello))
}
