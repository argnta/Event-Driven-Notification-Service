package handlers

import (
	"Event-Driven-Notification-Service/event"
	"encoding/json"
	"net/http"
	"time"
)

func MainHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.URL.Path != "/events" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var reqBody event.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if reqBody.Event == "" || reqBody.UserTarget == "" {
		http.Error(w, "missing fields", http.StatusBadRequest)
		return
	}

	event.EventQueue <- reqBody

	now := time.Now()
	resp := event.ResponseStructure{
		Status:     "success",
		Message:    "event accepted",
		EventID:    "Event",
		ReceivedAt: &now,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
