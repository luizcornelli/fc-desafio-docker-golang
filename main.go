package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Full Cycle Rocks!!")
	})

	fmt.Println("Server running on port 8081")
	http.ListenAndServe(":8080", nil)
}
