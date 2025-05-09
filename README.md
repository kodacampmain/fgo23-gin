# 🚀 FGO 23 GIN

![alt text](https://img.shields.io/badge/Go-black?style=for-the-badge&logo=go)
![alt text](https://img.shields.io/badge/PostgreSQL-black?style=for-the-badge&logo=postgresql)
![alt text](https://img.shields.io/badge/Redis-black?style=for-the-badge&logo=redis)

Backend project written in Go using gin-gonic framework as the backend engine and struct validation, redis for caching, and PostgreSQL as the database.

## 📦 Tech Stack

- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Redis](https://redis.io/docs/latest/operate/oss_and_stack/install/archive/install-redis/install-redis-on-windows/)
- [JWT](https://github.com/golang-jwt/jwt)
- [argon2](https://pkg.go.dev/golang.org/x/crypto/argon2)
- [migrate](https://github.com/golang-migrate/migrate)
- [Docker](https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository)
- [Swagger for API docs](https://swagger.io/) + [Swaggo](https://github.com/swaggo/swag)

## 🌎 Environment

```sh
// see env.example
DBNAME=<YOUR_DB_NAME>
DBUSER=<YOUR_DB_USER>
DBHOST=<YOUR_DB_HOST>
DBPORT=<YOUR_DB_PORT>
DBPASS=<YOUR_DB_PASS>

JWT_SECRET=<YOUR_JWT_SECRET>
JWT_ISSUER=<YOUR_JWT_ISSUER>

RDSHOST=<YOUR_REDIS_HOST>
RDSPORT=<YOUR_REDIS_PORT>
```

## 🔧 Installation

1. Clone the project

```sh
$ git clone https://github.com/kodacampmain/fgo23-gin.git
```

2. Navigate to project directory

```sh
$ cd fgo23-gin
```

3. Install dependencies

```sh
$ go mod tidy
```

4. Setup your [environment](#-environment)
5. Install [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation) for DB migration
6. Do the DB Migration

```sh
$ migrate -database YOUR_DATABASE_URL -path ./db/migrations up
```

7. Run the project

```sh
$ go run ./cmd/main.go
```

## 🚧 API Documentation

| Method | Endpoint  | Body                          | Description                                          |
| ------ | --------- | ----------------------------- | ---------------------------------------------------- |
| GET    | /img      |                               | Static File                                          |
| GET    | /ping     |                               | Connection Testing (should be responded with a pong) |
| GET    | /users    |                               | Get All User                                         |
| POST   | /auth     | email:string, password:string | Login                                                |
| POST   | /auth/new | email:string, password:string | Register                                             |

[Full Documentation]()

## 📄 LICENSE

MIT License

Copyright (c) 2025 Koda Tech Academy

## 📧 Contact Info

## 🎯 Related Project

[React-FGO23](https://github.com/kodacampmain/react-fgo23)
