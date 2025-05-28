package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, "<h1>Full Cycle Rocks!!</h1>")
    })
    
    fmt.Println("Server starting on port 8081...")
    http.ListenAndServe(":8081", nil)
}
