package Api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func setupRouter() *gin.Engine {
	r := gin.Default()
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to MySQL database")
	}
	// Create table for Task model
	db.AutoMigrate(&db.Task{})
	globalDB = db
	// Define routes
	public := r.Group("/public")
	{
		public.GET("/tasks", getTasks)
	}
	protected := r.Group("/tasks")
	protected.Use(authMiddleware())
	protected.Use(rateLimiterMiddleware())
	{
		protected.GET(":id", getTaskByID)
		protected.POST("", createTask)
		protected.PUT(":id", updateTask)
		protected.DELETE(":id", deleteTask)
	}
	return r
}

func TestGetTasksPublic(t *testing.T) {
	r := setupRouter()
	rdb := globalDB
	rdb.Create(&Task{Title: "Task 1", Description: "Description 1", Status: "pending"})
	rdb.Create(&Task{Title: "Task 2", Description: "Description 2", Status: "completed"})

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
