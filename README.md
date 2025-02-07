# Casino Wallet Service

Casino Wallet Service is a RESTful API service written in Go, built to manage casino wallet transactions. This service provides features like wallet authentication, balance checking, credit, and debit operations.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (1.16 or later)
- Docker and Docker Compose (For Redis and MySQL)

### Installing

```bash
# 1. Clone the repository
git clone https://github.com/yigitalpkilavuz/casino_wallet.git

# 2. Go to the project directory
cd casino_wallet

# 3. Run Docker Compose to set up the Redis and MySQL services
docker-compose up -d

# 4. Install the necessary Go dependencies
go mod tidy

# 5. Run the service
go run main.go
```


## Running the Tests

To run the tests for the service, navigate to the service directory and run the following command:
go test

## Endpoints

- POST /wallet/authenticate
- GET /wallet/{id}/balance
- POST /wallet/{id}/credit
- POST /wallet/{id}/debit

## Built With

- [Go](https://golang.org/) - The programming language used
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [GORM](https://gorm.io/) - ORM library for Go
- [Redis](https://redis.io/) - In-memory data structure store used as a cache
- [MySQL](https://www.mysql.com/) - Database used

## TODO

- Fix tests: Currently, there are issues with the tests. They need to be fixed to ensure proper test coverage.