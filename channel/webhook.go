package channel

import (
	"Event-Driven-Notification-Service/event"
	"bytes"
	"net/http"
)

func SendWebhook(evt event.RequestBody) bool {
	resp, err := http.Post("http://example.com/webhook", "application/json", bytes.NewBuffer([]byte(evt.Event)))
	if err != nil {
		return false
	}
	return resp.StatusCode == http.StatusOK
}
