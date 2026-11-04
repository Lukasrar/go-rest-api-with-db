# Go REST API with SQLite

A simple REST API for retrieving and creating events. The server uses Gin and
stores data in a local SQLite database.

## Requirements

- Go 1.27.3 or later
- A C compiler available on `PATH` to build the `go-sqlite3` dependency

## Run

From the project root, start the API:

```sh
go run .
```

The server is available at `http://localhost:8080`. On the first run, the
`api.db` file and the `events` table are created automatically. The database
is local and ignored by Git.

## Endpoints

### List events

`GET /events`

```sh
curl http://localhost:8080/events
```

Returns the saved events as JSON.

### Create an event

`POST /events`

Send a JSON body with `name`, `description`, `location`, and `dateTime`. The
date must use RFC 3339 format, for example:

```sh
curl -X POST http://localhost:8080/events \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "My event",
    "description": "Event description",
    "location": "San Francisco",
    "dateTime": "2025-01-01T15:30:00Z"
  }'
```

Request examples are also available in [`api-test/`](./api-test/) for use with
VS Code HTTP client extensions.

## Tests

Run the Go tests with:

```sh
go test ./...
```
