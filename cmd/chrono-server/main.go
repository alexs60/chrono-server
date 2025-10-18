package main

import (
	"chrono/internal/server"
	"chrono/internal/store"
	"log"
)

func main() {
	// Create the data store.
	s := store.NewStore()

	// Create the server, giving it a reference to the store.
	srv := server.NewServer(s)

	// Start the server.
	log.Println("Chrono server started on port 8181")
	if err := srv.Start(":8181"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
