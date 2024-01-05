# Load environment variables from .env
include .env
export

.PHONY: run-dev run-nodemon psql

# run the server using gin
run-dev:
	gin --immediate 

# run the server using nodemon
run-nodemon:
	nodemon --exec go run main.go --signal SIGTERM

# run psql with remote host
psql:
	psql postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}/${DB_NAME}?options=endpoint%3D${DB_ENDPOINT}

# Example of creating new migration scheme for session table
migrate-create-session:
	migrate create -ext sql -dir database/migration -seq session