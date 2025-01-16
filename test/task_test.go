package task_test

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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

func TestCreateTask(t *testing.T) {
	mockDB := new(MockDB)
	task := Task{Title: "Test Task", Description: "This is a test", Status: "pending", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	mockDB.On("Create", &task).Return(nil)

	// Call the CreateTask method with the mockDB
	// Replace the db variable in the real app with the mockDB

	c := new(gin.Context)
	createTask(c)

	mockDB.AssertExpectations(t)
}

func TestGetTasks(t *testing.T) {
	mockDB := new(MockDB)
	tasks := []Task{
		{ID: 1, Title: "Task 1", Description: "Description 1", Status: "completed"},
		{ID: 2, Title: "Task 2", Description: "Description 2", Status: "in-progress"},
	}

	mockDB.On("Find", &tasks).Return(nil)

	// Call the GetTasks method with the mockDB
	c := new(gin.Context)
	getTasks(c)

	mockDB.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockDB := new(MockDB)
	task := Task{ID: 1, Title: "Updated Task", Description: "Updated Description", Status: "in-progress"}

	mockDB.On("First", &task, 1).Return(nil)
	mockDB.On("Save", &task).Return(nil)

	// Call the UpdateTask method with the mockDB
	c := new(gin.Context)
	updateTask(c)

	mockDB.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockDB := new(MockDB)
	task := Task{ID: 1}

	mockDB.On("First", &task, 1).Return(nil)
	mockDB.On("Delete", &task).Return(nil)

	// Call the DeleteTask method with the mockDB
	c := new(gin.Context)
	deleteTask(c)

	mockDB.AssertExpectations(t)
}
