# Go REST API with SQLite

A learning-oriented REST API for user accounts and event management. It is
built with Go, Gin, and SQLite, and includes password hashing, JWT
authentication, and event registrations.

## Features

- Create users and log in with email and password.
- Hash passwords with bcrypt before storing them.
- Issue JWTs for authenticated requests.
- Create, list, retrieve, update, and delete events.
- Register authenticated users for events and cancel registrations.
- Create the SQLite database and tables automatically at startup.

## Requirements

- Go 1.27.3 or later (as declared in `go.mod`).
- A C compiler available on `PATH`; the SQLite driver
  (`github.com/mattn/go-sqlite3`) uses CGO.

## Run locally

From the project root:

```sh
go run .
```

The server listens on `http://localhost:8080`. On startup, it opens or creates
`api.db` in the current working directory and creates the `users`, `events`,
and `registrations` tables if they do not already exist. Local database files
are ignored by Git.

Run the Go test suite with:

```sh
go test ./...
```

## Authentication

Create an account with `POST /signup`, then exchange the same credentials for
a token with `POST /login`:

```sh
curl -X POST http://localhost:8080/signup \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "alex@example.com",
    "password": "a-strong-password"
  }'

curl -X POST http://localhost:8080/login \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "alex@example.com",
    "password": "a-strong-password"
  }'
```

The login response contains a `token` string. Send that token as the value of
the `Authorization` header when calling a protected endpoint. The API expects
the raw token without an authorization-scheme prefix:

```sh
curl http://localhost:8080/events \
  -H 'Authorization: <token-from-login>'
```

Tokens are signed with HS256 and expire after two hours. Their claims include
the user's email, user ID, and expiration time.

## API endpoints

| Method | Path | Authentication | Description |
|---|---|---|---|
| `POST` | `/signup` | No | Create a user account |
| `POST` | `/login` | No | Validate credentials and return a JWT |
| `GET` | `/users` | No | List users |
| `GET` | `/events` | No | List all events |
| `GET` | `/events/:eventId` | No | Get one event |
| `POST` | `/events` | Yes | Create an event owned by the authenticated user |
| `PUT` | `/events/:eventId` | Yes | Update an event owned by the authenticated user |
| `DELETE` | `/events/:eventId` | Yes | Delete an event owned by the authenticated user |
| `POST` | `/events/:eventId/register` | Yes | Register the authenticated user for an event |
| `DELETE` | `/events/:eventId/register` | Yes | Cancel the authenticated user's registration |

### Event request example

The event fields are `name`, `description`, `location`, and `dateTime`. Use an
RFC 3339 date-time:

```sh
curl -X POST http://localhost:8080/events \
  -H 'Content-Type: application/json' \
  -H 'Authorization: <token-from-login>' \
  -d '{
    "name": "Community meetup",
    "description": "A local developer meetup",
    "location": "San Francisco",
    "dateTime": "2026-12-01T18:30:00Z"
  }'
```

Update an event by sending the same event fields to
`PUT /events/:eventId`. Registration and cancellation do not require a JSON
body.

## Project structure

| Path | Purpose |
|---|---|
| [`main.go`](./main.go) | Initializes the database, creates the Gin server, and registers routes |
| [`routes/routes.go`](./routes/routes.go) | Maps HTTP methods and paths to handlers; applies authentication to protected routes |
| [`routes/users.go`](./routes/users.go) | Signup, login, and user-list handlers |
| [`routes/events.go`](./routes/events.go) | Event list, detail, create, update, and delete handlers |
| [`routes/register.go`](./routes/register.go) | Event registration and cancellation handlers |
| [`middlewares/auth.go`](./middlewares/auth.go) | Reads and validates the `Authorization` header and places the user ID in the request context |
| [`utils/jwt.go`](./utils/jwt.go) | JWT generation and validation, signing algorithm, claims, expiry, and signing key |
| [`utils/hash.go`](./utils/hash.go) | bcrypt password hashing and password comparison |
| [`models/user.go`](./models/user.go) | User persistence and credential validation |
| [`models/event.go`](./models/event.go) | Event persistence and registration persistence |
| [`db/db.go`](./db/db.go) | SQLite connection setup and table creation |
| [`api-test/`](./api-test/) | HTTP request examples for exercising the API, including users, events, and registrations |
| [`go.mod`](./go.mod) | Go version and module dependencies |

## Security notes

This project is suitable for learning and local development, not production
deployment in its current form:

- The JWT signing key is hard-coded as `secretKey123` in
  [`utils/jwt.go`](./utils/jwt.go). Replace it with a securely managed secret
  before deploying.
- `GET /users` currently returns user records that include the stored bcrypt
  password hashes. Do not expose this endpoint publicly; remove the password
  field from its response and add appropriate access controls before
  production use.
- Authentication currently expects the raw JWT in `Authorization`, rather
  than a token preceded by an authorization scheme.

## HTTP request examples

Files in [`api-test/`](./api-test/) use the VS Code REST Client format and can
be run with a compatible HTTP client extension. Set a valid token in the
authorization header for protected requests.
