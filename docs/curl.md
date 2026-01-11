# Curl Commands

Base URL: `http://localhost:5000/api`

## 1. Health Check

```bash
curl -X GET http://localhost:5000/api/
```

## 2. Get All Cards

```bash
curl -X GET http://localhost:5000/api/cards
```

## 3. Get Card By ID

```bash
curl -X GET http://localhost:5000/api/cards/1
```

## 4. Create New Card

```bash
curl -X POST http://localhost:5000/api/cards \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Go jogging",
    "description": "Morning exercise",
    "dueDate": "2026-01-12T07:00:00Z",
    "status": "todo"
  }'
```

## 5. Update Card (Partial Fields)

```bash
curl -X PUT http://localhost:5000/api/cards/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Buy groceries and snacks",
    "status": "in-progress"
  }'
```

## 6. Update Card (All Fields)

```bash
curl -X PUT http://localhost:5000/api/cards/1/all \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Go jogging daily",
    "description": "Morning exercise routine",
    "dueDate": "2026-01-12T07:00:00Z",
    "status": "done"
  }'
```

## 7. Delete Card

```bash
curl -X DELETE http://localhost:5000/api/cards/1
```