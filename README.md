# microservices-in-go

This service provides a basic API for managing users using Go, `chi` for routing, and Cassandra as a database.

## Endpoints

### 1. Save a User

**Endpoint:** `POST /SaveUserById`

**Description:** Saves a user with name, email, and ID.

**Request Body (JSON):**
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "id": "1234"
}

curl -X POST http://localhost:80/SaveUserById \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john.doe@example.com", "id": "1234"}'


curl -X GET http://localhost:80/getUserById/1234


docker run --name cassandra -p 9042:9042 -d cassandra:latest


docker exec -it cassandra cqlsh

docker-compose up --build
