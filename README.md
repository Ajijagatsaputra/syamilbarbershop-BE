# Syamil Barbershop Backend API

Backend service for Syamil Barbershop application, built with Go and Gin framework.

## Tech Stack
- **Language**: Go (Golang)
- **Framework**: Gin Gonic
- **ORM**: GORM
- **Database**: MySQL
- **Containerization**: Docker & Docker Compose

## Prerequisites
- [Go](https://golang.org/doc/install) (version 1.21 or higher)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### 1. Repository Setup
Clone the repository to your local machine:
```bash
git clone <repository-url>
cd syamilbarbershop-be
```

### 2. Environment Variables
Copy the `.env.example` file to `.env` and adjust the values as needed:
```bash
cp .env.example .env
```

### 3. Running with Docker (Recommended)
The easiest way to get the project running is using Docker Compose. This will start the API, MySQL database, and phpMyAdmin.
```bash
docker-compose up -d
```
- **API**: [http://localhost:8081](http://localhost:8081)
- **phpMyAdmin**: [http://localhost:8082](http://localhost:8082)

### 4. Running Locally
If you want to run the Go application directly:
1. Ensure you have a MySQL database running (you can use the one from docker-compose).
2. Update `.env` to point to `localhost` and the correct port (3309).
3. Run the application:
```bash
go run cmd/api/main.go
```

## Project Structure
- `cmd/api`: Entry point of the application.
- `internal/config`: Configuration and environment variable loading.
- `internal/handlers`: HTTP request handlers.
- `internal/models`: Database models and GORM structures.
- `internal/routes`: API route definitions.
- `internal/middleware`: Custom Gin middlewares (e.g., JWT Auth).

## Security Note
The `.env` file should **never** be committed to the repository. If you accidentally committed it, run the following command to remove it from the Git index:
```bash
git rm --cached .env
git commit -m "chore: remove .env from tracking"
git push
```
