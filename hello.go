package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/gg", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "you are in the endgame")
	})
	http.ListenAndServe(":8080", nil)

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s", r.URL.Path[1:])
}
