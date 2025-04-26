package models

import (
	"time"
)

type Category struct {
	ID          *string    `json:"id,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	RevisedAt   *time.Time `json:"revisedAt,omitempty"`
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
	Category    string     `json:"category,omitempty"`
}