package implementation

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type RoomRepository struct{}

func (c RoomRepository) FindAll() []entity.HotelRoom {
	rooms := []entity.HotelRoom{}
	entity.DB.Find(&rooms)
	return rooms
}

func (c RoomRepository) FindByID(id uint) (entity.HotelRoom, *exception.Exception) {
	room := entity.HotelRoom{}
	entity.DB.Where("ID = ?", id).Find(&room)
	if room.Capacity == 0 {
		return room, exception.New("Quarto não encontrado", 404)
	}
	return room, nil
}

func (c RoomRepository) FindByType(roomType entity.RoomType) []entity.HotelRoom {
	rooms := []entity.HotelRoom{}
	entity.DB.Where("Type = ?", roomType).Find(&rooms)
	return rooms
}

func (c RoomRepository) FindByNumber(number int) (entity.HotelRoom, *exception.Exception) {
	room := entity.HotelRoom{}
	entity.DB.Where("Number = ?", number).Find(&room)
	if room.Capacity == 0 {
		return room, exception.New("Quarto não encontrado", 404)
	}
	return room, nil
}

func (c RoomRepository) Update(room entity.HotelRoom) *exception.Exception {
	_, err := c.FindByID(room.ID)
	if err != nil {
		return err
	}
	entity.DB.Save(&room)
	return nil
}

func (c RoomRepository) Save(room entity.HotelRoom) *exception.Exception {
	result := entity.DB.Create(&room)
	if result.RowsAffected != 1 {
		return exception.New("Erro ao criar o quarto", 500)
	}
	return nil
}

func (c RoomRepository) Delete(id uint) *exception.Exception {
	result := entity.DB.Delete(&entity.HotelRoom{}, id)
	if result.RowsAffected != 1 {
		return exception.New("Erro ao deletar o quarto", 500)
	}
	return nil
}
