package model

import (
	"time"
)

type Url struct {
	Hash        string
	Link        string
	CreatedAt   time.Time
	ExpiresAt   time.Time
	ClickCounts int
}
