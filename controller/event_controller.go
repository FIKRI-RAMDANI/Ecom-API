package controller

import (
	"event-app/config"
	"event-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// dummy user id
	event.UserId = 1

	config.DB.Create(&event)
	context.JSON(http.StatusCreated, gin.H{
		"message": "Data berhasil dibuat",
		"event":   event,
	})
}

func GetEvents(context *gin.Context) {
	var events []models.Event

	config.DB.Find(&events)
	context.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil tampilkan semua event",
		"event":   events,
	})
}

func GetEventById(context *gin.Context) {
	var event models.Event
	paramsID := context.Param("id")

	var eventData = config.DB.First(&event, paramsID).Error
	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event Tidak ditemukan",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil tampilkan semua event",
		"event":   event,
	})
}

func UpdateEvent(context *gin.Context) {
	var event models.Event
	paramsID := context.Param("id")

	var eventData = config.DB.First(&event, paramsID).Error
	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event Tidak ditemukan",
		})
		return
	}

	var input models.Event
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Model(&event).Updates(input)
	context.JSON(http.StatusOK, gin.H{
		"message": "Data Behasil di Update",
		"event":   event,
	})

}

func DeleteEvent(context *gin.Context) {
	var event models.Event
	paramsID := context.Param("id")

	var eventData = config.DB.Delete(&event, paramsID).Error
	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event Tidak ditemukan",
		})
		return
	}

	config.DB.Unscoped().Delete(&event)
	context.JSON(http.StatusOK, gin.H{
		"message": "Data Behasil di Delete",
	})
}
