package savedmodel

import (
	"encoding/json"
	"time"
)

const (
	SavedTypeBoost        = "boost"
	SavedTypeEpisode      = "episode"
	SavedTypeCollection   = "collection"
	SavedTypeHabitEpisode = "habitEpisode"
	SavedTypeEpisodeStory = "episodeStory"
)

type Saved struct {
	ID       string  `json:"id" gorm:"primarykey"`
	Type     string  `json:"type" gorm:"primarykey"`
	UserID   uint    `json:"user_id" gorm:"primarykey"`
	Metadata *string `json:"metadata"`
	Created  int64   `json:"created"`
	Updated  int64   `json:"updated"`
}

func NewUserSave(id string, saveType string, userID uint, metadata json.RawMessage) Saved {
	saved := Saved{
		ID:      id,
		Type:    saveType,
		UserID:  userID,
		Created: time.Now().Unix(),
		Updated: time.Now().Unix(),
	}

	if metadata != nil {
		data := string(metadata)
		saved.Metadata = &data
	}

	return saved
}

func (s Saved) TableName() string {
	return "app_user_saved"
}
