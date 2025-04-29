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
	Tags         []Tag      `json:"tags,omitempty"`
	Category    *Category   `json:"category,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Overview 	*string    `json:"overview,omitempty"`
}
