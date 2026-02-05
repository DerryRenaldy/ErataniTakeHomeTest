# TestCase_3 - Local Development Guide

A Go-based REST API service for managing users and credit card transactions.

## Prerequisites

Ensure you have the following installed on your machine:

- **Go** 1.25 or higher
- **PostgreSQL** 15 or higher (running locally or via Docker)
- **Air** (for hot-reloading during development)
- **Make**

## Installing Air

Air provides hot-reload functionality for development:

```bash
go install github.com/air-verse/air@latest
```

Make sure your PATH includes the Go bin directory:

```bash
export PATH=$PATH:$HOME/go/bin
```

## Database Setup

Before running the application, you need to set up the PostgreSQL database.

### 1. Start PostgreSQL

Start your local PostgreSQL service:

**Using Homebrew (macOS):**
```bash
brew services start postgresql
```

**Using Docker:**
```bash
docker run --name postgres-local \
  -e POSTGRES_DB=sandbox \
  -e POSTGRES_USER=postgres \
  -p 5432:5432 \
  -d postgres:15-alpine
```

**Using system service (Linux):**
```bash
sudo systemctl start postgresql
```

### 2. Create the Database

Create a new database named `sandbox`:

**Using your PostgreSQL client** (pgAdmin, DBeaver, TablePlus, DataGrip, psql, or any other tool):

```sql
CREATE DATABASE sandbox;
```

### 3. Create the Tables

Execute the following SQL statements in your PostgreSQL client:

```sql
-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    country VARCHAR(100),
    credit_card_type VARCHAR(50),
    credit_card_number VARCHAR(20),
    first_name VARCHAR(100),
    last_name VARCHAR(150)
);

-- Create user_transactions table
CREATE TABLE IF NOT EXISTS user_transactions (
    id BIGSERIAL PRIMARY KEY,
    id_user INT NOT NULL,
    total_buy BIGINT NOT NULL
);

-- Add foreign key relationship
ALTER TABLE user_transactions
ADD CONSTRAINT fk_user_transactions_user
FOREIGN KEY (id_user)
REFERENCES users(id);
```

You can also copy the contents from the file `migrations/00001_create_schema.sql` and run it in your PostgreSQL client.

### 4. Verify the Setup

In your PostgreSQL client, run:

```sql
\dt
```

You should see:

```
            List of relations
 Schema |         Name         | Type  |  Owner
--------+----------------------+-------+----------
 public | users                | table | postgres
 public | user_transactions    | table | postgres
(2 rows)
```

If both tables appear, your database is ready!

## Quick Start

Start the development server with hot-reload:

```bash
make dev
```

The application will be available at: **http://localhost:8080**

## Available Commands

| Command | Description |
|---------|-------------|
| `make dev` | Start development server with hot-reload |
| `make build` | Build the application binary |
| `make run` | Build and run the application |
| `make test` | Run all tests |
| `make wire` | Regenerate wire dependency injection |
| `make deps` | Download and tidy Go dependencies |

## API Endpoints

### Health Check

```bash
curl http://localhost:8080/health
```

### Users API

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/v1/users` | Get all users |
| GET | `/v1/users/{id}` | Get user by ID |
| POST | `/v1/users` | Create a new user |

**Get all users:**
```bash
curl http://localhost:8080/v1/users
```

**Get user by ID:**
```bash
curl http://localhost:8080/v1/users/1
```

**Create a new user:**
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "country": "Indonesia",
    "credit_card_type": "Visa",
    "credit_card_number": "4111111111111111",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

### Transactions API

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/v1/transactions` | Get all transactions with user info |
| GET | `/v1/transactions/credit-card-types` | Get credit card type statistics |

**Get all transactions:**
```bash
curl http://localhost:8080/v1/transactions
```

**Get credit card type statistics:**
```bash
curl http://localhost:8080/v1/transactions/credit-card-types
```

## Configuration

The application uses environment variables defined in `.env`. Here are the default settings:

```bash
# Application Settings
APP.NAME=TestCase_3
APP.URL=http://localhost:8080

# Database Connection (Read)
DB.PG.READ.HOST=localhost
DB.PG.READ.PORT=5432
DB.PG.READ.USER=postgres
DB.PG.READ.PASSWORD=
DB.PG.READ.NAME=sandbox

# Database Connection (Write)
DB.PG.WRITE.HOST=localhost
DB.PG.WRITE.PORT=5432
DB.PG.WRITE.USER=postgres
DB.PG.WRITE.PASSWORD=
DB.PG.WRITE.NAME=sandbox

# Server Settings
SERVER.ENV=development
SERVER.LOG_LEVEL=debug
SERVER.PORT=8080
```

To modify settings, edit the `.env` file in the project root.

## Project Structure

```
TestCase_3/
├── main.go                     # Application entry point
├── Makefile                    # Build and development commands
├── .env                        # Environment configuration
├── go.mod                      # Go module definition
├── go.sum                      # Go module checksums
├── configs/
│   └── config.go              # Configuration loading
├── database/
│   └── postgres.go            # PostgreSQL connection
├── internal/
|   └── handlers/
│       └── credit_card/
│           └── handler/        # HTTP handlers
│   └── domain/
│       └── credit_card/
│           ├── model/         # Data models
│           ├── repository/    # Data access layer
│           └── service/       # Business logic
├── transport/
│   └── http/
│       ├── http.go            # HTTP server setup
│       ├── router/            # Route definitions
│       └── response/          # Response utilities
├── migrations/
│   ├── 00001_create_schema.sql
│   └── 00002_seed_data.sql
└── seed-data/
    ├── users.csv
    └── user_transactions.csv
```

## Troubleshooting

### Cannot Connect to Database

1. **Verify PostgreSQL is running**
   - Check that your PostgreSQL service or Docker container is running

2. **Confirm the database exists**
   - Ensure you created the `sandbox` database

3. **Check your connection settings**
   - Host: `localhost`
   - Port: `5432`
   - Database: `sandbox`
   - Username: `postgres`
   - Password: (leave empty)

### Tables Already Exist

If you see errors about tables already existing:

1. Drop the existing tables in your PostgreSQL client:

```sql
DROP TABLE IF EXISTS user_transactions;
DROP TABLE IF EXISTS users;
```

2. Re-run the CREATE TABLE statements from Step 3 in the Database Setup section.

### Port 8080 Already in Use

If you get an error about port 8080 being in use:

1. Stop the process using that port, or

2. Change the port in `.env`:

```bash
SERVER.PORT=8081
```

### Air Command Not Found

If `make dev` fails with "air: command not found":

```bash
# Install Air
go install github.com/air-verse/air@latest

# Add Go bin to PATH
export PATH=$PATH:$HOME/go/bin

# Verify installation
which air
```

## Docker Development

For containerized development with automatic database setup and migrations, see [README-DOCKER.md](README-DOCKER.md).


## SQL Answers

Answer 3a and 3b are located in the `sql/` folder.

## License

This project is part of an assessment test.
