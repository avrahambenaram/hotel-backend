package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/middleware"
	"github.com/avrahambenaram/hotel-backend/internal/model"
)

type ReservationController struct {
	Handler          http.Handler
	reservationModel *model.ReservationModel
}

func NewReservationController(reservationModel *model.ReservationModel) *ReservationController {
	mux := http.NewServeMux()
	reservationController := &ReservationController{
		middleware.SendJSON(mux),
		reservationModel,
	}

	mux.Handle(
		"GET /id/{ID}",
		middleware.GetId(
			http.HandlerFunc(reservationController.findByID),
		),
	)
	mux.Handle(
		"GET /{$}",
		http.HandlerFunc(reservationController.query),
	)
	mux.Handle(
		"POST /add",
		middleware.ParseBody(
			http.HandlerFunc(reservationController.addReservation),
			entity.Reservation{},
		),
	)
	mux.Handle(
		"DELETE /{ID}",
		middleware.GetId(
			http.HandlerFunc(reservationController.findByID),
		),
	)

	return reservationController
}

func (c *ReservationController) findByID(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(uint)

	reservation, err := c.reservationModel.FindByID(id)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", reservation)
	*r = *r.WithContext(ctx)
}

func (c *ReservationController) query(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	clientStr := query.Get("client")
	roomStr := query.Get("room")
	var clientID uint
	var roomID uint

	if clientStr != "" {
		client, err := strconv.Atoi(clientStr)
		if err != nil {
			http.Error(w, "ID do cliente inválido", http.StatusForbidden)
			return
		}
		clientID = uint(client)
	}
	if roomStr != "" {
		room, err := strconv.Atoi(roomStr)
		if err != nil {
			http.Error(w, "ID do quarto inválido", http.StatusForbidden)
			return
		}
		roomID = uint(room)
	}

	if clientStr != "" && roomStr != "" {
		reservations := c.reservationModel.FindByClientAndRoom(clientID, roomID)
		ctx := context.WithValue(r.Context(), "json", reservations)
		*r = *r.WithContext(ctx)
		return
	}
	if clientStr != "" {
		reservations := c.reservationModel.FindByClient(clientID)
		ctx := context.WithValue(r.Context(), "json", reservations)
		*r = *r.WithContext(ctx)
		return
	}
	if roomStr != "" {
		reservations := c.reservationModel.FindByRoom(roomID)
		ctx := context.WithValue(r.Context(), "json", reservations)
		*r = *r.WithContext(ctx)
		return
	}

	reservations := c.reservationModel.FindAll()

	ctx := context.WithValue(r.Context(), "json", reservations)
	*r = *r.WithContext(ctx)
}

func (c *ReservationController) addReservation(w http.ResponseWriter, r *http.Request) {
	reservation := r.Context().Value("body").(entity.Reservation)
	reservationCreated, err := c.reservationModel.Save(reservation)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", reservationCreated)
	*r = *r.WithContext(ctx)
}

func (c *ReservationController) delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(uint)

	errDelete := c.reservationModel.Delete(id)
	if errDelete != nil {
		http.Error(w, errDelete.Message, errDelete.Status)
		return
	}

	w.WriteHeader(204)
}
