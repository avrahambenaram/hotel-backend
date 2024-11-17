package model

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
	"github.com/avrahambenaram/hotel-backend/internal/repository"
)

type ReservationModel struct {
	reservationRepository repository.ReservationRepository
}

func NewReservationModel(reservationRepository repository.ReservationRepository) *ReservationModel {
	return &ReservationModel{
		reservationRepository,
	}
}

func (c *ReservationModel) FindAll() []entity.Reservation {
	return c.reservationRepository.FindAll()
}

func (c *ReservationModel) FindByID(id uint) (entity.Reservation, *exception.Exception) {
	return c.reservationRepository.FindByID(id)
}

func (c *ReservationModel) FindByClientAndRoom(clientID uint, roomID uint) []entity.Reservation {
	return c.reservationRepository.FindByClientAndRoom(clientID, roomID)
}

func (c *ReservationModel) FindByClient(clientID uint) []entity.Reservation {
	return c.reservationRepository.FindByClient(clientID)
}

func (c *ReservationModel) FindByRoom(roomID uint) []entity.Reservation {
	return c.reservationRepository.FindByRoom(roomID)
}

func (c *ReservationModel) Save(reservation entity.Reservation) (entity.Reservation, *exception.Exception) {
	if c.isRoomAlreadyReserved(reservation) {
		return reservation, exception.New("Já há uma reserva no horário solicitado", 409)
	}

	err := c.reservationRepository.Save(reservation)
	if err != nil {
		return reservation, err
	}

	reservations := c.reservationRepository.FindByClientAndRoom(reservation.ClientID, reservation.RoomID)
	index := len(reservations) - 1
	return reservations[index], nil
}

func (c ReservationModel) isRoomAlreadyReserved(reservation entity.Reservation) bool {
	reservCheckIn, _ := c.reservationRepository.FindByRoomAndTime(reservation.RoomID, reservation.CheckIn)
	reservCheckOut, _ := c.reservationRepository.FindByRoomAndTime(reservation.RoomID, reservation.CheckOut)

	if !reservCheckIn.CheckIn.IsZero() || !reservCheckOut.CheckIn.IsZero() {
		return true
	}
	return false
}

func (c *ReservationModel) Delete(id uint) *exception.Exception {
	_, err := c.FindByID(id)
	if err != nil {
		return err
	}

	errDelete := c.reservationRepository.Delete(id)
	if errDelete != nil {
		return errDelete
	}

	return nil
}
