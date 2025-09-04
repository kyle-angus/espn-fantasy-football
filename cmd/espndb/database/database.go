package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Instance() *gorm.DB {
	if DB == nil {
		Connect()
	}

	return DB
}

func Connect() {
	dsn := "root:L0calDB@tcp(10.3.0.10:3306)/fantasy_football"
	// dsn := "host=10.3.0.10 user=root password=L0calDB dbname=fantasy_football port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	DB = db
}
