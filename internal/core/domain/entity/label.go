package entity

import (
	"time"
)

type Labels []*Label
type Label struct {
	ID          string
	Title       string
	Description string
	Color       string

	CreatedAt time.Time
	UpdatedAt time.Time
}
