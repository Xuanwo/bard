package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Poem is sang by the bard.
type Poem struct {
	gorm.Model

	ID          string
	Type        string
	ContentType string
	CreatedAt   time.Time
	ExpiresIn   time.Time
}
