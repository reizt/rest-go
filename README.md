# REST Go

A simple REST API boilerplate written in Go, including authentication, database, and email.

## Stack

- 🔥 Framework: Echo
- 🔄 ORM: Ent
- 🥫 Database: PostgreSQL
- 📬 Mailing: SendGrid
- 🔑 Signing: JWT
- 🥔 Hashing: Bcrypt

## Directory structure

- `e2e` - E2E test
- `ent` - ent schema files
- `entities` - Domain models
- `handlers` - Echo handlers
- `iservices` - Interface of services
- `iusecases` - Interface of usecases
- `mservices` - Structs for mocking services
- `router` - Router
- `services` - Implementation of services
- `usecases` - Implementation of usecases

## Try it

Set environment variables in `.env`
```
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
JWT_PRIVATE_KEY="xxx"
JWT_PUBLIC_KEY="xxx"
SENDGRID_API_KEY="xxx"
MAILER_FROM="xxx"

TEST_CLEAR_DATABASE=on
TEST_GENERATE_CODE_FIXED_VALUE=123456
```

Generate ent code
```sh
go generate ./ent
```

Launch database & API server
```sh
docker compose up
```

Run tests
```sh
go test -v ./...
```
