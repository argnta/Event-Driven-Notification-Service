package dispatcher

import (
	"Event-Driven-Notification-Service/channel"
	"Event-Driven-Notification-Service/event"
)

func DispatchEvent(evt event.RequestBody) []bool {
	switch evt.Event {
	case "user.registered":
		return []bool{channel.SendEmail(evt)}
	case "password.reset":
		return []bool{channel.SendEmail(evt), channel.SendInApp(evt)}
	case "order.completed":
		return []bool{channel.SendWebhook(evt)}
	default:
		return []bool{}
	}
}
