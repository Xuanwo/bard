package model

import (
	"time"

	"github.com/Xuanwo/bard/utils"
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
	t := time.Now()
	p := &Poem{
		Model: Model{
			ID:        utils.GenerateID(),
			CreatedAt: t,
			UpdatedAt: t,
		},
		ShortID:     utils.GenerateShortID(),
		Name:        name,
		ContentType: contentType,
		ExpiresIn:   nil,
	}
	return p
}
