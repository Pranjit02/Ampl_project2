package handlers

import (
	"Ampl_project2/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []database.Task
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	database.DB.Offset(offset).Limit(limit).Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	var task database.Task
	if err := database.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var task database.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	var task database.Task
	if err := database.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}
func DeleteTask(c *gin.Context) {
	var task database.Task
	if err := database.DB.First(&task, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
