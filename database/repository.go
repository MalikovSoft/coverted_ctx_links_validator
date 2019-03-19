package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// GetAllLinksToResources возвращает коллекцию всех ссылок на ресурсы из таблицы в БД
func GetAllLinksToResources(db *gorm.DB) ([]*LinkToResource, map[string]string) {
	allLinksToResources := make([]*LinkToResource, 0)
	mapOfLinks := make(map[string]string)
	db.AutoMigrate(&LinkToResource{})
	db.Find(&allLinksToResources)
	for _, link := range allLinksToResources {
		mapOfLinks[link.OldLink] = link.NewLink
	}
	return allLinksToResources, mapOfLinks
}

// PutLinkRecordToDatabase добавляет запись ссылки в БД
func PutLinkRecordToDatabase(db *gorm.DB, link *LinkToResource) {
	db.AutoMigrate(&LinkToResource{})
	db.NewRecord(link)
	db.Create(&link)
	if !db.NewRecord(link) {
		fmt.Printf("New record was created\n")
	}
}

// UpdateLinkRecordFromDatabase изменяет запись ссылки в БД
func UpdateLinkRecordFromDatabase(db *gorm.DB, link *LinkToResource) {
	db.AutoMigrate(&LinkToResource{})
	db.Where("old_link = ?", link.OldLink).Save(&link)
}
