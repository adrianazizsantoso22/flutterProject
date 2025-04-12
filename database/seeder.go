package database

import (
	"gorm.io/gorm"
	"go-notes-taker/database/seeders"
)

func Seeder(db *gorm.DB) error {
	if err := seeders.UserSeeder(db); err != nil {
		return err
	}

	if err := seeders.NoteSeeder(db); err != nil {
		return err
	}

	return nil
}

