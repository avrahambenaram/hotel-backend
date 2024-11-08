package entity

import "time"

type Reservation struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	CheckIn  time.Time `json:"checkIn" gorm:"not null"`
	CheckOut time.Time `json:"checkOut" gorm:"not null"`
	RoomID   uint
	ClientID uint

	Room   HotelRoom `json:"room"`
	Client Client    `json:"client"`
}
