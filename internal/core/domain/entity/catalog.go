package entity

import (
	"time"

	"github.com/google/uuid"
)

type Catalogs []*Catalog
type Catalog struct {
	ID          uuid.UUID
	Name        string
	Version     string
	Description string
	Priority    int
	WorkloadID  WorkloadID

	Artifacts ArtifactTag

	CreatedAt time.Time
	UpdatedAt time.Time
}
