package storymodel

type Story struct {
	ID             string  `json:"id"`
	Type           string  `json:"type"`
	Title          string  `json:"title"`
	Image          string  `json:"image"`
	PlaybackID     string  `json:"playback_id"`
	Duration       float32 `json:"duration"`
	Thumbnail      string  `json:"thumbnail"`
	ButtonTitle    string  `json:"button_title"`
	ButtonURL      string  `json:"url"`
	HighlightImage string  `json:"highlight_image"`
}
