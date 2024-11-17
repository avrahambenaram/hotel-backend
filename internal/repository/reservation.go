package repository

import (
	"time"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type ReservationRepository interface {
	FindAll() []entity.Reservation
	FindByID(id uint) (entity.Reservation, *exception.Exception)
	FindByClientAndRoom(clientID uint, roomID uint) []entity.Reservation
	FindByClient(clientID uint) []entity.Reservation
	FindByRoom(roomID uint) []entity.Reservation
	FindByRoomAndTime(roomID uint, when time.Time) (entity.Reservation, *exception.Exception)
	Save(reservation entity.Reservation) *exception.Exception
	Delete(id uint) *exception.Exception
}
