# TestCase_3 - Docker Setup with Dbmate Migrations

## Quick Start

```bash
cd TestCase_3
docker compose up --build
```

This will:
1. Start PostgreSQL database (internal only)
2. Run Dbmate migrations in dedicated service
3. Seed sample data from CSV files
4. Build and start the Go application (port 8090)

## Access Points

- **Application**: http://localhost:8090
- **Database**: Not exposed externally (internal only)

## Architecture

```
┌─────────────────────────────────────────────────────┐
│              docker-compose.yml                      │
│                                                      │
│  postgres:15 ──► migrate:latest ──► app:latest     │
│      :5432         dbmate up           :8090        │
└─────────────────────────────────────────────────────┘

Execution Flow:
1. postgres starts and becomes healthy
2. migrate runs dbmate migrations, exits (code 0)
3. app starts after migrations succeed
```

## Services

### PostgreSQL
- **Image**: postgres:15-alpine
- **Database**: sandbox
- **Credentials**: postgres/(empty password)
- **No port exposure**: Only accessible internally
- **No persistence**: Data is ephemeral (no volumes)

### Migrate (Init Container)
- **Purpose**: Runs database migrations once
- **Tool**: Dbmate
- **Runs on**: Every deployment
- **Restart policy**: "no" (exits after completion)

### Application
- **Port**: 8090 (mapped to internal 8080)
- **Runs**: After migrations complete successfully
- **Auto-seeding**: Seeds sample data from CSV

## Commands

```bash
# Full deployment (builds all images)
docker compose up --build

# Start in detached mode
docker compose up --build -d

# View all logs
docker compose logs -f

# View only migration logs
docker compose logs -f migrate

# View only app logs
docker compose logs -f app

# Run migrations only (for debugging)
docker compose run migrate

# Stop all services
docker compose down

# Stop and remove volumes (fresh start)
docker compose down -v
```

## Project Structure

```
TestCase_3/
├── migrations/           # Dbmate migration files
│   ├── 00001_create_schema.sql
│   └── 00002_seed_data.sql
├── seed-data/           # CSV files for seeding
│   ├── users.csv
│   └── user_transactions.csv
├── scripts/
│   └── docker-entrypoint.sh  # App startup script
├── Dockerfile           # Multi-stage Docker build
├── docker-compose.yml   # Service orchestration (3 services)
├── .env.docker          # Docker environment configuration
└── .dockerignore       # Docker build exclusions
```

## Configuration

Edit `.env.docker` to customize:
- Database connection settings
- Application settings
- Server configuration

## Customization

### Replace Sample Data

Replace the CSV files in `seed-data/` with your own:
- `seed-data/users.csv`
- `seed-data/user_transactions.csv`

### Add New Migrations

Create new migration files in `migrations/`:
- `00003_add_new_table.sql`
- `00004_update_schema.sql`

Dbmate will automatically run them on next deployment.

Migration file format:
```sql
-- +dbmate Up

CREATE TABLE ...;

-- +dbmate Down

DROP TABLE ...;
```

## Troubleshooting

### Migration Fails

Check migration logs:
```bash
docker compose logs migrate
```

Run migrations manually:
```bash
docker compose run migrate
```

### Database Connection Issues

Ensure PostgreSQL is healthy:
```bash
docker compose ps
docker compose exec postgres pg_isready -U postgres
```

### Clean Restart

```bash
docker compose down -v
docker compose up --build
```

### View All Service Logs

```bash
docker compose logs -f
```

### Check Migration Status

```bash
docker compose logs migrate
```

### Rebuild After Changes

If you modify migrations or configuration:
```bash
docker compose down
docker compose up --build
```

## Requirements

- Docker Engine 20.10+
- Docker Compose V2
- 2GB RAM available
