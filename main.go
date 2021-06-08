package main

import (
	"fmt"
	"net/http"
	"os"
)

func serve(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v, ", name, h)
			fmt.Fprintf(os.Stdout, "%v: %v, ", name, h)
		}
	}
	fmt.Println()
}

func main() {
	http.HandleFunc("/api/post", serve)
	fmt.Println("Starting server on :8090")
	http.ListenAndServe(":8090", nil)
}
