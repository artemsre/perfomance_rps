package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func main() {

	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
