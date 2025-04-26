package models

import (
	"time"
)

type Blog struct {
	ID          *string    `json:"id,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	RevisedAt   *time.Time `json:"revisedAt,omitempty"`
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Body        *string    `json:"body,omitempty"`
	Tag         []Tag      `json:"tag,omitempty"`
	Category    *Category   `json:"category,omitempty"`
}
