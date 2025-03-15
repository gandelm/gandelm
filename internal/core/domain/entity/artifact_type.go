package entity

type ArtifactType = int32

const (
	ArtifactTypeNone      ArtifactType = iota
	ArtifactTypeContainer ArtifactType = iota
	ArtifactTypeAsset     ArtifactType = iota
)
