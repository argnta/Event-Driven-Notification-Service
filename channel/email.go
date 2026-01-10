package channel

import (
	"Event-Driven-Notification-Service/event"
	"fmt"
)

func SendEmail(evt event.RequestBody) bool {
	fmt.Println("Sending email to:", evt.Payload.Email)
	return true
}
