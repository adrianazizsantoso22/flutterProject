package database

import (
	"fmt"
	"go-notes-taker/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		entity.User{},
		entity.Note{},
	); err != nil {
		return err
	}
	fmt.Println("Migration success!")

	return nil
}
