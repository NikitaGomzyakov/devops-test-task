# DevOps/Backend SWE Take Home Test

This project implements a simple Backend API in Go that processes data asynchronously using Redis as a storage backend. It is fully containerized and includes Kubernetes manifests for deployment.

## Tech Stack
- **Language:** Go 1.23
- **Database:** Redis
- **Orchestration:** Docker Compose / Kubernetes
- **Framework:** Gorilla Mux

## Features
- `POST /process-data`: Receives JSON payload, initiates a 7-second mock processing task, and stores status in Redis.
- `GET /results/{id}`: Retrieves the current status and result of the processing task.
- **Asynchronous Processing:** Uses Go routines to handle data processing without blocking the API.
- **Containerization:** Multistage Dockerfile for optimized image size.

## Project Structure
```text
.
├── cmd/api/main.go          # Application entry point
├── internal/handler/        # HTTP handlers logic
├── internal/storage/        # Redis client configuration
├── k8s/                     # Kubernetes Deployment & Service manifests
├── Dockerfile               # Multistage build file
└── docker-compose.yml       # Local development setup
