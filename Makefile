run-dev:
	gin --immediate 
run-nodemon:
	nodemon --exec go run main.go --signal SIGTERM
# Example of creating new migration scheme for session table
migrate-create-session:
	migrate create -ext sql -dir database/migration -seq session