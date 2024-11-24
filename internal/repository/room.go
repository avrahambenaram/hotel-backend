package repository

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type RoomQuery struct {
	Capacity   uint
	Type       uint
	PriceDiary float32
}

type RoomRepository interface {
	FindAll() []entity.HotelRoom
	FindByQuery(query RoomQuery) []entity.HotelRoom
	FindByID(id uint) (entity.HotelRoom, *exception.Exception)
	FindByNumber(number int) (entity.HotelRoom, *exception.Exception)
	Update(room entity.HotelRoom) *exception.Exception
	Save(room entity.HotelRoom) *exception.Exception
	Delete(id uint) *exception.Exception
}
