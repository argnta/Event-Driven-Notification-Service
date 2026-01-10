package main

import (
	"Event-Driven-Notification-Service/handlers"
	"Event-Driven-Notification-Service/worker"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Loading configuration...")

	worker.StartWorkers(5)

	http.HandleFunc("/events", handlers.MainHandlers)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
