package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avrahambenaram/hotel-backend/internal/configuration"
	"github.com/avrahambenaram/hotel-backend/internal/controller"
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/model"
	"github.com/avrahambenaram/hotel-backend/internal/repository"
)

func main() {
	configuration.Setup()
	entity.Setup()
	server := http.NewServeMux()

	clientRepository := &repository.ClientRepository{}
	clientModel := model.NewClientModel(clientRepository)
	clientController := controller.NewClientController(clientModel)

	server.Handle("/client/", http.StripPrefix("/client", clientController.Handler))

	log.Printf("Server running on port %d\n", configuration.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", configuration.Server.Port), server)
}
