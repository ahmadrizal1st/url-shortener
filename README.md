# URL Shortener - Development Documentation

A high-performance URL shortener service built with Go, PostgreSQL, and Redis.

## Development Setup

### Prerequisites

- **Go 1.19+** or later
- **Docker** and **Docker Compose** (recommended for containerized setup)
- **Git**

### Environment Setup

1. **Clone the repository**:
   
   ```bash
   git clone https://github.com/ahmadrizal1st/url-shortener.git
   cd url-shortener
   ```
2. **Switch to development branch**:
   
   ```bash
   git checkout develop
   ```
3. **Create environment file**:
   
   ```bash
   cp .env.example .env
   ```
   
   Edit the `.env` file with your configuration:
   
   ```
   DB_ADDRESS="localhost:6379"
   DB_PASSWORD=""
   APP_PORT="8000"
   DOMAIN="localhost:5173"
   API_QUOTA=10
   DATABASE_URL="postgres://user:password@localhost:5432/urlshortener?sslmode=disable"
   ```
4. **Install Go dependencies**:
   
   ```bash
   go mod download
   ```

### Running the Application

#### Option 1: Using Docker Compose (Recommended)

```bash
# Start all services (app, PostgreSQL, Redis)
docker-compose up --build

# Run in detached mode
docker-compose up -d --build
```

#### Option 2: Running Locally with External Services

1. Ensure PostgreSQL and Redis are running
2. Start the application:
   ```bash
   go run main.go
   ```

#### Option 3: Running Tests Only

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./api/routes/...
```

### Project Structure

```
url-shortener/
├── api/                    # API-related code
│   ├── database/          # Database connection and operations
│   │   └── database.go    # DB initialization and utilities
│   ├── models/            # Data models
│   │   └── models.go      # Struct definitions
│   ├── routes/            # HTTP route handlers
│   │   ├── addTag.go      # Add tag to URL endpoint
│   │   ├── deleteUrl.go   # Delete URL endpoint
│   │   ├── editUrl.go     # Edit URL endpoint
│   │   ├── getUrl.go      # Get URL endpoint
│   │   └── shorten.go     # URL shortening endpoint
│   └── utils/             # Utility functions
│       └── utils.go       # Helper functions
├── db/                    # Database-related files
├── Dockerfile            # Application container configuration
├── docker-compose.yml    # Multi-container setup
├── go.mod               # Go module dependencies
├── go.sum               # Go dependency checksums
├── main.go              # Application entry point
└── README.md            # Project documentation
```

### Key Dependencies

The project uses the following main Go modules:

- **Gin**: Web framework for HTTP routing
- **GORM**: ORM for database interactions with PostgreSQL
- **Redis Go**: Redis client for caching and rate limiting
- **Godotenv**: Environment variable loading
- **Validator**: Input validation

### Development Workflow

1. **Make changes** to the codebase
2. **Run tests** to ensure functionality:
   ```bash
   go test ./...
   ```
3. **Format code**:
   ```bash
   go fmt ./...
   ```
4. **Check for issues**:
   ```bash
   go vet ./...
   ```
5. **Build the application**:
   ```bash
   go build -o url-shortener main.go
   ```

### API Endpoints Overview

- `POST /api/v1` - Create a short URL
- `GET /api/v1/:shortID` - Retrieve original URL
- `PUT /api/v1/:shortID` - Update a short URL
- `DELETE /api/v1/:shortID` - Delete a short URL
- `POST /api/v1/tag` - Add tag to a URL

For detailed API documentation with request/response examples, refer to the [testing branch documentation](https://github.com/ahmadrizal1st/url-shortener/tree/testing/#readme).

### Common Development Tasks

#### Adding a New Endpoint

1. Create a new handler in `api/routes/`
2. Register the route in `main.go`
3. Add corresponding tests
4. Update documentation

#### Database Changes

1. Modify models in `api/models/models.go`
2. Update database operations in `api/database/database.go`
3. Consider database migrations if needed

### Debugging

#### View Logs with Docker

```bash
# View application logs
docker-compose logs app

# View all service logs
docker-compose logs -f

# View specific service logs
docker-compose logs redis
docker-compose logs db
```

#### Running with Debug Output

```bash
# Enable debug mode
DEBUG=true go run main.go
```

### Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Make your changes and test thoroughly
4. Ensure code is formatted: `go fmt ./...`
5. Run tests: `go test ./...`
6. Submit a pull request to the `develop` branch

### Getting Help

- Check the [testing branch documentation](https://github.com/ahmadrizal1st/url-shortener/tree/testing/#readme) for API details
- Review existing issues on GitHub
- Create a new issue for bugs or feature requests

### Useful Commands Reference

```bash
# Start development environment
docker-compose up --build

# Stop services
docker-compose down

# Run tests
go test ./...

# Format code
go fmt ./...

# Check for issues
go vet ./...

# Clean build
go clean && go build -o url-shortener main.go
```

This development documentation focuses on setting up and working with the codebase. For API usage and testing instructions, please refer to the testing branch documentation.

