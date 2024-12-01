package events

import "encoding/json"

type UserActivityEvent struct {
	Header EventHeader `json:"header"`
	UserID int64       `json:"user_id"`
	Route  string      `json:"route"`
}

func NewUserActivityEvent(correlationID string, userID int64, route string) UserActivityEvent {
	return UserActivityEvent{
		Header: NewEventHeader(correlationID, "UserActivityEvent"),
		UserID: userID,
		Route:  route,
	}
}

func (e UserActivityEvent) Deserialize() ([]byte, error) {
	return json.Marshal(e)
}
