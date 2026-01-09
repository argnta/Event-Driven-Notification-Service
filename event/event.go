package event

var EventQueue = make(chan RequestBody, 100)
var DeadLetterQueue = make(chan RequestBody, 100)