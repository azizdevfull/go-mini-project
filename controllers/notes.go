package controllers

import (
	"go-tutorial/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	noteService services.NoteService
}

func (n *NoteController) InitNoteControllerRouter(router *gin.Engine, noteService services.NoteService) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	n.noteService = noteService
}

func (n *NoteController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Query("status")
		actualStatus, err := strconv.ParseBool(status)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		data, err := n.noteService.GetNoteService(actualStatus)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"notes": data,
		})
	}
}
func (n *NoteController) CreateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}
	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.noteService.CreateNoteService(noteBody.Title, noteBody.Status)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"note": note,
		})

	}
}
