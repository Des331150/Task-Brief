package main

import (
	"fmt"
	"log"
)

func main() {
	pool, err := database.ConnectionPool()
	if err != nil {
		log.Fatalf("Fatal: %v\n", err)
	}

	defer pool.Close()

	fmt.Println("Database connection secured!")

	// mux := http.NewServeMux()

	// mux.HandleFunc("POST /job", createJob)
	// mux.HandleFunc("GET /jobs", listJobs)
	// mux.HandleFunc("GET /job/", getJob)
	// mux.HandleFunc("DELETE /job/", deleteJob)

	// err = http.ListenAndServe(":8080", mux)

	// if err != nil {
	// 	log.Fatal("Unable to start server")
	// }
	// fmt.Println("Server scaffolded on port 8080.")
}
