# Load environment variables from .env
include .env
export

POSTGRE_CONN_STRING = "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}/${DB_NAME}?options=endpoint%3D${DB_ENDPOINT}"

.PHONY: run-dev run-nodemon psql

# run the server using gin
run-dev:
	gin --immediate 

# run the server using nodemon
run-nodemon:
	nodemon --exec go run main.go --signal SIGTERM

# run psql with remote host
psql:
	psql $(POSTGRE_CONN_STRING)

# Example of creating new migration scheme for session table
migrate-create:
	migrate create -ext sql -dir database/migration -seq $(name)
# Command to migrate
migrate-up:
	migrate -path database/migration -database $(POSTGRE_CONN_STRING) -verbose up $(n)
migrate-down:
	migrate -path database/migration -database $(POSTGRE_CONN_STRING) -verbose down $(n)
migrate-fix:
	migrate -path database/migration -database $(POSTGRE_CONN_STRING) force $(v) 
	@echo "y" | make migrate-down
