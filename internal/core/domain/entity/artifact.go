package entity

import (
	"time"

	constantv1 "github.com/gandelm/gandelm/generated/protocol/constant/v1"
)

type ArtifactType = constantv1.ArtifactType

type ArtifactID string
type Artifacts []*Artifact
type Artifact struct {
	ID          ArtifactID
	Title       string
	Description string
	Type        ArtifactType
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
