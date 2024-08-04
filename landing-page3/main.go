package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve index.html at the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	// fs := http.FileServer(http.Dir("static1"))
	// http.Handle("/static1/", http.StripPrefix("/static1/", fs))

	fmt.Println("Starting server at port 7998")
	if err := http.ListenAndServe(":7998", nil); err != nil {
		fmt.Println(err)
	}
}
