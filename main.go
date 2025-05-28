package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Define o cabeçalho como HTML
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    
    // HTML simples com estilo
    html := `
    <!DOCTYPE html>
    <html>
        <head>
            <title>Full Cycle Rocks!</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    margin: 0;
                    background-color: #f0f0f0;
                }
                .message {
                    padding: 20px;
                    background-color: #4CAF50;
                    color: white;
                    border-radius: 5px;
                    box-shadow: 0 2px 5px rgba(0,0,0,0.2);
                }
            </style>
        </head>
        <body>
            <div class="message">
                <h1>Full Cycle Rocks!!</h1>
            </div>
        </body>
    </html>
    `
    
    fmt.Fprint(w, html)
}

func main() {
    // Registra o handler para a rota raiz
    http.HandleFunc("/", handler)
    
    // Log de inicialização
    fmt.Println("Server starting on port 8081...")
    
    // Inicia o servidor na porta 8081
    if err := http.ListenAndServe(":8081", nil); err != nil {
        fmt.Printf("Server failed to start: %v\n", err)
    }
}
