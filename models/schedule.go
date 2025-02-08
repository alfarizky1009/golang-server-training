package models

import "gorm.io/gorm"

type Schedule struct {
	Id        	*string `json:"id" gorm:"type:varchar(50);unique;primaryKey;autoIncrement;notNull"`
	StudentId   *string `json:"student_id" gorm:"type:varchar(50)"`
	BookId		*string `json:"book_id" gorm:"type:varchar(50)"`
	BorrowDate 	*string `json:"borrow_date" gorm:"type:varchar(50)"`
	ReturnDate 	*string `json:"return_date" gorm:"type:varchar(50)"`
}

func MigrateSchedule(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}