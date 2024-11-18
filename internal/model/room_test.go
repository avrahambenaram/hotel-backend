package model_test

import (
	"testing"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/model"
	inmemory "github.com/avrahambenaram/hotel-backend/internal/repository/implementation/in-memory"
	"github.com/stretchr/testify/assert"
)

type RoomSuite struct {
	room1      entity.HotelRoom
	room2      entity.HotelRoom
	room3      entity.HotelRoom
	roomToSave entity.HotelRoom

	roomModel *model.RoomModel
}

func SetupRoomSuite() *RoomSuite {
	roomRepository := new(inmemory.RoomRepository)
	roomModel := model.NewRoomModel(roomRepository)

	room1 := entity.HotelRoom{
		ID:         0,
		Number:     1,
		Type:       entity.SimpleRoom,
		Capacity:   1,
		PriceDiary: 40,
	}
	room2 := entity.HotelRoom{
		ID:         1,
		Number:     2,
		Type:       entity.FamilyRoom,
		Capacity:   4,
		PriceDiary: 120,
	}
	room3 := entity.HotelRoom{
		ID:         2,
		Number:     3,
		Type:       entity.FamilyRoom,
		Capacity:   4,
		PriceDiary: 120,
	}
	roomToSave := entity.HotelRoom{
		ID:         3,
		Number:     4,
		Type:       entity.Suite,
		Capacity:   2,
		PriceDiary: 100,
	}
	roomRepository.Save(room1)
	roomRepository.Save(room2)
	roomRepository.Save(room3)

	return &RoomSuite{
		room1,
		room2,
		room3,
		roomToSave,
		roomModel,
	}
}

func TestGetAllRooms(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	rooms := suite.roomModel.FindAll()

	assert.Len(rooms, 3)
}

func TestGetRoomByIDNotFound(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	_, err := suite.roomModel.FindByID(suite.roomToSave.ID)

	if assert.NotNil(err) {
		assert.Equal(404, err.Status)
	}
}

func TestGetRoomByIDSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	room, _ := suite.roomModel.FindByID(suite.room1.ID)

	assert.Equal(suite.room1.Number, room.Number)
}

func TestGetRoomByNumberNotFound(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	_, err := suite.roomModel.FindByNumber(suite.roomToSave.Number)

	if assert.NotNil(err) {
		assert.Equal(404, err.Status)
	}
}

func TestGetRoomByNumberSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	room, _ := suite.roomModel.FindByNumber(suite.room1.Number)

	assert.Equal(suite.room1.ID, room.ID)
}

func TestGetRoomsByInvalidType(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	_, err := suite.roomModel.FindByType(100)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestGetRoomsByFamilyType(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	rooms, _ := suite.roomModel.FindByType(uint(entity.FamilyRoom))

	assert.Len(rooms, 2)
}

func TestGetRoomsBySimpleType(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	rooms, _ := suite.roomModel.FindByType(uint(entity.SimpleRoom))

	assert.Len(rooms, 1)
}

func TestGetRoomsBySuiteType(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	rooms, _ := suite.roomModel.FindByType(uint(entity.Suite))

	assert.Len(rooms, 0)
}

func TestUpdateRoomNumberAlreadyInUse(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()
	updateRoom := suite.room1
	updateRoom.Number = suite.room2.Number

	_, err := suite.roomModel.Update(updateRoom)

	if assert.NotNil(err) {
		assert.Equal(409, err.Status)
	}
}

func TestUpdateRoomInvalidType(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()
	updateRoom := suite.room1
	updateRoom.Type = entity.RoomType(100)

	_, err := suite.roomModel.Update(updateRoom)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestUpdateRoomSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()
	updateRoom := suite.room1
	updateRoom.Number = suite.roomToSave.Number
	updateRoom.Type = entity.FamilyRoom

	updatedRoom, _ := suite.roomModel.Update(updateRoom)

	assert.Equal(suite.roomToSave.Number, updatedRoom.Number)
	assert.Equal(entity.FamilyRoom, updatedRoom.Type)
}

func TestSaveRoomNumberAlreadyInUse(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()
	saveRoom := suite.roomToSave
	saveRoom.Number = suite.room2.Number

	_, err := suite.roomModel.Save(saveRoom)

	if assert.NotNil(err) {
		assert.Equal(409, err.Status)
	}
}

func TestSaveRoomInvalidType(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()
	saveRoom := suite.roomToSave
	saveRoom.Type = entity.RoomType(100)

	_, err := suite.roomModel.Save(saveRoom)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestSaveRoomSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()
	saveRoom := suite.roomToSave

	savedRoom, _ := suite.roomModel.Save(saveRoom)

	assert.Equal(suite.roomToSave.ID, savedRoom.ID)
	assert.Equal(suite.roomToSave.Number, savedRoom.Number)
	assert.Equal(suite.roomToSave.Type, savedRoom.Type)
}

func TestDeleteRoomNotFound(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	err := suite.roomModel.Delete(suite.roomToSave.ID)

	if assert.NotNil(err) {
		assert.Equal(404, err.Status)
	}
}

func TestDeleteRoomSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupRoomSuite()

	err := suite.roomModel.Delete(suite.room1.ID)

	assert.Nil(err)
}
