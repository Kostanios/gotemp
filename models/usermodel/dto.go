package usermodel

import (
	"time"
)

type UserDto struct {
	ID            uint      `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Pronoun       string    `json:"pronoun"`
	CustomPronoun *string   `json:"custom_pronoun"`
	Birthday      time.Time `json:"birthday"`
	CreateDate    time.Time `json:"create_date"`
}

type PersonalDetailsDto struct {
	UserDto
	Phone                       string        `json:"phone"`
	Email                       string        `json:"email"`
	HasAccess                   bool          `json:"has_access"`
	IsProfileComplete           bool          `json:"is_profile_complete"`
	TwoFactorEnabled            bool          `json:"two_factor_enabled"`
	ShouldReceiveTextsAndEmails bool          `json:"should_receive_texts_and_emails"`
	Subscription                *Subscription `json:"subscription,omitempty"`
	HasUsedFreeTrial            bool          `json:"has_used_free_trial"`
}

func NewUserDto(user *AppUser) UserDto {
	return UserDto{
		ID:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Pronoun:       parsePronouns(user.Pronoun),
		CustomPronoun: user.CustomPronoun,
		Birthday:      user.Birthday,
		CreateDate:    time.Unix(user.Created, 0),
	}
}

func NewPersonalDetailsDto(user *AppUser) PersonalDetailsDto {
	return PersonalDetailsDto{
		UserDto:                     NewUserDto(user),
		Phone:                       user.Phone,
		Email:                       user.Email,
		HasAccess:                   user.HasAccess,
		IsProfileComplete:           user.IsProfileComplete,
		TwoFactorEnabled:            user.TwoFactorEnabled,
		ShouldReceiveTextsAndEmails: user.ShouldReceiveEmails && user.ShouldReceiveTexts,
		Subscription:                user.Subscription,
		HasUsedFreeTrial:            user.HasUsedFreeTrial,
	}
}
