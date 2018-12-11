package user

import (
	"errors"
	"time"
)

var (
	ErrUsernameAlreadyTaken   = errors.New("This username already taken")
	ErrEmailAlreadyRegistered = errors.New("This email already registered")
	ErrUserNotFound           = errors.New("User not found")
)

type User struct {
	UUID     string     `json:"uuid"`
	Name     string     `json:"name"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password,omitempty"`
	Created  *time.Time `json:"dtm_crt,omitempty"`
	Modified *time.Time `json:"dtm_upd,omitempty"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AvailableUser struct {
	Username     string `json:"username,omitempty"`
	IsRegistered bool   `json:"is_registered"`
}
