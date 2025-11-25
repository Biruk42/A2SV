# Task Manager API Documentation

Base URL: `http://localhost:8080`

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
