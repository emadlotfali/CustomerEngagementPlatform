package domain

import "time"
//value objects

type SessionStatus string
const (
	StatusWaiting SessionStatus = "WAITING"
	StatusActive SessionStatus = "ACTIVE"
	StatusClosed SessionStatus = "CLOSED"
)
//Domin Event
type OperatorAssignedEvent struct{
	SessionID string
	OperatorID string
	OccurredOn time.Time
}
type MessageSentEvent struct{
	MessageID string
    SessionID string
    Content   string
    OccurredOn time.Time
}