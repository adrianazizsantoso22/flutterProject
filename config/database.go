package config

import (
	"fmt"
	"go-notes-taker/entity"
	"go-notes-taker/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	dbUser := helpers.MustGetenv("DB_USER")
	dbPass := helpers.MustGetenv("DB_PASS")
	dbHost := helpers.MustGetenv("DB_HOST")
	dbName := helpers.MustGetenv("DB_NAME")
	dbPort := helpers.MustGetenvInt("DB_PORT")

	fmt.Println(dbUser)
	fmt.Println(dbPass)
	fmt.Println(dbHost)
	fmt.Println(dbName)
	fmt.Println(dbPort)

	dsn := fmt.Sprintf("TimeZone=Asia/Jakarta host=%v user=%v password=%v dbname=%v port=%v", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := db.AutoMigrate(
		entity.User{},
	); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}
