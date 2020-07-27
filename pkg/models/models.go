package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found.")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

// Snippet model struct
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User model struct
type User struct {
	ID       int
	Name     string
	Email    string
	Password []byte
	Created  time.Time
}
