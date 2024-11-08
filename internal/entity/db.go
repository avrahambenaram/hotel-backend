package entity

import (
	"github.com/avrahambenaram/hotel-backend/internal/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	db, err := gorm.Open(mysql.Open(configuration.MysqlDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Client{})
	db.AutoMigrate(&HotelRoom{})
	db.AutoMigrate(&Reservation{})
	DB = db
}
