package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type CurrentStage string

const (
	Queued = "queued"
	Done   = "done"
)

type Jobs struct {
	ID            int64
	AudioID       int64
	CurrentStatus CurrentStage
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func ConnectionPool() (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connection, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = connection.Ping(context.Background())
	if err != nil {
		return nil, err

	}

	return connection, nil
}

func main() {
	pool, err := ConnectionPool()
	if err != nil {
		log.Fatalf("Fatal: %v\n", err)
	}

	defer pool.Close()

	fmt.Println("Database connection secured!")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /job", createJob)
	mux.HandleFunc("GET /jobs", listJobs)
	mux.HandleFunc("GET /job/", getJob)
	mux.HandleFunc("DELETE /job/", deleteJob)

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Unable to start server")
	}
	fmt.Println("Server scaffolded on port 8080.")
}

func createJob(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating job...."))
}

func listJobs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listing jobs...."))
}
func getJob(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting job...."))
}
func deleteJob(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting job...."))
}
