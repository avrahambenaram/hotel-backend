package implementation

import (
	"time"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type ReservationRepository struct{}

func (c ReservationRepository) FindAll() []entity.Reservation {
	reservations := []entity.Reservation{}
	entity.DB.Preload("Room").Preload("Client").Find(&reservations)
	return reservations
}

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
	entity.DB.Preload("Room").Preload("Client").Where("client_id = ? AND room_id = ?", clientID, roomID).Find(&reservations)
	return reservations
}

func (c ReservationRepository) FindByClient(clientID uint) []entity.Reservation {
	reservations := []entity.Reservation{}
	entity.DB.Preload("Room").Preload("Client").Where("client_id = ?", clientID).Find(&reservations)
	return reservations
}

func (c ReservationRepository) FindByRoom(roomID uint) []entity.Reservation {
	reservations := []entity.Reservation{}
	entity.DB.Preload("Room").Preload("Client").Where("room_id = ?", roomID).Find(&reservations)
	return reservations
}

func (c ReservationRepository) FindByRoomAndTime(roomID uint, when time.Time) (entity.Reservation, *exception.Exception) {
	var reservation entity.Reservation
	entity.DB.Preload("Room").Preload("Client").Where("room_id = ? AND ? >= check_in AND ? <= check_out", roomID, when, when).Find(&reservation)
	if reservation.CheckIn.IsZero() {
		return reservation, exception.New("Reserva não encontrada", 404)
	}
	return reservation, nil
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
