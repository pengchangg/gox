package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
	})

	// favicon.ico
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./favicon.ico")
	})

	// 404
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./404.html")
	})

	// 500
	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./500.html")
	})

  // post mothon create post
  http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "post")
  })

  log.Println("Server is running on port http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
