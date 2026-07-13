# Golang Banking API

A RESTful banking API built with Go, featuring user management, accounts, and financial transactions with JWT-based authentication and role-based access control.

## Features

- **Authentication & Authorization**: JWT access tokens + refresh tokens, role-based middleware (admin / user)
- **User Management**: Register, profile CRUD (admin can manage all users)
- **Account Management**: Create, update, delete accounts; each user owns their accounts
- **Transactions**: Deposit, withdrawal, and transfer with decimal precision
- **Swagger Documentation**: Auto-generated OpenAPI docs served at `/swagger/`

## Tech Stack

| Component | Technology |
|-----------|-----------|
| Language | Go 1.26.4 |
| Router | `julienschmidt/httprouter` |
| Database | MySQL (`go-sql-driver/mysql`) |
| Auth | `golang-jwt/jwt/v5` (HS256) + bcrypt |
| Decimal | `shopspring/decimal` |
| Validation | `go-playground/validator/v10` |
| Docs | `swaggo/swag` |

## Prerequisites

- Go 1.26+
- MySQL 5.7+ running on `localhost:3306`
- Database name: `golang_banking_api`

## Setup

1. **Clone and install dependencies**
   ```bash
   go mod download
   ```

2. **Create the database schema**
   ```bash
   mysql -u root -p < database.sql
   ```

3. **Configure environment variables**
   ```bash
   cp .env.example .env
   ```
   Then edit `.env`:
   ```env
   DB_URL="root:password@tcp(localhost:3306)/golang_banking_api?parseTime=true"
   JWT_SECRET="your-secret-key-here"
   ```

   If `DB_URL` is not set, it falls back to `ardhian:afnan@tcp(localhost:3306)/golang_banking_api?parseTime=true`.

4. **Run the server**
   ```bash
   go run main.go
   ```

   Server starts at `http://localhost:3000`.

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DB_URL` | No | `ardhian:afnan@tcp(localhost:3306)/golang_banking_api?parseTime=true` | MySQL connection DSN |
| `JWT_SECRET` | No | `kunci_rahasia_super_aman` | HMAC secret for JWT signing |

## API Endpoints

All requests to protected endpoints must include the header:

```
Authorization: Bearer <access_token>
```

### Public Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/register` | Register a new user |
| `POST` | `/login` | Login and receive tokens |
| `POST` | `/refresh` | Refresh access token |
| `POST` | `/logout` | Invalidate refresh token |

### Admin Endpoints (`/api/admin`)

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/admin/users` | List all users |
| `POST` | `/api/admin/users` | Create user |
| `GET` | `/api/admin/users/:userId` | Get user by ID |
| `PUT` | `/api/admin/users/:userId` | Update user |
| `DELETE` | `/api/admin/users/:userId` | Delete user |
| `GET` | `/api/admin/accounts` | List all accounts |
| `POST` | `/api/admin/accounts` | Create account |
| `GET` | `/api/admin/accounts/:accountId` | Get account by ID |
| `PUT` | `/api/admin/accounts/:accountId` | Update account |
| `DELETE` | `/api/admin/accounts/:accountId` | Delete account |
| `GET` | `/api/admin/transactions` | List all transactions |
| `POST` | `/api/admin/transactions` | Create transaction |
| `GET` | `/api/admin/transactions/:transactionId` | Get transaction by ID |
| `PUT` | `/api/admin/transactions/:transactionId` | Update transaction |
| `DELETE` | `/api/admin/transactions/:transactionId` | Delete transaction |

### User Endpoints (`/api/user`)

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/user/profile` | Get current user profile |
| `PUT` / `PATCH` | `/api/user/profile` | Update current profile |
| `GET` | `/api/user/accounts` | List own accounts |
| `POST` | `/api/user/accounts` | Create account |
| `GET` | `/api/user/accounts/:accountId` | Get own account |
| `PATCH` | `/api/user/accounts/:accountId` | Update own account |
| `DELETE` | `/api/user/accounts/:accountId` | Delete own account |
| `POST` | `/api/user/transactions/deposit` | Deposit funds |
| `POST` | `/api/user/transactions/withdraw` | Withdraw funds |
| `POST` | `/api/user/transactions/transfer` | Transfer funds |
| `GET` | `/api/user/transactions` | List own transactions |
| `GET` | `/api/user/transactions/:transactionId` | Get own transaction |

