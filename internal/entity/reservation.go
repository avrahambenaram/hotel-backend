package entity

import "time"

type Reservation struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	CheckIn  time.Time `json:"checkIn" gorm:"not null"`
	CheckOut time.Time `json:"checkOut" gorm:"not null"`
	RoomID   uint      `json:"roomId" gorm:"not null"`
	ClientID uint      `json:"clientId" gorm:"not null"`

	Room   *HotelRoom `json:"room" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Client *Client    `json:"client" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
