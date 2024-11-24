package implementation

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
	"github.com/avrahambenaram/hotel-backend/internal/repository"
	"gorm.io/gorm"
)

type RoomRepository struct{}

func (c RoomRepository) FindAll() []entity.HotelRoom {
	rooms := []entity.HotelRoom{}
	entity.DB.Find(&rooms)
	return rooms
}

func (c RoomRepository) FindByQuery(query repository.RoomQuery) []entity.HotelRoom {
	rooms := []entity.HotelRoom{}
	queries := []*gorm.DB{}

	if query.Capacity != 0 {
		queries = append(queries, entity.DB.Where("capacity = ?", query.Capacity))
	}
	if query.Type != 0 {
		queries = append(queries, entity.DB.Where("type = ?", query.Type))
	}
	if query.PriceDiary != 0 {
		queries = append(queries, entity.DB.Where("price_diary = ?", query.PriceDiary))
	}

	var finalQuery *gorm.DB
	for i, queryORM := range queries {
		if i == 0 {
			finalQuery = entity.DB.Where(queryORM)
			continue
		}
		finalQuery.Where(queryORM)
	}
	finalQuery.Find(&rooms)
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
