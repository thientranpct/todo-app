package userstorage

import "gorm.io/gorm"

type postgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(db *gorm.DB) *postgresStorage {
	return &postgresStorage{db: db}
}
