package model

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID string `gorm:"primary_key"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Poem is sang by the bard.
type Poem struct {
	Model

	ShortID     string `gorm:"index"`
	Name        string
	ContentType string
	ExpiresIn   *time.Time
}

func NewPoem(name, contentType string) *Poem {
	id := uuid.New().String()
	p := &Poem{
		Model: Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ShortID:     id[32:],
		Name:        name,
		ContentType: contentType,
		ExpiresIn:   nil,
	}
	return p
}
