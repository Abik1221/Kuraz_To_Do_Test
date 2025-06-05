package controllers

import (
	"net/http"
	"strconv"

	"github.com/abik1221/Kuraz_To_Do_Test/config"
	"github.com/abik1221/Kuraz_To_Do_Test/models"

	"github.com/gin-gonic/gin"
)

// GET /api/tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	result := config.DB.Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// POST /api/tasks
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// PUT /api/tasks/:id
func MarkTaskCompleted(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	task.Completed = true
	config.DB.Save(&task)

	c.JSON(http.StatusOK, task)
}

// DELETE /api/tasks/:id
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	result := config.DB.Delete(&models.Task{}, idUint)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
