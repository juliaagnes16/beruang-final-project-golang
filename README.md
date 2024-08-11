# BerUang

## Overview
BerUang adalah aplikasi manajemen keuangan pribadi yang memungkinkan pengguna untuk mencatat dan mengelola pengeluaran mereka. Aplikasi ini menggunakan Gin Gonic sebagai web framework dan PostgreSQL sebagai database. Autentikasi dilakukan menggunakan JWT.

## Installation
1. Clone repository ini.
2. Install dependensi dengan `go mod tidy`.
3. Setup database PostgreSQL.
4. Jalankan aplikasi dengan `go run main.go`.

## API Endpoints
### Auth
- **POST /register**: Register a new user.
- **POST /login**: Login to get a JWT token.

### Expenses
- **GET /expenses**: Get all expenses (requires authentication).
- **POST /expenses**: Create a new expense (requires authentication).

### Categories
- **GET /categories**: Get all categories.
- **POST /categories**: Create a new category.
- **PUT /categories/:id**: Update a category.
- **DELETE /categories/:id**: Delete a category.

### Budgets
- **GET /budgets**: Get all budgets (requires authentication).
- **POST /budgets**: Create a new budget (requires authentication).

## Deployment
Deploy the application using your preferred method, such as Docker or a cloud service like Heroku or Railway.

## Database Schema
- `User` table with columns `ID`, `Username`, `Password`, `CreatedAt`, `UpdatedAt`.
- `Expense` table with columns `ID`, `UserID`, `CategoryID`, `Amount`, `Note`, `CreatedAt`, `UpdatedAt`.
- `Category` table with columns `ID`, `Name`, `CreatedAt`, `UpdatedAt`.
- `Budget` table with columns `ID`, `UserID`, `CategoryID`, `Amount`, `CreatedAt`, `UpdatedAt`.