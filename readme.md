# Gin Money Manager API

A backend application built with Go as a learning project to explore modern backend development patterns and software design principles.

## Overview

This project serves as a playground for experimenting with concepts commonly found in real-world applications, including:

* Domain-Driven Design (DDD)
* CQRS (Command Query Responsibility Segregation)
* Repository Pattern
* JWT Authentication
* Authorization Policies
* Middleware-based Access Control
* Go Generics
* Modular Architecture

The primary goal of this repository is learning, experimentation, and continuous improvement while building a maintainable Go application.

## Tech Stack

* Go
* Gin
* GORM
* PostgreSQL
* JWT

## Running the Project

Install dependencies:

```bash
go mod tidy
```

Create environment file:

```bash
cp .env.example .env
```

Run the application:

```bash
go run main.go
```

## Running Tests

```bash
go test ./...
```

## Notes

This project is actively developed as part of my Go learning journey. The architecture and implementation may evolve over time as I continue exploring new patterns, practices, and ideas within the Go ecosystem.

## License

MIT
