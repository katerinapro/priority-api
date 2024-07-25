# About this project

This is a basic priorities CRUD API which is still a work in progress. 

It builds and it runs.

What's left to do:
- unit tests
- add CI/CD

# How to run this API

## Prerequisites

Ensure you have the following installed on your system:
- PostgreSQL
- Go

### 1. Create a Database

First, create the PostgreSQL database. Connect to PostgreSQL using a superuser account (e.g., `postgres`).

```sh
psql -U postgres

CREATE DATABASE mydb;
```

### 2. Create a user
```
CREATE USER myuser WITH PASSWORD 'mypassword';
```

### 3. Create a schema

```
\c mydb
CREATE SCHEMA lo AUTHORIZATION myuser;
```

### 4. Grant user permissions

```
GRANT USAGE ON SCHEMA lo TO myuser;
GRANT CREATE ON SCHEMA lo TO myuser;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA lo TO myuser;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA lo TO myuser;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA lo TO myuser;
```

### 5. Set your DATABASE_URL environment variable
export DATABASE_URL="postgresql://username:password@localhost/dbname?sslmode=disable"

### 6. Run Migrations
```
go build -o run-migrations cmd/migrate/main.go
./run-migrations
```

### 6. Run API

```
go build -o api-server cmd/server/main.go
./api-server
```
