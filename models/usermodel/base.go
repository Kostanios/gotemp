package usermodel

import (
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type AppUser struct {

	// Basic information
	ID uint `json:"id" gorm:"primarykey"`

	FirstName           string        `json:"first_name"`
	LastName            string        `json:"last_name"`
	Phone               string        `json:"phone"`
	Email               string        `json:"email"`
	IsEmailVerified     bool          `json:"is_email_verified"`
	Pronoun             UserPronounce `json:"pronoun"`
	CustomPronoun       *string       `json:"custom_pronoun"`
	Birthday            time.Time     `json:"birthday"`
	PushToken           *string       `json:"push_token"`
	InviteCode          *string       `json:"invite_code"`
	HasAccess           bool          `json:"has_access"`
	IsProfileComplete   bool          `json:"is_profile_complete"`
	ShouldReceiveTexts  bool          `json:"should_receive_texts"`
	ShouldReceiveEmails bool          `json:"should_receive_emails"`
	HasUsedFreeTrial    bool          `json:"has_used_free_trial"`

	Password         *string `json:"password,omitempty"`
	Salt             *string `json:"salt,omitempty"`
	TwoFactorEnabled bool    `json:"two_factor_enabled"`

	Subscription *Subscription `json:"subscription" gorm:"foreignKey:user_id"`

	// Timestamps
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
}

func (user *AppUser) TableName() string {
	return "app_users"
}

func (user *AppUser) SetSubscription(newSubscription *Subscription) {
	user.Subscription = newSubscription
}

func (user *AppUser) SetUserData(firstname string, lastname string, email string, pronounce string, customPronounce *string) error {
	userPronoun, err := getPronouns(strings.ToLower(pronounce))
	if err != nil {
		return err
	}
	err = user.setUserEmail(email)
	if err != nil {
		return err
	}
	user.FirstName = firstname
	user.LastName = lastname
	user.Pronoun = userPronoun
	user.IsProfileComplete = true
	if userPronoun == Other {
		user.CustomPronoun = customPronounce
	} else {
		user.CustomPronoun = nil
	}
	return nil
}

func (user *AppUser) SetUserBirthday(birthday time.Time) error {
	user.Birthday = birthday
	return nil
}

func (user *AppUser) setUserEmail(newEmail string) error {
	isValidEmail := govalidator.IsEmail(newEmail)
	if !isValidEmail {
		return ErrInvalidEmail
	}

	if user.Email != strings.ToLower(newEmail) {
		user.Email = strings.ToLower(newEmail)
		user.IsEmailVerified = false
	}
	return nil
}

func (user *AppUser) ShouldSendVerificationEmail() bool {
	return user.IsEmailVerified
}

func (user *AppUser) SetShouldReceiveNotifications(value bool) {
	user.ShouldReceiveTexts = value
	user.ShouldReceiveEmails = value
}

func (user *AppUser) ShouldSendEmail() bool {
	return user.ShouldReceiveEmails && user.IsEmailVerified
}

func (user *AppUser) ShouldSendText() bool {
	return user.ShouldReceiveTexts
}

type Subscription struct {
	UserID      uint             `json:"user_id" gorm:"primarykey"`
	Type        SubscriptionType `json:"type"`
	Status      string           `json:"status"`
	AutoRenew   bool             `json:"auto_renew"`
	Ending      *int             `json:"ending"`
	IsFreeTrial bool             `json:"is_free_trial"`

	// Apple related
	AppleTransactionId *string `json:"-"`

	// Google Related
	GoogleTransactionId *string `json:"-"`
}

func NewGoogleSubscription(userID uint, ending int, subscriptionType string, googleSubscriptionID string) (*Subscription, error) {
	subType, err := getSubscribeType(subscriptionType)
	if err != nil {
		return nil, err
	}

	newSubscription := &Subscription{
		UserID:              userID,
		Type:                subType,
		Status:              SubscriptionStatusActive,
		AutoRenew:           false,
		Ending:              &ending,
		GoogleTransactionId: &googleSubscriptionID,
	}

	return newSubscription, nil
}

func (sub *Subscription) TableName() string {
	return "app_users_subscriptions"
}

func (sub *Subscription) IsValid() bool {
	if sub.Ending == nil {
		return false
	}

	if int64(*sub.Ending) < time.Now().UTC().Unix() {
		return false
	}

	return true
}

func getSubscribeType(subType string) (SubscriptionType, error) {
	switch subType {
	case string(Monthly):
		return SubscriptionTypeMonthly, nil
	case string(Yearly):
		return SubscriptionTypeYearly, nil
	default:
		return "", ErrInvalidSubscription
	}
}

func getPronouns(pronounce string) (UserPronounce, error) {
	switch pronounce {
	case "he":
		return He, nil
	case "she":
		return She, nil
	case "they":
		return They, nil
	case "other":
		return Other, nil
	case "prefer_not_to_say":
		return PreferNotToSay, nil
	}
	return Other, ErrInvalidPronounce
}

func parsePronouns(pronounce UserPronounce) string {
	switch pronounce {
	case He:
		return "he"
	case She:
		return "she"
	case They:
		return "they"
	case Other:
		return "other"
	case PreferNotToSay:
		return "prefer_not_to_say"
	}
	return ""
}
