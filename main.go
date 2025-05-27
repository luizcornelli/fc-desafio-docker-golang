package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Full Cycle Rocks!!")
	
	// Mantém o container rodando
	for {
		time.Sleep(1 * time.Hour) // Mantém o processo ativo
	}
}
