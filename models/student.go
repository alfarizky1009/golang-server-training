package models

import "gorm.io/gorm"

type Student struct {
	Id        	*string `json:"id" gorm:"type:varchar(50);unique;primaryKey;autoIncrement;notNull"`
	Name    	*string `json:"name" gorm:"type:varchar(50)"`
	Gender     	*string `json:"gender" gorm:"type:varchar(50)"`
	Major 		*string `json:"major" gorm:"type:varchar(50)"`
	Birthplace 	*string `json:"birthplace" gorm:"type:varchar(50)"`
	Birthdate 	*string `json:"birthdate" gorm:"type:varchar(50)"`
}

func MigrateStudent(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}