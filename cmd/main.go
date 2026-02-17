package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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
}
