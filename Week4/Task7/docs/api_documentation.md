# Task Manager API Documentation

Base URL: `http://localhost:8080`

## Configuration

### MongoDB Setup

1. Ensure you have MongoDB installed and running locally on port 27017.
2. The application connects to `mongodb://localhost:27017` by default. You can modify the connection string in `main.go` if needed.
3. The application uses a database named `task_manager` and collections named `tasks` and `users`.

## Authentication

This API uses JWT (JSON Web Tokens) for authentication.
To access protected routes, you must include the JWT token in the `Authorization` header.

**Header Format:**
`Authorization: Bearer <your_token>`

## Endpoints

### User Management

#### POST /register

Register a new user. The first registered user will be an admin.

**Request Body:**

```json
{
  "username": "user1",
  "password": "password123"
}
```

**Response** (201 Created):

```json
{
  "message": "User registered successfully"
}
```

#### POST /login

Authenticate a user and return a JWT token.

**Request Body:**

```json
{
  "username": "user1",
  "password": "password123"
}
```

**Response** (200 OK):

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### POST /promote (Admin Only)

Promote a user to admin role.

**Request Body:**

```json
{
  "username": "user2"
}
```

**Response** (200 OK):

```json
{
  "message": "User promoted successfully"
}
```

### Tasks

#### GET /tasks (Authenticated)

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

#### GET /tasks/:id (Authenticated)

Get a specific task by ID.

**Response** (200 OK):

```json
{
  "id": "...",
  "title": "...",
  "description": "...",
  "due_date": "2025-01-01T12:00:00Z",
  "status": "pending"
}
```

#### POST /tasks (Admin Only)

Create a new task.

**Request Body:**

```json
{
  "title": "New Task",
  "description": "Task description",
  "due_date": "2025-01-01T12:00:00Z",
  "status": "pending"
}
```

**Response** (201 Created):

```json
{
  "id": "...",
  "title": "New Task",
  "description": "Task description",
  "due_date": "2025-01-01T12:00:00Z",
  "status": "pending"
}
```

#### PUT /tasks/:id (Admin Only)

Update an existing task.

**Request Body:**

```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2025-01-01T12:00:00Z",
  "status": "completed"
}
```

**Response** (200 OK):

```json
{
  "id": "...",
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2025-01-01T12:00:00Z",
  "status": "completed"
}
```

#### DELETE /tasks/:id (Admin Only)

Delete a task.

**Response** (204 No Content)
