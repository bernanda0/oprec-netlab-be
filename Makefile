run-dev:
	gin --immediate 


# Example of creating new migration scheme for session table
migrate-create-session:
	migrate create -ext sql -dir database/migration -seq session