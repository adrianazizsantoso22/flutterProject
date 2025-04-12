package seeders

import (
	"encoding/json"
	"go-notes-taker/entity"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

func NoteSeeder(db *gorm.DB) error {
	file, err := os.Open("database/seeders/data/note.json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return err
	}

	var notes []entity.Note
	err = json.Unmarshal(data, &notes)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, note := range notes {
		if err := db.Create(&note).Error; err != nil {
			return err
		}
	}

	return nil
}
