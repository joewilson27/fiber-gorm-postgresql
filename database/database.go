package database

import (
	"log"

	"gorm-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// connectDb
func ConnectDb() {
	dsn := "host=localhost user=postgres password='Aiforesee123' dbname=gorm-postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migration")
	db.AutoMigrate(&models.Book{})

	DB = Dbinstance{
		Db: db,
	}

}