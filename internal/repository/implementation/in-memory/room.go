package inmemory

import (
	"slices"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
	"github.com/avrahambenaram/hotel-backend/internal/repository"
)

type RoomRepository struct {
	rooms []entity.HotelRoom
}

func (c RoomRepository) FindAll() []entity.HotelRoom {
	return c.rooms
}

func (c RoomRepository) FindByQuery(query repository.RoomQuery) []entity.HotelRoom {
	rooms := []entity.HotelRoom{}
	for _, room := range c.rooms {
		if query.Capacity != 0 && room.Capacity != query.Capacity {
			continue
		}
		if query.Type != 0 && room.Type != entity.RoomType(query.Type) {
			continue
		}
		if query.PriceDiary != 0 && room.PriceDiary != query.PriceDiary {
			continue
		}
		rooms = append(rooms, room)
	}
	return rooms
}

func (c RoomRepository) FindByID(id uint) (entity.HotelRoom, *exception.Exception) {
	for _, room := range c.rooms {
		if room.ID == id {
			return room, nil
		}
	}
	return entity.HotelRoom{}, exception.New("Quarto não encontrado", 404)
}

func (c RoomRepository) FindByNumber(number int) (entity.HotelRoom, *exception.Exception) {
	for _, room := range c.rooms {
		if room.Number == number {
			return room, nil
		}
	}
	return entity.HotelRoom{}, exception.New("Quarto não encontrado", 404)
}

func (c *RoomRepository) Update(room entity.HotelRoom) *exception.Exception {
	for i, current := range c.rooms {
		if current.ID == room.ID {
			c.rooms[i] = room
			return nil
		}
	}

	return exception.New("Quarto não encontrado", 404)
}

func (c *RoomRepository) Save(room entity.HotelRoom) *exception.Exception {
	c.rooms = append(c.rooms, room)
	return nil
}

func (c *RoomRepository) Delete(id uint) *exception.Exception {
	for i, current := range c.rooms {
		if current.ID == id {
			c.rooms = slices.Delete(c.rooms, i, i+1)
			return nil
		}
	}
	return exception.New("Quarto não encontrado", 404)
}
