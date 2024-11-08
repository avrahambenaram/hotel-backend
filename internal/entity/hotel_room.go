package entity

type RoomType uint

const (
	SimpleRoom RoomType = 0
	DoubleRoom RoomType = 1
	Suite      RoomType = 4
	FamilyRoom RoomType = 5
)

func (c RoomType) IsValid() bool {
	if c == SimpleRoom {
		return true
	}
	if c == DoubleRoom {
		return true
	}
	if c == Suite {
		return true
	}
	if c == FamilyRoom {
		return true
	}
	return false
}

type HotelRoom struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Number     int      `json:"number" gorm:"not null"`
	Type       RoomType `json:"type" gorm:"not null"`
	Capacity   uint     `json:"capacity" gorm:"not null"`
	PriceDiary float32  `json:"priceDiary" gorm:"not null"`
}