## Database Schema

### `users`

| Column | Type | Description |
|--------|------|-------------|
| `id` | INT (PK, AI) | User ID |
| `email` | VARCHAR(100) | Unique email |
| `password` | VARCHAR(100) | Bcrypt hash |
| `name` | VARCHAR(100) | Login username |
| `role` | ENUM('admin','user') | Default: `user` |
| `createdAt` | TIMESTAMP | Created timestamp |
| `updatedAt` | TIMESTAMP | Updated timestamp |

### `accounts`

| Column | Type | Description |
|--------|------|-------------|
| `id` | INT (PK, AI) | Account ID |
| `user_id` | INT (FK) | Owner user ID |
| `account_bank` | VARCHAR(100) | Bank name |
| `balance` | DECIMAL(15,2) | Account balance |
| `account_type` | ENUM('savings','checking') | Account type |
| `createdAt` | TIMESTAMP | Created timestamp |
| `updatedAt` | TIMESTAMP | Updated timestamp |

### `transactions`

| Column | Type | Description |
|--------|------|-------------|
| `id` | INT (PK, AI) | Transaction ID |
| `from_account_id` | INT (FK) | Source account |
| `to_account_id` | INT (FK) | Destination account |
| `amount` | DECIMAL(15,2) | Transaction amount |
| `type` | ENUM('transfer','deposit','withdrawal') | Transaction type |
| `description` | VARCHAR(100) | Optional description |
| `createdAt` | TIMESTAMP | Created timestamp |

### `refresh_tokens`

| Column | Type | Description |
|--------|------|-------------|
| `id` | INT (PK, AI) | Token ID |
| `user_id` | INT (FK) | Owner user ID |
| `token` | VARCHAR(255) | UUID token string |
| `expires_at` | TIMESTAMP | Expiry time |
| `created_at` | TIMESTAMP | Created timestamp |

## Authentication Flow

1. **Register** — POST `/register` with `name`, `email`, `password`, `role`
2. **Login** — POST `/login` with `name`, `password` → returns `access_token` (15m) + `refresh_token` (7d)
3. **Access protected endpoints** — Include `Authorization: Bearer <access_token>` header
4. **Refresh** — POST `/refresh` with `{"refresh_token": "..."}` when access token expires
5. **Logout** — POST `/logout` with `{"refresh_token": "..."}` to invalidate

## Swagger Docs

Swagger UI is available at `http://localhost:3000/swagger/` with the JSON spec at `http://localhost:3000/swagger.json`.

To regenerate docs after changing annotations:
```bash
swag init
```

## Testing

```bash
go test ./...
```

## Error Handling

The API returns errors in a consistent format:

```json
{
  "code": 404,
  "message": "not found",
  "status": "NOT_FOUND_ERROR"
}
```

Common HTTP status codes: `400` (bad request), `401` (unauthorized), `403` (forbidden), `404` (not found), `500` (internal server error).

## Project Structure

```
.
├── main.go                   # Entry point
├── database.sql              # MySQL schema
├── .env                      # Environment configuration
├── app/
│   ├── router.go             # Route definitions
│   ├── database.go           # DB connection
│   └── swagger.go            # Swagger handler
├── controller/               # HTTP handlers
├── service/                  # Business logic
├── repository/               # Data access layer
├── model/
│   ├── domain/               # Entities & DTOs
│   └── web/                  # Request/response models
├── middleware/                # Auth & routing middleware
├── exception/                # Error types
├── helper/                   # Utilities
├── docs/                     # Swagger generated files
└── doc.go                    # Swagger metadata
```

## Security Notes

- Change `JWT_SECRET` in production
- Change the default `DB_URL` credentials in production
- Consider enabling HTTPS in production
- The default JWT secret fallback is for development only
