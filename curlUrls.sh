#!/bin/bash

BASE_URL="http://localhost:8080"

AUTH_HEADER="Authorization: Bearer mysecrettoken"

echo "GET /public/tasks"
curl -X GET "$BASE_URL/public/tasks"

# Protected Routes

# Create a new task (protected)
echo -e "\nPOST /tasks - Create a task"
curl -X POST "$BASE_URL/tasks" \
     -H "$AUTH_HEADER" \
     -H "Content-Type: application/json" \
     -d '{"title": "New Task", "description": "Task description", "status": "pending"}'

# Get a task by ID (protected)
echo -e "\nGET /tasks/{id} - Get a task by ID"
curl -X GET "$BASE_URL/tasks/1" \
     -H "$AUTH_HEADER"

# Update a task (protected)
echo -e "\nPUT /tasks/{id} - Update a task"
curl -X PUT "$BASE_URL/tasks/1" \
     -H "$AUTH_HEADER" \
     -H "Content-Type: application/json" \
     -d '{"title": "Updated Task", "description": "Updated description", "status": "in-progress"}'

# Delete a task (protected)
echo -e "\nDELETE /tasks/{id} - Delete a task"
curl -X DELETE "$BASE_URL/tasks/1" \
     -H "$AUTH_HEADER"

# Error Handling Example (Unauthorized Request)

echo -e "\nGET /tasks/1 (Unauthorized)"
curl -X GET "$BASE_URL/tasks/1"

# Rate Limiting Example

# Sending multiple requests to trigger rate limit
echo -e "\nGET /tasks/{id} - Test Rate Limiting"
for i in {1..6}; do
    curl -X GET "$BASE_URL/tasks/1" -H "$AUTH_HEADER"
done
