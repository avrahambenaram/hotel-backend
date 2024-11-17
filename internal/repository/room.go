package repository

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type RoomRepository interface {
	FindAll() []entity.HotelRoom
	FindByID(id uint) (entity.HotelRoom, *exception.Exception)
	FindByType(roomType entity.RoomType) []entity.HotelRoom
	FindByNumber(number int) (entity.HotelRoom, *exception.Exception)
	Update(room entity.HotelRoom) *exception.Exception
	Save(room entity.HotelRoom) *exception.Exception
	Delete(id uint) *exception.Exception
}
