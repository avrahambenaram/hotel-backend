package model_test

import (
	"testing"
	"time"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/model"
	inmemory "github.com/avrahambenaram/hotel-backend/internal/repository/implementation/in-memory"
	"github.com/stretchr/testify/assert"
)

type ReservationSuite struct {
	reservation1      entity.Reservation
	reservation2      entity.Reservation
	reservationToSave entity.Reservation

	reservationModel *model.ReservationModel
}

func SetupReservationSuite() *ReservationSuite {
	reservationRepository := new(inmemory.ReservationRepository)
	reservationModel := model.NewReservationModel(reservationRepository)

	checkIn1, _ := time.Parse(time.RFC3339, "2024-01-10T20:00:00Z")
	checkOut1, _ := time.Parse(time.RFC3339, "2024-01-14T20:00:00Z")
	reservation1 := entity.Reservation{
		ID:       0,
		RoomID:   0,
		ClientID: 0,
		CheckIn:  checkIn1,
		CheckOut: checkOut1,
	}

	checkIn2, _ := time.Parse(time.RFC3339, "2024-01-20T20:00:00Z")
	checkOut2, _ := time.Parse(time.RFC3339, "2024-01-24T20:00:00Z")
	reservation2 := entity.Reservation{
		ID:       1,
		RoomID:   1,
		ClientID: 0,
		CheckIn:  checkIn2,
		CheckOut: checkOut2,
	}

	checkInToSave, _ := time.Parse(time.RFC3339, "2024-01-10T20:00:00Z")
	checkOutToSave, _ := time.Parse(time.RFC3339, "2024-01-14T20:00:00Z")
	reservationToSave := entity.Reservation{
		ID:       2,
		RoomID:   1,
		ClientID: 1,
		CheckIn:  checkInToSave,
		CheckOut: checkOutToSave,
	}

	reservationRepository.Save(reservation1)
	reservationRepository.Save(reservation2)

	return &ReservationSuite{
		reservation1,
		reservation2,
		reservationToSave,
		reservationModel,
	}
}

func TestGetAllReservations(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()

	reservations := suite.reservationModel.FindAll()

	assert.Len(reservations, 2)
}

func TestGetReservationByIDNotFound(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()

	_, err := suite.reservationModel.FindByID(suite.reservationToSave.ID)

	if assert.NotNil(err) {
		assert.Equal(404, err.Status)
	}
}

func TestGetReservationByIDSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()

	reservation, _ := suite.reservationModel.FindByID(suite.reservation1.ID)

	assert.Equal(suite.reservation1.ID, reservation.ID)
	assert.Equal(suite.reservation1.ClientID, reservation.ClientID)
	assert.Equal(suite.reservation1.RoomID, reservation.RoomID)
}

func TestGetReservationsByClientAndRoom(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()

	reservations := suite.reservationModel.FindByClientAndRoom(0, 1)

	assert.Len(reservations, 1)
}

func TestGetReservationsByClient(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()

	reservations := suite.reservationModel.FindByClient(0)

	assert.Len(reservations, 2)
}

func TestGetReservationsByRoom(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()

	reservations := suite.reservationModel.FindByRoom(1)

	assert.Len(reservations, 1)
}

func TestSaveReservationAlreadyReserved(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()
	save := suite.reservationToSave
	save.CheckOut = suite.reservation2.CheckIn.Add(time.Hour * 24)

	_, err := suite.reservationModel.Save(save)

	if assert.NotNil(err) {
		assert.Equal(409, err.Status)
	}
}

func TestSaveReservationCheckOutBeforeCheckIn(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()
	save := suite.reservationToSave
	save.CheckOut = save.CheckIn.Add(time.Hour * -24)

	_, err := suite.reservationModel.Save(save)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestSaveReservationSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()
	save := suite.reservationToSave

	reservation, _ := suite.reservationModel.Save(save)

	assert.Equal(suite.reservationToSave.ID, reservation.ID)
	assert.Equal(suite.reservationToSave.ClientID, reservation.ClientID)
	assert.Equal(suite.reservationToSave.RoomID, reservation.RoomID)
}

func TestDeleteReservationNotFound(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()

	err := suite.reservationModel.Delete(10)

	if assert.NotNil(err) {
		assert.Equal(404, err.Status)
	}
}

func TestDeleteReservationSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupReservationSuite()

	err := suite.reservationModel.Delete(suite.reservation1.ID)

	assert.Nil(err)
}
