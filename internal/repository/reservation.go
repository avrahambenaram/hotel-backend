package repository

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type ReservationRepository struct{}

func (c ReservationRepository) FindByID(id uint) (entity.Reservation, *exception.Exception) {
	var reservation entity.Reservation
	entity.DB.Preload("Room").Preload("Client").First(&reservation, id)
	if reservation.CheckIn.IsZero() {
		return reservation, exception.New("Reserva não encontrada", 404)
	}
	return reservation, nil
}

func (c ReservationRepository) FindByClientAndRoom(clientID uint, roomID uint) []entity.Reservation {
	reservations := []entity.Reservation{}
	entity.DB.Where("clientID = ? AND roomID = ?", clientID, roomID).Find(&reservations)
	return reservations
}

func (c ReservationRepository) FindByClient(clientID uint) []entity.Reservation {
	reservations := []entity.Reservation{}
	entity.DB.Preload("Room").Preload("Client").Where("clientID = ?", clientID).Find(&reservations)
	return reservations
}

func (c ReservationRepository) FindByRoom(roomID uint) []entity.Reservation {
	reservations := []entity.Reservation{}
	entity.DB.Where("roomID = ?", roomID).Find(&reservations)
	return reservations
}

func (c ReservationRepository) Update(reservation entity.Reservation) *exception.Exception {
	_, err := c.FindByID(reservation.ID)
	if err != nil {
		return err
	}
	entity.DB.Save(&reservation)
	return nil
}

func (c ReservationRepository) Save(reservation entity.Reservation) *exception.Exception {
	result := entity.DB.Create(&reservation)
	if result.RowsAffected != 1 {
		return exception.New("Erro ao criar a reserva", 500)
	}
	return nil
}

func (c ReservationRepository) Delete(id uint) *exception.Exception {
	result := entity.DB.Delete(&entity.Reservation{}, id)
	if result.RowsAffected != 1 {
		return exception.New("Erro ao deletar a reserva", 500)
	}
	return nil
}
