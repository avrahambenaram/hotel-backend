package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avrahambenaram/hotel-backend/internal/configuration"
	"github.com/avrahambenaram/hotel-backend/internal/controller"
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/model"
	repositoryImp "github.com/avrahambenaram/hotel-backend/internal/repository/implementation"
	"github.com/rs/cors"
)

func main() {
	configuration.Setup()
	entity.Setup()
	server := http.NewServeMux()

	clientRepository := &repositoryImp.ClientRepository{}
	clientModel := model.NewClientModel(clientRepository)
	clientController := controller.NewClientController(clientModel)

	roomRepository := &repositoryImp.RoomRepository{}
	roomModel := model.NewRoomModel(roomRepository)
	roomController := controller.NewRoomController(roomModel)

	reservationRepository := &repositoryImp.ReservationRepository{}
	reservationModel := model.NewReservationModel(reservationRepository)
	reservationController := controller.NewReservationController(reservationModel)

	server.Handle("/client/", http.StripPrefix("/client", clientController.Handler))
	server.Handle("/room/", http.StripPrefix("/room", roomController.Handler))
	server.Handle("/reservation/", http.StripPrefix("/reservation", reservationController.Handler))

	corsSetup := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "DELETE"},
	})
	allowedServer := corsSetup.Handler(server)

	log.Printf("Server running on port %d\n", configuration.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", configuration.Server.Port), allowedServer)
}
