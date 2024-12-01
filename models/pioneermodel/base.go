package pioneermodel

type PioneerNewsSection struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Body        string `json:"body"`
	Title       string `json:"title"`
	Created     int64  `json:"created"`
	ButtonTitle string `json:"button_title"`
}
