package domain
import(
	"errors"
	"time"
	"github.com/google/uuid"
)
//Message Entity
type Message struct{
	ID string
	SenderID string
	Content string
	CreatedAt time.Time
}
//ChatSession Aggregate Root
type ChatSession struct{
	id           string
    websiteID    string
    visitorID    string
    operatorID   *string
    status       SessionStatus
    domainEvents []interface{}
}
// Constructor (Factory Method)
func NewChatSession(websiteID, visitorID string) *ChatSession {
    return &ChatSession{
        id:        uuid.New().String(),
        websiteID: websiteID,
        visitorID: visitorID,
        status:    StatusWaiting,
    }
}
// Behaviors
func (s *ChatSession) AssignOperator(operatorID string) error {
    if s.status == StatusClosed {
        return errors.New("cannot assign operator to a closed session")
    }
    
    s.operatorID = &operatorID
    s.status = StatusActive
    
    s.domainEvents = append(s.domainEvents, OperatorAssignedEvent{
        SessionID:  s.id,
        OperatorID: operatorID,
        OccurredOn: time.Now(),
    })
    return nil
}
func (s *ChatSession) SendMessage(senderID, content string) (Message, error) {
    if s.status == StatusClosed {
        return Message{}, errors.New("session is closed")
    }

    msg := Message{
        ID:        uuid.New().String(),
        SenderID:  senderID,
        Content:   content,
        CreatedAt: time.Now(),
    }

    s.domainEvents = append(s.domainEvents, MessageSentEvent{
        MessageID:  msg.ID,
        SessionID:  s.id,
        Content:    content,
        OccurredOn: time.Now(),
    })

    return msg, nil
}

// Getters
func (s *ChatSession) ID() string { return s.id }
func (s *ChatSession) Status() SessionStatus { return s.status }
func (s *ChatSession) Events() []interface{} { return s.domainEvents }
func (s *ChatSession) ClearEvents() { s.domainEvents = nil }