# URL Shortener

A high-performance URL shortener service built with Go, PostgreSQL, and Redis, designed to efficiently shorten long URLs and provide analytics.

## Features

- **URL Shortening**: Convert long URLs into short, manageable links
- **Custom Short Codes**: Option to specify custom short codes
- **Expiration Control**: Set expiration time for shortened URLs
- **Rate Limiting**: Built-in rate limiting to prevent abuse
- **Tagging System**: Add tags to organize and categorize shortened URLs
- **High Performance**: Utilizes Redis for caching and fast redirects
- **RESTful API**: Clean and simple API design
- **Docker Support**: Easy deployment with Docker and Docker Compose

## Architecture

```
Client → Load Balancer → Go Application → Redis (Cache) → PostgreSQL (Persistence)
```

## Technology Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Containerization**: Docker & Docker Compose

## Quick Start

### Prerequisites

- Docker and Docker Compose

### Using Docker Compose (Recommended)

1. Clone the repository:

```bash
git clone https://github.com/ahmadrizal1st/url-shortener.git
cd url-shortener
```

2. Copy the environment file and configure:

```bash
cp .env.example .env
# Edit .env with your configuration if needed
```

3. Start the services:

```bash
docker-compose up -d
```

4. The application will be available at `http://localhost:3000`

## Configuration

The application uses the following environment variables:

| Variable      | Description                 | Default Value    |
| ------------- | --------------------------- | ---------------- |
| `DB_ADDRESS`  | Redis database address      | `db:6379`        |
| `DB_PASSWORD` | Redis database password     | Empty            |
| `APP_PORT`    | Application server port     | `8000`           |
| `DOMAIN`      | Domain for shortened URLs   | `localhost:3000` |
| `API_QUOTA`   | Rate limit quota per window | `10`             |

### Example .env file

```env
DB_ADDRESS="db:6379"
DB_PASSWORD=""
APP_PORT="8000"
DOMAIN="localhost:3000"
API_QUOTA=10
```

## API Documentation

For comprehensive API documentation including testing procedures and examples, please refer to the [API Documentation in the testing branch](https://github.com/ahmadrizal1st/url-shortener/tree/testing/#readme).

The testing branch contains detailed information about:

- All available API endpoints
- Request and response formats
- Testing examples using Postman
- Error codes and responses
- Rate limiting details

### Base URL

- Development: `http://localhost:3000`

## Development

For the latest development features, improvements, and contribution guidelines, please visit the [development branch](https://github.com/ahmadrizal1st/url-shortener/tree/develop/#readme).

The development branch includes:

- Latest features and enhancements
- Development setup instructions
- Contribution guidelines
- Code standards and best practices
- Roadmap and upcoming features

## Docker Deployment

### Docker Compose Configuration

The `docker-compose.yml` file includes:

1. **App Service**: Go application server
2. **Redis Service**: For caching and rate limiting
3. **PostgreSQL Service**: For persistent data storage

### Running with Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Building Custom Image

```bash
docker build -t url-shortener .
docker run -p 3000:8000 --env-file .env url-shortener
```

## Project Structure

```
url-shortener/
├── api/
│   ├── database/
│   │   └── database.go
│   ├── models/
│   │   └── models.go
│   ├── routes/
│   │   ├── addTag.go
│   │   ├── deleteUrl.go
│   │   ├── editUrl.go
│   │   ├── getUrl.go
│   │   └── shorten.go
│   └── utils/
│       └── utils.go
├── db/
│   └── (database migration files, if any)
├── Dockerfile
├── .env
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Contributing

We welcome contributions! Please see our [development branch](https://github.com/ahmadrizal1st/url-shortener/tree/develop/#readme) for detailed contribution guidelines.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

If you have any questions or issues, please open an issue on GitHub.

## Roadmap

- [ ] Analytics and tracking
- [ ] User authentication and management
- [ ] API key management
- [ ] Bulk URL shortening
- [ ] QR code generation

## Acknowledgments

- Go community for excellent libraries and tools
- PostgreSQL and Redis for reliable data storage and caching
- Docker for containerization support

