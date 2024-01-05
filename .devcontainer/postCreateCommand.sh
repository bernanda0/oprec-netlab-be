#!/bin/bash

# Update package list and install PostgreSQL client
apt-get update && apt-get install -y postgresql-client

# Install Golang Gin
go install github.com/codegangsta/gin@latest

# Install Golang sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Install Golang migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

