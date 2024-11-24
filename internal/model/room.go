package model

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
	"github.com/avrahambenaram/hotel-backend/internal/repository"
)

type RoomModel struct {
	roomRepository repository.RoomRepository
}

func NewRoomModel(roomRepository repository.RoomRepository) *RoomModel {
	return &RoomModel{
		roomRepository,
	}
}

func (c *RoomModel) FindAll() []entity.HotelRoom {
	return c.roomRepository.FindAll()
}

func (c *RoomModel) FindByQuery(query repository.RoomQuery) []entity.HotelRoom {
	return c.roomRepository.FindByQuery(query)
}

func (c *RoomModel) FindByID(id uint) (entity.HotelRoom, *exception.Exception) {
	return c.roomRepository.FindByID(id)
}

func (c *RoomModel) FindByNumber(number int) (entity.HotelRoom, *exception.Exception) {
	return c.roomRepository.FindByNumber(number)
}

func (c *RoomModel) Update(room entity.HotelRoom) (entity.HotelRoom, *exception.Exception) {
	_, err := c.roomRepository.FindByID(room.ID)
	if err != nil {
		return entity.HotelRoom{}, err
	}

	roomNumberUsed, _ := c.roomRepository.FindByNumber(room.Number)
	if roomNumberUsed.Capacity != 0 && roomNumberUsed.ID != room.ID {
		return entity.HotelRoom{}, exception.New("Número de quarto já usado", 409)
	}

	if !room.Type.IsValid() {
		return entity.HotelRoom{}, exception.New("Tipo de quarto inválido", 403)
	}

	errUpdate := c.roomRepository.Update(room)
	if errUpdate != nil {
		return entity.HotelRoom{}, errUpdate
	}

	roomUpdated, _ := c.roomRepository.FindByID(room.ID)
	return roomUpdated, nil
}

func (c *RoomModel) Save(room entity.HotelRoom) (entity.HotelRoom, *exception.Exception) {
	roomNumberUsed, _ := c.roomRepository.FindByNumber(room.Number)
	if roomNumberUsed.Capacity != 0 {
		return entity.HotelRoom{}, exception.New("Número de quarto já usado", 409)
	}

	if !room.Type.IsValid() {
		return entity.HotelRoom{}, exception.New("Tipo de quarto inválido", 403)
	}

	errSave := c.roomRepository.Save(room)
	if errSave != nil {
		return entity.HotelRoom{}, errSave
	}

	roomSaved, _ := c.roomRepository.FindByNumber(room.Number)
	return roomSaved, nil
}

func (c *RoomModel) Delete(id uint) *exception.Exception {
	_, err := c.roomRepository.FindByID(id)
	if err != nil {
		return err
	}

	errDelete := c.roomRepository.Delete(id)
	if errDelete != nil {
		return errDelete
	}

	return nil
}
