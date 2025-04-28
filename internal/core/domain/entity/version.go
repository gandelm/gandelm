package entity

import "time"

type Versions []*Version
type Version struct {
	ID            string
	BaseArtifacts map[string]string
	BasePriority  uint32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
