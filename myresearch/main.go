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
	fs := http.FileServer(http.Dir("static2"))
	http.Handle("/static2/", http.StripPrefix("/static2/", fs))

	fmt.Println("Starting server at port 8002")
	if err := http.ListenAndServe(":8002", nil); err != nil {
		fmt.Println(err)
	}
}
