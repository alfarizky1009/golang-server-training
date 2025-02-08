package models

import "gorm.io/gorm"

type Author struct {
	Id        		*string `json:"id" gorm:"type:varchar(50);unique;primaryKey;autoIncrement;notNull"`
	Name    		*string `json:"name" gorm:"type:varchar(50)"`
	Gender     		*string `json:"gender" gorm:"type:varchar(50)"`
	Age 			*string `json:"age" gorm:"type:varchar(50)"`
	IsStillActive 	*bool `json:"is_still_active" gorm:"type:boolean"`
}

func MigrateAuthor(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}