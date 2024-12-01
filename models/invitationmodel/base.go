package invitationmodel

import (
	"errors"
	"time"
)

var (
	ErrExpiredInviteCode  = errors.New("invite_code_expired")
	ErrInactiveInviteCode = errors.New("invite_code_not_active")
)

type InviteCode struct {
	Code           string `json:"code"`
	Limit          *uint  `json:"limit"`
	Active         bool   `json:"active"`
	ExpirationDate *int64 `json:"expiration_date"`
	Created        int64  `json:"created"`
	Updated        int64  `json:"updated"`
}

func (ic *InviteCode) TableName() string {
	return "app_invitation_codes"
}

func (ic *InviteCode) ValidateCode() error {
	if !ic.Active {
		return ErrInactiveInviteCode
	}

	if ic.ExpirationDate != nil && *ic.ExpirationDate < time.Now().Unix() {
		return ErrExpiredInviteCode
	}

	return nil
}
