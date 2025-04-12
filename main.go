package main

import (
	// "go-notes-taker/common"
	"go-notes-taker/config"
	"go-notes-taker/controller"
	"go-notes-taker/repository"
	"go-notes-taker/routes"
	"go-notes-taker/service"
	"go-notes-taker/database"
	// "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	// "gorm.io/gorm"

	"log"
	"os"
)

func main() {
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	var (	
		jwtService 			service.JWTService 			= service.NewJWTService()

		userRepository 		repository.UserRepository 	= repository.NewUserRepository(db)
		noteRepository		repository.NoteRepository	= repository.NewNoteRepository(db)

		userService 		service.UserService 		= service.NewUserService(userRepository)
		noteService			service.NoteService			= service.NewNoteService(noteRepository)

		userController 		controller.UserController 	= controller.NewUserController(userService, jwtService)
		noteController		controller.NoteController	= controller.NewNoteController(noteService, jwtService)
	)

	if err := database.Migrate(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	if err := database.Seeder(db); err != nil {
		log.Fatalf("Error seeding database: %v", err)
	}

	server := gin.Default()
	routes.UserRoutes(server, userController, jwtService)
	routes.NoteRoutes(server, noteController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Run("127.0.0.1:" + port)
}