package converter

import (
	catalogv1 "github.com/gandelm/gandelm/generated/protocol/catalog/v1"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
)

func MakeArtifactPb(artifact *entity.Artifact) *catalogv1.Artifact {
	return &catalogv1.Artifact{
		Name: artifact.Name,
		Type: catalogv1.ArtifactType(artifact.Type),
		Tag:  "",
	}
}

func MakeArtifactsPb(artifacts entity.Artifacts) []*catalogv1.Artifact {
	results := make([]*catalogv1.Artifact, 0, len(artifacts))
	for _, artifact := range artifacts {
		results = append(results, MakeArtifactPb(artifact))
	}

	return results
}
