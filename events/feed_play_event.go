package events

type FeedPlayEvent struct {
	Header     EventHeader `json:"header"`
	UserID     int64       `json:"user_id"`
	IsPlusUser bool        `json:"is_plus_user"`
	PlusType   string      `json:"plus_type"`
	Duration   int32       `json:"duration"`
}

func NewFeedPlayEvent(corelationID string, userID int64, isPlus bool, plusType string, duration int32) FeedPlayEvent {
	return FeedPlayEvent{
		Header:     NewEventHeader(corelationID, "FeedPlayEvent"),
		UserID:     userID,
		IsPlusUser: isPlus,
		PlusType:   plusType,
		Duration:   duration,
	}
}
