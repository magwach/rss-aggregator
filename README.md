# RSS Aggregator

The scraper periodically fetches feed sources, parses new items, stores them in DB, and provides access via the API. You can trigger the scraper manually or rely on scheduled polling.

## Table of Contents
- [Overview](#overview)  
- [Features](#features)  
- [Tech Stack](#tech-stack)  
- [Getting Started](#getting-started)  
  - [Prerequisites](#prerequisites)  
  - [Installation](#installation)  
  - [Configuration](#configuration)  
  - [Database Setup](#database-setup)  
  - [Running Locally](#running-locally)
- [Project Structure](#project-structure)
- [Testing](#testing)  
- [Contributing](#contributing)  
- [License](#license)  

## Overview  
This project provides a backend service for aggregating RSS (and/or Atom) feeds, parsing them, and storing feed items in a database, so you can build downstream functionality — for example a UI feed reader, notifications of new items, or automated processing of feed items.

## Features  
- Fetches and parses RSS / Atom feeds  
- Supports multiple sources and keeps track of previously-seen items  
- REST API to manage users, feed sources, subscriptions, and retrieved items  
- Lightweight Go implementation for performance and ease of deployment  
- Ready to integrate with frontend UIs or further processing pipelines  

## Tech Stack  
- Language: **Go** (Golang)  
- Database: e.g., **PostgreSQL** (via `sqlc` or similar query tooling)  
- HTTP API: built using Go standard libs or/minimal framework  
- Feed parsing / scraping logic in `rss.go` / `scraper.go`  

## Getting Started  

### Prerequisites  
- Go version (e.g., 1.20 or later)  
- A running PostgreSQL (or compatible) database  
- Git  
- (Optional) Docker if you prefer containerised setup  

### Installation  

    git clone https://github.com/magwach/rss-aggregator.git
    cd rss-aggregator
    go mod tidy

### Configuration
Create an environment file (e.g., .env) with values such as:

      PORT=8080
      DB_URL=postgresql://user:password@localhost:5432/rssagg?sslmode=disable

### Database Setup
- Create the database/schema in PostgreSQL (see sql/ folder or sqlc.yaml).
- Run migrations if included (e.g., sqlc, manual SQL scripts).
- Ensure that tables for users, feed sources, feed items, subscriptions (or follows) are present.

### Running Locally

    go build -o rss-aggregator
    ./rss-aggregator
    
### Project Structure
    .
    ├── handler/             # HTTP handlers: authentication, users, feeds, readiness
    ├── middleware/          # e.g., auth middleware
    ├── models.go            # DB model definitions
    ├── rss.go               # Feed parsing logic
    ├── scraper.go           # Polling / scraping logic
    ├── json.go              # JSON request/response mapping
    ├── main.go              # Application entry point
    ├── sql/                  # SQL schema and queries (sqlc)  
    ├── sqlc.yaml            # Configuration for sqlc (if used)
    ├── internal/             # Any internal packages 
    └── vendor/              # Vendor dependencies (if vendored)

### Testing

- Unit test your feed parsing logic (rss.go) and handler logic.
- Integrate with your database in test mode (or use a SQLite/ test DB) for full stack API tests.
-Use a tool like curl, Postman or HTTPie to test endpoints manually.

### Contributing

-Fork the repository.
-Create a new feature branch (git checkout -b feature/my-feature).
-Write code + tests.
-Submit a pull request with a clear description of your changes.
-Ensure existing tests pass and new functionality is well documented.

### License

This project is licensed under the MIT License 


