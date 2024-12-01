package events

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_NewSubscriptionEvent(t *testing.T) {
	newEvent := NewSubscriptionEvent(uuid.New().String(), 1, Android, Monthly, 10, 20)
	bytes, err := newEvent.Deserialize()
	assert.NoError(t, err)
	existingEvent := new(SubscriptionEvent)
	err = json.Unmarshal(bytes, existingEvent)
	assert.NoError(t, err)
	assert.Equal(t, newEvent.UserID, existingEvent.UserID)
	assert.Equal(t, newEvent.CreateTimestamp, existingEvent.CreateTimestamp)
	assert.Equal(t, newEvent.ExpireTimestamp, existingEvent.ExpireTimestamp)
	assert.Equal(t, newEvent.Platform, existingEvent.Platform)
	assert.Equal(t, newEvent.SubscriptionType, existingEvent.SubscriptionType)
}
