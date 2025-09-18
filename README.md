# URL Shortener

A simple URL shortener built with Golang and PostgreSQL. It stores original URLs and generates short links for easy sharing.

## Development Documentation

This documentation focuses on development setup, running, and testing the URL shortener application.

### Prerequisites

- Go 1.25.1 or later
- Docker and Docker Compose (for containerized setup)
- PostgreSQL (if running locally without Docker)
- Git

### Environment Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/Anurag/url-shortner.git
   cd url-shortner
   ```

2. Create a `.env` file in the root directory with the following variables:
   ```
   APP_PORT=8000
   DATABASE_URL=postgres://user:password@localhost:5432/urlshortener?sslmode=disable
   REDIS_URL=redis://localhost:6379
   ```
   Adjust the values according to your local setup.

3. Install Go dependencies:
   ```bash
   go mod download
   ```

### Running Locally

1. Ensure PostgreSQL and Redis are running locally, or use Docker Compose for services.

2. Run the application:
   ```bash
   go run main.go
   ```

3. The server will start on port 8000 (or the port specified in `.env`).

### Running with Docker

1. Use Docker Compose to run the entire stack:
   ```bash
   docker-compose up --build
   ```

2. The application will be available at `http://localhost:8000`.

### Testing

For testing instructions, refer to the testing branch README:
[https://github.com/ahmadrizal1st/url-shortener/tree/testing/#readme](https://github.com/ahmadrizal1st/url-shortener/tree/testing/#readme)

### Project Structure

- `main.go`: Entry point of the application, sets up the server and routes.
- `api/`: Contains API-related code.
  - `database/`: Database connection and initialization.
  - `models/`: Data models.
  - `routes/`: API route handlers.
  - `utils/`: Utility functions.
- `db/`: Database-related Docker files.
- `Dockerfile`: Docker configuration for the application.
- `docker-compose.yml`: Docker Compose setup for the full stack.

### Dependencies

The project uses the following Go modules (from `go.mod`):
- Gin (web framework)
- GORM (ORM for database interactions)
- Redis (caching)
- Godotenv (environment variable loading)
- And others for validation, JSON handling, etc.

### API Endpoints

- `POST /api/v1`: Shorten a URL
- `GET /api/v1/:shortID`: Retrieve original URL by short ID
- `PUT /api/v1/:shortID`: Edit a shortened URL
- `DELETE /api/v1/:shortID`: Delete a shortened URL
- `POST /api/v1/tag`: Add a tag to a URL

### Useful Commands

- Run tests: `go test ./...`
- Format code: `go fmt ./...`
- Lint code: `go vet ./...`
- Build: `go build -o url-shortener main.go`

### Contributing

1. Fork the repository.
2. Create a feature branch.
3. Make your changes.
4. Run tests and ensure code is formatted.
5. Submit a pull request.

For more details, refer to the testing branch documentation.
