# Azure Service Bus Queue Emulator!

The Azure Service Bus Queue Emulator allows developers to emulate Azure Service Bus locally for development and testing purposes. It provides an HTTP API mirroring the Azure Service Bus, and uses AMQP for communication, backed by RabbitMQ as the message broker.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Emulator](#running-the-emulator)
- [API Documentation](#api-documentation)
- [Development](#development)
  - [Project Structure](#project-structure)
  - [Contributing](#contributing)
- [Testing](#testing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Features

- Emulates Azure Service Bus's HTTP API.
- Backed by RabbitMQ for reliable message delivery and queue management.
- Extendable broker layer allows for easy integration of other message brokers in the future.
- Swagger UI for API documentation and testing.

## Getting Started

### Prerequisites

- Go (version >= 1.16)
- Docker & Docker Compose
- RabbitMQ (if not using Docker)

### Installation

1. Clone the repository:

   ```bash
   git clone https://garage-09.visualstudio.com/ASB-Queue-Emulator-Hackathon/_git/ASB-Queue-Emulator-Hackathon
   cd ASB-Queue-Emulator-Hackathon
   ```

2. Build the project:

   ```bash
   go build ./cmd
   ```

3. Run the project:
    ```
    go run ./cmd/main.go [--config "path/to/config.json"]
    ```

    > Note: config/default_config.json is used if no value for config is supplied

### Running the Emulator

Using Docker Compose:

```bash
docker compose up
```

This will start the API server, set up the AMQP connection, and start the RabbitMQ instance.

## API Documentation

Access the Swagger UI at: `http://localhost:your-port/swagger-ui/`  
This provides comprehensive documentation on all available API endpoints, including request/response formats and example payloads.

## Development

### Project Structure

- `/cmd`: Application's entry point.
- `/api`: API endpoint handlers and middleware.
- `/pkg`: Core logic, including broker implementations and AMQP-specific connections.
- `/swagger`: OpenAPI specification and generated code.
- `/tests`: Unit and integration tests.

```
/ASB-Queue-Emulator-Hackathon
|-- /cmd
|   |-- main.go  # Entry point of the application, sets up the server, logging, and DI
|
|-- /api
|   |-- handlers.go  # Contains API endpoint handlers based on the generated stubs
|   |-- middlewares.go  # Contains any middleware (e.g., for logging, CORS)
|
|-- /pkg
|   |-- /broker
|   |   |-- /abstract
|   |   |   |-- broker.go  # Interface definitions for the broker
|   |   |   |-- models.go  # Shared models (like Message) used across brokers
|   |   |
|   |   |-- /rabbitmq
|   |       |-- client.go  # Initializes and provides the RabbitMQ client using AMQP
|   |       |-- operations.go  # Implements the broker interfaces for RabbitMQ
|   |
|   |-- /amqp
|       |-- gateway.go  # Handles AMQP to HTTP translation and vice versa
|       |-- connection.go  # AMQP-specific connection setup, SAS token validation
|       |-- utils.go  # Utility functions specific to AMQP operations (like parsing connection strings)
|
|-- /models
|   |-- api_requests.go  # Models for API request payloads
|   |-- api_responses.go  # Models for API response payloads
|
|-- /swagger
|   |-- azure-servicebus-spec.yaml  # Swagger/OpenAPI specification for the API
|   |-- /gen  # Directory for the generated code using Go Swagger
|
|-- /tests
|   |-- /unit
|   |   |-- broker_tests.go
|   |   |-- api_tests.go
|   |   |-- amqp_tests.go  # Unit tests related to AMQP operations
|   |
|   |-- /integration
|       |-- end_to_end_tests.go  # High-level tests simulating real-world scenarios
|
|-- /scripts
|   |-- init_db.sh  # (If needed) Initialization scripts for RabbitMQ setup
|   |-- load_test.sh  # Scripts to run load testing tools like vegeta
|
|-- Dockerfile  # Docker configuration for the Go application
|-- docker-compose.yml  # Sets up linked containers, e.g., Go app + RabbitMQ
|-- README.md  # Comprehensive documentation of the project
|-- .gitignore  # To exclude binaries, vendor folders, etc.
|-- go.mod  # Go modules file
|-- go.sum  # Go modules checksums

```

### Contributing

For contributing guidelines, please refer to the [CONTRIBUTING.md](./CONTRIBUTING.md) file. *Not currently needed*

## Testing

After making any changes, ensure all unit tests pass:

```bash
go test ./...
```

## License
TBD

## Acknowledgments

- Azure Service Bus team for their comprehensive documentation.
- The Go, Docker, RabbitMQ, and AMQP communities for their invaluable resources.