# Cards API Documentation

**Base URL**: `http://localhost:5000/api`

---

## 1. Health Check

Check if the API is running.

- **Endpoint**: `/api/`
- **Method**: `GET`

### Responses

**200 OK**
```json
{
  "status": "ok"
}
```

---

## 2. Get All Cards

Fetch all cards.

- **Endpoint**: `/api/cards`
- **Method**: `GET`

### Responses

**200 OK**
```json
[
  {
    "id": 1,
    "title": "Buy groceries",
    "description": "Milk, Eggs, Bread",
    "dueDate": "2026-01-15T00:00:00Z",
    "status": "todo"
  },
  {
    "id": 2,
    "title": "Complete project",
    "description": "Finish API documentation",
    "dueDate": "2026-01-20T00:00:00Z",
    "status": "in-progress"
  }
]
```

**500 Internal Server Error**
```json
{
  "error": "Failed to fetch card"
}
```

---

## 3. Get Card By ID

Fetch a card by ID.

- **Endpoint**: `/api/cards/:id`
- **Method**: `GET`

### Responses

**200 OK**
```json
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Milk, Eggs, Bread",
  "dueDate": "2026-01-15T00:00:00Z",
  "status": "todo"
}
```

**404 Not Found**
```json
{
  "error": "Card not found"
}
```

---

## 4. Create New Card

Create a new card.

- **Endpoint**: `/api/cards`
- **Method**: `POST`

### Request Body
```json
{
  "title": "Go jogging",
  "description": "Morning exercise",
  "dueDate": "2026-01-12T07:00:00Z",
  "status": "todo"
}
```

### Responses

**201 Created**
```json
{
  "id": 3,
  "title": "Go jogging",
  "description": "Morning exercise",
  "dueDate": "2026-01-12T07:00:00Z",
  "status": "todo"
}
```

**400 Bad Request**
```json
{
  "message": "invalid request body"
}
```

**500 Internal Server Error**
```json
{
  "message": "failed to create card"
}
```

---

## 5. Update Card (Partial Fields)

Update partial fields of a card.

- **Endpoint**: `/api/cards/:id`
- **Method**: `PUT`

### Request Body
```json
{
  "title": "Buy groceries and snacks",
  "status": "in-progress"
}
```

### Responses

**200 OK**
```json
{
  "message": "card update successfully"
}
```

**400 Bad Request**
```json
{
  "message": "invalid request body"
}
```

**500 Internal Server Error**
```json
{
  "message": "failed to update card"
}
```

---

## 6. Update Card (All Fields)

Update all fields of a card.

- **Endpoint**: `/api/cards/:id/all`
- **Method**: `PUT`

### Request Body
```json
{
  "title": "Go jogging daily",
  "description": "Morning exercise routine",
  "dueDate": "2026-01-12T07:00:00Z",
  "status": "done"
}
```

### Responses

**200 OK**
```json
{
  "message": "card update successfully"
}
```

**404 Not Found**
```json
{
  "error": "Card not found"
}
```

**400 Bad Request**
```json
{
  "error": "invalid request body"
}
```

**500 Internal Server Error**
```json
{
  "error": "failed to update card"
}
```

---

## 7. Delete Card

Delete a card by ID.

- **Endpoint**: `/api/cards/:id`
- **Method**: `DELETE`

### Responses

**200 OK**
```json
{
  "message": "Card deleted successfully"
}
```

**404 Not Found**
```json
{
  "error": "Card not found"
}
```

**500 Internal Server Error**
```json
{
  "error": "Failed to delete card"
}
```
