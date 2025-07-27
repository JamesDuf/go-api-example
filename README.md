# Go API Example

A simple REST API built with Go that demonstrates authentication middleware and coin balance retrieval functionality.

## Features

- RESTful API with Chi router
- Token-based authentication middleware
- Mock database implementation
- Structured error handling
- JSON responses
- Logging with Logrus

## Prerequisites

- Go 1.20 or higher
- Git

## Installation

1. Clone the repository:

```bash
git clone https://github.com/JamesDuf/go-api-example.git
cd go-api-example
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the application:

```bash
go run cmd/api/main.go
```

The server will start on `localhost:8080`.

## API Documentation

### Authentication

All endpoints under `/account` require authentication via header token.

**Required Headers:**

- `Authorization`: User's auth token

**Required Query Parameters:**

- `username`: Username for authentication

### Endpoints

#### Get Coin Balance

Get the current coin balance for a user.

```http
GET /account/coins?username={username}
```

**Headers:**

```http
Authorization: {auth_token}
```

**Response:**

```json
{
  "Code": 200,
  "Balance": 1000
}
```

**Error Response:**

```json
{
  "Code": 400,
  "Message": "invalid credentials"
}
```

## Testing the API

### Using cURL

1. **Valid Request:**

```bash
curl -H "Authorization: 123ABC" "http://localhost:8080/account/coins?username=alex"
```

2. **Invalid Token:**

```bash
curl -H "Authorization: invalid" "http://localhost:8080/account/coins?username=alex"
```

3. **Missing Username:**

```bash
curl -H "Authorization: 123ABC" "http://localhost:8080/account/coins"
```

### Mock Users

The application includes three mock users for testing:

| Username | Auth Token | Coin Balance |
|----------|------------|--------------|
| alex     | 123ABC     | 1000         |
| jane     | 456DEF     | 2000         |
| marie    | 789GHI     | 3000         |

## Project Structure

```text
.
├── README.md
├── go.mod
├── go.sum
├── api/
│   └── api.go              # API types and error handlers
├── cmd/
│   └── api/
│       └── main.go         # Application entry point
└── internal/
    ├── handlers/
    │   ├── api.go          # Route definitions
    │   └── get_coin_balance.go  # Coin balance handler
    ├── middleware/
    │   └── authorization.go     # Authentication middleware
    └── tools/
        ├── database.go     # Database interface
        └── mockdb.go       # Mock database implementation
```

## Dependencies

- **[Chi](https://github.com/go-chi/chi)**: Lightweight, idiomatic HTTP router
- **[Logrus](https://github.com/sirupsen/logrus)**: Structured logger
- **[Gorilla Schema](https://github.com/gorilla/schema)**: URL query parameter decoder

## Security Features

- Token-based authentication
- Input validation
- Structured error responses (no sensitive data exposure)
- Authorization middleware for protected routes

## Development

### Adding New Endpoints

1. Define your handler function in `internal/handlers/`
2. Add the route in `internal/handlers/api.go`
3. Add any required types in `api/api.go`

### Database Integration

Currently uses a mock database. To integrate with a real database:

1. Implement the `DatabaseInterface` in `internal/tools/database.go`
2. Update the `NewDatabase()` function to return your implementation
3. Configure connection settings

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

This project was built following a tutorial by Alex Mux ([link-to-tutorial](https://youtu.be/8uiZC0l4Ajw)). Special thanks for the clear explanations and guidance on building REST APIs with Go.
