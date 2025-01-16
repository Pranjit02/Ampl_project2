package test

import (
	"Ampl_project2/database"
	"Ampl_project2/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the GORM DB
type MockDB struct {
	mock.Mock
}

func (mdb *MockDB) Create(value interface{}) {
	mdb.Called(value)
}

func (mdb *MockDB) Find(out interface{}, where ...interface{}) {
	mdb.Called(out, where)
}

func (mdb *MockDB) First(out interface{}, where ...interface{}) {
	mdb.Called(out, where)
}

func (mdb *MockDB) Save(value interface{}) {
	mdb.Called(value)
}

func (mdb *MockDB) Delete(value interface{}, where ...interface{}) {
	mdb.Called(value, where)
}

// Helper to create a test context
func createTestContext(method, path string) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(method, path, nil)
	return c, w
}

func TestCreateTask(t *testing.T) {
	mockDB := new(MockDB)
	task := database.Task{
		Title:       "Test Task",
		Description: "This is a test",
		Status:      "pending",
	}

	mockDB.On("Create", &task).Return(nil)

	c, w := createTestContext("POST", "/tasks")
	handlers.CreateTask(c)

	mockDB.AssertExpectations(t)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetTasks(t *testing.T) {
	mockDB := new(MockDB)
	tasks := []database.Task{
		{ID: 1, Title: "Task 1", Description: "Description 1", Status: "completed"},
		{ID: 2, Title: "Task 2", Description: "Description 2", Status: "in-progress"},
	}

	mockDB.On("Find", &tasks).Return(nil)

	c, w := createTestContext("GET", "/tasks")
	handlers.GetTasks(c)

	mockDB.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTask(t *testing.T) {
	mockDB := new(MockDB)
	task := database.Task{ID: 1, Title: "Updated Task", Description: "Updated Description", Status: "in-progress"}

	mockDB.On("First", &task, 1).Return(nil)
	mockDB.On("Save", &task).Return(nil)

	c, w := createTestContext("PUT", "/tasks/1")
	handlers.UpdateTask(c)

	mockDB.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask(t *testing.T) {
	mockDB := new(MockDB)
	task := database.Task{ID: 1}

	mockDB.On("First", &task, 1).Return(nil)
	mockDB.On("Delete", &task).Return(nil)

	c, w := createTestContext("DELETE", "/tasks/1")
	handlers.DeleteTask(c)

	mockDB.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}
