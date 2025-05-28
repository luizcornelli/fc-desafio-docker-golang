package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    fmt.Println("Full Cycle Rocks!!")

    // Configurar o servidor HTTP em uma goroutine
    go func() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "text/html")
            fmt.Fprintf(w, "<h1>Full Cycle Rocks!!</h1>")
        })
        
        http.ListenAndServe(":8081", nil)
    }()

    // Mantém o container rodando
    for {
        time.Sleep(1 * time.Hour)
    }
}
