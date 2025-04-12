package routes

import (
	"go-notes-taker/controller"
	"go-notes-taker/middleware"
	"go-notes-taker/service"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(router *gin.Engine, noteController controller.NoteController, jwtService service.JWTService) {
	noteRoutes := router.Group("/api/notes")
	{
		noteRoutes.POST("", noteController.CreateNote)
		noteRoutes.GET("", middleware.Authenticate(jwtService), noteController.GetAllNotes)
		noteRoutes.GET("/my", middleware.Authenticate(jwtService), noteController.GetMyNotes)
		noteRoutes.DELETE("/:note_id", middleware.Authenticate(jwtService), noteController.DeleteNote)
		noteRoutes.PUT("/:note_id", middleware.Authenticate(jwtService), noteController.UpdateNote)
	}
}