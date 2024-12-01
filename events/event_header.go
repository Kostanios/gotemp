package events

import (
	"time"

	"github.com/google/uuid"
)

type EventHeader struct {
	ID              string `json:"id"`
	EventName       string `json:"event_name"`
	CorelationID    string `json:"corelation_id"`
	CreateTimestamp int64  `json:"create_timestamp"`
}

func NewEventHeader(corelationID string, eventName string) EventHeader {
	return EventHeader{
		ID:              uuid.New().String(),
		EventName:       eventName,
		CorelationID:    corelationID,
		CreateTimestamp: time.Now().Unix(),
	}
}
