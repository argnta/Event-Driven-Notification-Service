package channel

import (
	"Event-Driven-Notification-Service/event"
	"fmt"
)

func SendInApp(evt event.RequestBody) bool {
	fmt.Println("In-App notification for:", evt.UserTarget)
	return true
}
