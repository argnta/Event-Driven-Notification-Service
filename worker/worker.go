package worker

import (
	"Event-Driven-Notification-Service/dispatcher"
	"Event-Driven-Notification-Service/event"
	"fmt"
)

const maxRetry = 3


func StartWorkers(n int) {
	for i := 0; i < n; i++ {
		go func(id int) {
			for req := range event.EventQueue {
				evt := event.Event{Body: req}
				processEvent(evt, id)
			}
		}(i)
	}
}

func processEvent(evt event.Event, workerID int) {
	fmt.Printf("Worker %d processing event: %s (id=%s, retry=%d)\n",
		workerID, evt.Body.Event, evt.ID, evt.RetryCount)

	results := dispatcher.DispatchEvent(evt.Body)

	success := true
	for _, r := range results {
		if !r {
			success = false
			break
		}
	}

	if success {
		fmt.Printf("Event %s processed successfully\n", evt.ID)
		return
	}

	evt.RetryCount++
	if evt.RetryCount < maxRetry {
		fmt.Printf("Retrying event %s (retry=%d)\n", evt.ID, evt.RetryCount)
		event.EventQueue <- evt.Body
	} else {
		fmt.Printf("Moving event %s to dead-letter queue\n", evt.ID)
		event.DeadLetterQueue <- evt.Body
	}
}
