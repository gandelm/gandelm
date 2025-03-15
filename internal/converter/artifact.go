package converter

import (
	catalogv1 "github.com/gandelm/gandelm/generated/protocol/catalog/v1"
	workloadv1 "github.com/gandelm/gandelm/generated/protocol/workload/v1"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
)

func MakeArtifactWithObjectPb(artifact *entity.Artifact) *workloadv1.ArtifactWithObject {
	return &workloadv1.ArtifactWithObject{
		Artifact: MakeArtifactPb(artifact),
		Version:  artifact.Version,
		Url:      artifact.URL,
	}
}

func MakeArtifactsWithObjectPb(artifact entity.Artifacts) []*workloadv1.ArtifactWithObject {
	results := make([]*workloadv1.ArtifactWithObject, 0, len(artifact))
	for _, artifact := range artifact {
		results = append(results, MakeArtifactWithObjectPb(artifact))
	}

	return results
}

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
