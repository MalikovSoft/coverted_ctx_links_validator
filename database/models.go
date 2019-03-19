package database

import (
	"github.com/jinzhu/gorm"
)

// LinkToResource модель ссылки на ресурс (содержит старую и новую ссылки)
type LinkToResource struct {
	gorm.Model
	OldLink string `gorm:"column:old_link;type:varchar(255);not null;unique" json:"old_link"`
	NewLink string `gorm:"column:new_link;type:varchar(255)" json:"new_link"`
}

// TableName возвращает имя таблицы ссылок
func (LinkToResource) TableName() string {
	return "convert_links"
}
