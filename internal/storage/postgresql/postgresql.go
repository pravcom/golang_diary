package postgresql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func New(storagePath string) *Storage {
	db, err := gorm.Open(postgres.Open(storagePath), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	return &Storage{DB: db}
}

func (s *Storage) Close() {
	sqlDB, err := s.DB.DB()
	if err != nil {
		log.Fatal("Failed to close connection to DB", err)
	}

	_ = sqlDB.Close()
}
