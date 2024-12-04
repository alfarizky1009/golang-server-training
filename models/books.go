package models

import "gorm.io/gorm"

// Golang can't understand json response 3rd string is to point the json field
type Books struct {
	Id        uint    `json:"id" gorm:"unique;primaryKey;autoIncrement;notNull"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

// function for migrating
func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}