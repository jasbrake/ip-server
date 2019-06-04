package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ipHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, req.Header.Get("X-Forwarded-For"))
	}

	http.HandleFunc("/", ipHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
