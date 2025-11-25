package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:8080/gg")
	if err != nil {
		log.Fatal("http get failure ", err)
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	fmt.Printf("%s\n", b)
}
