package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port = flag.Int("p", 8081, "Port to listen on")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("> Starting on http://0.0.0.0:%d\n", *port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Go!")
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Println("Error starting!")
	}
}
