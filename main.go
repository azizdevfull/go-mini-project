package main

import (
	"go-tutorial/controllers"
	internal "go-tutorial/internal/database"
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := internal.InitDB()

	if db == nil {
		return
	}
	noteService := &services.NoteService{}
	noteService.InitNoteService(db)

	notesController := &controllers.NoteController{}
	notesController.InitController(*noteService)
	notesController.InitRoutes(router)

	authService := services.InitAuthService(db)
	authController := controllers.InitController(authService)
	authController.InitRoutes(router)

	router.Run()
}
