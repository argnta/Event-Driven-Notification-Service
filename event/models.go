package event

import "time"

type ResponseStructure struct {
	Status     string     `json:"status"`
	Error      string     `json:"error,omitempty"`
	Message    string     `json:"message"`
	EventID    string     `json:"event_id,omitempty"`
	ReceivedAt *time.Time `json:"received_at,omitempty"`
	RetryCount int        `json:"retry_count,omitempty"`
}

type UserPayload struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	ResetLink string `json:"reset_link"`
	Msg       string `json:"msg"`
}

type RequestBody struct {
	Event      string      `json:"event"`
	UserTarget string      `json:"user_target"`
	Payload    UserPayload `json:"payload"`
}
