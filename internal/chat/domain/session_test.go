package domain_test

import (
    "testing"
    "CustomerEngagementPlatform/internal/chat/domain"
)

func TestChatSession_AssignOperator(t *testing.T) {
    // Arrange
    session := domain.NewChatSession("web-123", "vis-456")

    // Act
    err := session.AssignOperator("op-789")

    // Assert
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if session.Status() != domain.StatusActive {
        t.Errorf("expected status ACTIVE, got %v", session.Status())
    }
    
    events := session.Events()
    if len(events) != 1 {
        t.Errorf("expected 1 event, got %d", len(events))
    }
}