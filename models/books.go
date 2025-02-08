package models

import "gorm.io/gorm"

// Golang can't understand json response 3rd string is to point the json field
type Books struct {
	Id        *string `json:"id" gorm:"type:varchar(50);unique;primaryKey;autoIncrement;notNull"`
	Author    *string `json:"author" gorm:"type:varchar(50)"`
	Title     *string `json:"title" gorm:"type:varchar(50)"`
	Publisher *string `json:"publisher" gorm:"type:varchar(50)"`
}

// function for migrating
func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}