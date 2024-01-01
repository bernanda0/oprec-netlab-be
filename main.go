package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bernanda0/oprec-netlab-be/database"
	"github.com/bernanda0/oprec-netlab-be/database/sqlc"
	"github.com/bernanda0/oprec-netlab-be/handlers"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	l := log.New(os.Stdout, "NETLAB-SERVER-", log.LstdFlags)
	ctx := context.Background()

	env := os.Getenv("APP_ENV")
	// If it's in development environment then load the .env
	if env != "prod" {
		// load the .env
		err := godotenv.Load(".env")
		if err != nil {
			l.Fatalf("Error reding the .env %s", err)
		}
	}

	// CRUD
	db, queries := database.Instantiate(l)
	if db == nil || queries == nil {
		l.Println("Exiting due to database connection error")
		return
	}
	defer db.Close()

	server := &http.Server{
		Addr:        ":" + os.Getenv("PORT_NUMBER"),
		Handler:     defineMultiplexer(l, queries),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: time.Second,
	}

	// now the startServer is run by a routine
	go startServer(server, l)

	// inorder to block the routine, we might use a channel (we can use wait group also)
	shut := make(chan os.Signal, 1)
	signal.Notify(shut, syscall.SIGINT, syscall.SIGTERM)

	<-shut // Block until a signal is received

	timeout_ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stopServer(server, l, &timeout_ctx, &cancel)
}

func startServer(s *http.Server, l *log.Logger) {
	l.Println("🔥 Server is running on", s.Addr)

	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		l.Fatalln("Server is failed due to", err)
	}
}

func stopServer(s *http.Server, l *log.Logger, ctx *context.Context, cancel *context.CancelFunc) {
	l.Println("💅 Shutting down the server")
	s.Shutdown(*ctx)
	c := *cancel
	c()
}

func defineMultiplexer(l *log.Logger, q *sqlc.Queries) http.Handler {
	// handle multiplexer
	mux := http.NewServeMux()

	hello_handler := handlers.NewHello(l)
	mux.Handle("/hello", hello_handler)

	corsMiddleware := cors.AllowAll().Handler

	handler := corsMiddleware(mux)
	return handler
}
