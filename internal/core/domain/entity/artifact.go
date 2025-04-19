package entity

import artifactv1 "github.com/gandelm/gandelm/generated/protocol/artifact/v1"

type ArtifactType = artifactv1.ArtifactType

type ArtifactID string
type Artifacts []*Artifact
type Artifact struct {
	ID          ArtifactID
	Title       string
	Description string
	Type        ArtifactType
}
