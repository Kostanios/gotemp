package waitlistmodel

type WaitListUser struct {
	UserID  uint  `json:"user_id" gorm:"primarykey"`
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
}

func (w WaitListUser) TableName() string {
	return "app_waitlist"
}
