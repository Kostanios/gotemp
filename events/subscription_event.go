package events

import "encoding/json"

type Platform string

const (
	Ios     Platform = "ios"
	Android Platform = "android"
)

type SubscriptionEventType string

const (
	Monthly SubscriptionEventType = "monthly"
	Yearly  SubscriptionEventType = "yearly"
)

type SubscriptionEvent struct {
	Header           EventHeader `json:"header"`
	UserID           int64       `json:"user_id"`
	Platform         string      `json:"platform"`
	SubscriptionType string      `json:"subscription_type"`
	CreateTimestamp  int64       `json:"create_timestamp"`
	ExpireTimestamp  int64       `json:"expire_timestamp"`
}

func NewSubscriptionEvent(
	correlationID string,
	userID int64,
	platform Platform,
	subscriptionType SubscriptionEventType,
	createTimestamp int64,
	expireTimestamp int64,
) SubscriptionEvent {
	return SubscriptionEvent{
		Header:           NewEventHeader(correlationID, "SubscriptionEvent"),
		UserID:           userID,
		Platform:         string(platform),
		SubscriptionType: string(subscriptionType),
		CreateTimestamp:  createTimestamp,
		ExpireTimestamp:  expireTimestamp,
	}
}

func (e SubscriptionEvent) Deserialize() ([]byte, error) {
	return json.Marshal(e)
}
