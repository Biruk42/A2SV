# Task Manager API Documentation

Base URL: `http://localhost:8080`

## Configuration

### MongoDB Setup

1. Ensure you have MongoDB installed and running locally on port 27017.
2. The application connects to `mongodb://localhost:27017` by default. You can modify the connection string in `main.go` if needed.
3. The application uses a database named `task_manager` and a collection named `tasks`.

## Endpoints

### GET /tasks

Return a list of all tasks.

**Response** (200 OK):

```json
{
  "tasks": [
    {
      "id": "...",
      "title": "...",
      "description": "...",
      "due_date": "2025-01-01T12:00:00Z",
      "status": "pending"
    }
  ]
}
```
