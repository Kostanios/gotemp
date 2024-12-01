package usermodel

import "errors"

var (
	ErrInvalidSubscription = errors.New("invalid subscription type")
	ErrInvalidPronounce    = errors.New("invalid pronounce")
	ErrInvalidEmail        = errors.New("invalid email")
)
