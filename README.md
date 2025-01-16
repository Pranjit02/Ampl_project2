# Go Task Management API

This project implements a RESTful API for managing tasks using Go, Gin framework, and MySQL. 

## Features

* **Create Tasks:** Create new tasks with descriptions, due dates, and priorities.
* **Read Tasks:** Retrieve a list of tasks, or get details of a specific task.
* **Update Tasks:** Modify existing tasks (e.g., update status, change due date).
* **Delete Tasks:** Remove tasks.
* **Authentication:** Basic API key authentication for protected endpoints.
* **Rate Limiting:** Prevent abuse by limiting the number of requests per IP address.

## Technologies Used

* **Go:** Programming language.
* **Gin:** HTTP web framework for Go.
* **MySQL:** Relational database.
* **GORM:** Object-Relational Mapper (ORM) for Go.

## Getting Started

1. **Clone the repository:**

   git clone https://github.com/Pranjit02/Ampl_project2.git
   cd Ampl_project2

Install dependencies:
go mod tidy
Set up environment variables:

Create a .env file (or use an environment variable manager) to store sensitive information like:
DATABASE_URL: Connection string to your MySQL database (e.g., user:password@tcp(host:port)/database_name)

/tasks (GET): Retrieve a list of tasks.
/tasks/{id} (GET): Get details of a specific task.
/tasks (POST): Create a new task.
/tasks/{id} (PUT): Update an existing task.
/tasks/{id} (DELETE): Delete a task.
Authentication:

Protected endpoints require an Authorization header with the value Bearer <API_KEY>.
Rate Limiting:

Rate limits are enforced to prevent abuse.
TODO: Implement and configure rate limiting middleware (e.g., using gin-contrib/ratelimit).
Testing
Unit tests: Cover individual functions and components.

Integration tests: Verify the interaction between different parts of the system (e.g., API endpoints and database).

TODO: Implement unit and integration tests.
#git config --global core.autocrlf true
