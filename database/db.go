package database

import (
	"github.com/jinzhu/gorm"
)

// InitDatabase инициализирует подключение к БД (возвращает указатель на соединение и ошибку)
func InitDatabase(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
