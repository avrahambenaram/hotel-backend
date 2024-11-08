package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avrahambenaram/hotel-backend/internal/configuration"
	"github.com/avrahambenaram/hotel-backend/internal/entity"
)

func main() {
	configuration.Setup()
	entity.Setup()
	server := http.NewServeMux()

	log.Printf("Server running on port %d\n", configuration.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", configuration.Server.Port), server)
}
