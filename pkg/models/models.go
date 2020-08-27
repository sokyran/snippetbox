package models

import (
	"errors"
	"time"
)

// ErrNoRecord made for using when no record was found
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is the model for snippets inside database
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
