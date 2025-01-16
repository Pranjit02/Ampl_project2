package test

import (
	"Ampl_project2/database"
	"Ampl_project2/handlers"
	"Ampl_project2/middleware"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func setupRouter() *gin.Engine {
	database.InitDB()
	r := gin.Default()

	public := r.Group("/public")
	{
		public.GET("/tasks", handlers.GetTasks)
	}

	protected := r.Group("/tasks")
	protected.Use(middleware.AuthMiddleware())
	protected.Use(middleware.RateLimiterMiddleware())
	{
		protected.GET(":id", handlers.GetTaskByID)
		protected.POST("", handlers.CreateTask)
		protected.PUT(":id", handlers.UpdateTask)
		protected.DELETE(":id", handlers.DeleteTask)
	}

	return r
}

func TestGetTasksPublic(t *testing.T) {
	r := setupRouter()
	database.DB.Create(&database.Task{Title: "Task 1", Description: "Description 1", Status: "pending"})
	database.DB.Create(&database.Task{Title: "Task 2", Description: "Description 2", Status: "completed"})

	w := performRequest(r, "GET", "/public/tasks", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task 1")
	assert.Contains(t, w.Body.String(), "Task 2")
}

func TestCreateTaskProtected(t *testing.T) {
	r := setupRouter()
	taskData := `{"title":"New Task","description":"New Description","status":"pending"}`
	w := performRequest(r, "POST", "/tasks", []byte(taskData))

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New Task")
}

func TestRateLimiting(t *testing.T) {
	r := setupRouter()
	for i := 0; i < 6; i++ {
		w := performRequest(r, "GET", "/tasks/1", nil)
		if i < 5 {
			assert.Equal(t, http.StatusOK, w.Code)
		} else {
			assert.Equal(t, http.StatusTooManyRequests, w.Code)
		}
	}
}

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer mysecrettoken")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
