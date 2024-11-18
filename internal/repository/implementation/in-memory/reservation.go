package inmemory

import (
	"slices"
	"time"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type ReservationRepository struct {
	reservations []entity.Reservation
}

func (c ReservationRepository) FindAll() []entity.Reservation {
	return c.reservations
}

func (c ReservationRepository) FindByID(id uint) (entity.Reservation, *exception.Exception) {
	for _, reservations := range c.reservations {
		if reservations.ID == id {
			return reservations, nil
		}
	}
	return entity.Reservation{}, exception.New("Reserva não encontrada", 404)
}

func (c ReservationRepository) FindByClientAndRoom(clientID uint, roomID uint) []entity.Reservation {
	reservations := []entity.Reservation{}
	for _, reservation := range c.reservations {
		if reservation.ClientID == clientID && reservation.RoomID == roomID {
			reservations = append(reservations, reservation)
		}
	}
	return reservations
}

func (c ReservationRepository) FindByClient(clientID uint) []entity.Reservation {
	reservations := []entity.Reservation{}
	for _, reservation := range c.reservations {
		if reservation.ClientID == clientID {
			reservations = append(reservations, reservation)
		}
	}
	return reservations
}

func (c ReservationRepository) FindByRoom(roomID uint) []entity.Reservation {
	reservations := []entity.Reservation{}
	for _, reservation := range c.reservations {
		if reservation.RoomID == roomID {
			reservations = append(reservations, reservation)
		}
	}
	return reservations
}

func (c ReservationRepository) FindByRoomAndTime(roomID uint, when time.Time) (entity.Reservation, *exception.Exception) {
	whenUnix := when.Unix()
	for _, reservation := range c.reservations {
		checkInUnix := reservation.CheckIn.Unix()
		checkOutUnix := reservation.CheckOut.Unix()
		if reservation.RoomID == roomID && whenUnix >= checkInUnix && whenUnix <= checkOutUnix {
			return reservation, nil
		}
	}

	return entity.Reservation{}, exception.New("Reserva não encontrada", 404)
}

func (c *ReservationRepository) Save(reservation entity.Reservation) *exception.Exception {
	c.reservations = append(c.reservations, reservation)
	return nil
}

func (c *ReservationRepository) Delete(id uint) *exception.Exception {
	for i, current := range c.reservations {
		if current.ID == id {
			c.reservations = slices.Delete(c.reservations, i, i+1)
			return nil
		}
	}
	return exception.New("Reserva não encontrada", 404)
}
