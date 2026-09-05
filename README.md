# Web Load Tests: Golang vs Laravel

Performance comparison of API response times between Go (Gin) and Laravel (PHP) for an event registration system.

## Tech Stack

- **Go** — Gin web framework, raw SQL (`database/sql`), MySQL driver
- **Laravel** — Eloquent ORM, MySQL
- **Database** — MySQL 8 @ `localhost:3306`, root (no password)
- **Load Testing** — k6

## Database Schema

| Table | Description |
|-------|-------------|
| `participants` | id, name, email, phone |
| `tickets` | id, name, price, bib_prefix, bib_padding, bib_increment |
| `payments` | id, participant_id, ticket_id, order_id, subtotal, status, bib_number |

BIB is auto-generated on successful payment via webhook: `{prefix}{zero-padded increment}` (e.g. `EB00001`, `RG00003`).

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/tickets` | List available tickets |
| POST | `/api/register` | Register participant + create pending payment |
| GET | `/api/participants/:id` | Get participant with payments |
| GET | `/api/payments/:orderId` | Get payment status + bib number |
| POST | `/api/webhook/payment` | Mock payment — flips status to `paid`, generates bib |

## Setup

### Prerequisites

- Go 1.21+
- PHP 8.3+ with Composer
- MySQL 8 @ `localhost:3306` (root, no password)
- k6 (for load testing)

### 1. Create databases

```bash
mysql -u root -e "CREATE DATABASE IF NOT EXISTS loadtest_golang; CREATE DATABASE IF NOT EXISTS loadtest_laravel;"
```

### 2. Golang

```bash
cd golang
go mod tidy
go run ./cmd/seed    # seed tickets
go run ./cmd/server  # start on :8080
```

### 3. Laravel

```bash
cd laravel
composer install
php artisan migrate --force
php artisan db:seed --class=TicketSeeder
php artisan serve --port=8000
```

### 4. Load test

```bash
# Against Go
k6 run --env BASE_URL=http://localhost:8080 k6/load-test.js

# Against Laravel
k6 run --env BASE_URL=http://localhost:8000 k6/load-test.js
```

Or use the batch files: `k6\run-go.bat`, `k6\run-laravel.bat`

## Project Structure

```
web-load-tests/
├── golang/
│   ├── cmd/server/main.go         # entry point
│   ├── cmd/seed/main.go           # ticket seeder
│   └── internal/
│       ├── config/database.go     # MySQL connection
│       ├── handlers/              # route handlers
│       ├── models/                # DB queries
│       └── services/bib.go        # bib generation
├── laravel/
│   ├── app/Http/Controllers/Api/  # controllers
│   ├── app/Models/                # Eloquent models
│   ├── app/Services/BibService.php
│   ├── database/migrations/
│   ├── database/seeders/
│   └── routes/api.php
└── k6/
    ├── load-test.js               # main k6 script
    ├── load-test-detailed.js      # with per-endpoint metrics
    ├── run-go.bat
    └── run-laravel.bat
```

## Ticket Seed Data

| Name | Price | Bib Prefix | Padding |
|------|-------|-----------|---------|
| Early Bird | $25.00 | EB | 5 |
| Regular | $50.00 | RG | 5 |
| VIP | $100.00 | VIP | 4 |

## License

MIT
