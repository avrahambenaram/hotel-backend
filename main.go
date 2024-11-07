package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avrahambenaram/hotel-backend/internal/configuration"
)

func main() {
	server := http.NewServeMux()

	log.Printf("Server running on port %d\n", configuration.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", configuration.Server.Port), server)
}
