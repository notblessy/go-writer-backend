# go-writter-backend

goWriter is a Backend article API based on Golang, Gin and Sqlx.

## Tech Requirements

- Go
- Mysql
- Elasticsearch
- RabbitMQ

## How to Run

### Clone

First clone this repo by run:

```sh
$ git clone git@github.com:notblessy/go-writer-backend.git
```

### Environtment

- Don't forget to set `.env` for web and workers (in directory `/workers`), you can copy from `env.sample`

### Database Migration

- To migrate tables, ensure you create `Makefile` from `Makefile.sample` then run

```sh
$ make migrate-up
```

### Running project

- run for debugging by `go run main.go`
- ensure that you run worker in folder `workers/` then `go run *.go`

### API Test

- API can be tested by running `go test`
- All API docs are available in docs folder. The API docs are documented using Postman in the directory `docs/`

## Author

```
Frederich Blessy
```
