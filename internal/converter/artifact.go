package converter

import (
	v1 "github.com/gandelm/gandelm/api/v1"
	artifactv1 "github.com/gandelm/gandelm/generated/protocol/artifact/v1"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MakeArtifactPb(artifact *entity.Artifact) *artifactv1.Artifact {
	return &artifactv1.Artifact{
		Id:          string(artifact.ID),
		Title:       artifact.Title,
		Description: artifact.Description,
		Type:        artifact.Type,
		CreatedAt:   timestamppb.New(artifact.CreatedAt),
		UpdatedAt:   timestamppb.New(artifact.UpdatedAt),
	}
}

func MakeArtifactsPb(artifacts entity.Artifacts) []*artifactv1.Artifact {
	results := make([]*artifactv1.Artifact, 0, len(artifacts))
	for _, artifact := range artifacts {
		results = append(results, MakeArtifactPb(artifact))
	}

	return results
}

func MakeArtifact(artifact *v1.GandelmArtifact) *entity.Artifact {
	return &entity.Artifact{
		ID:          entity.ArtifactID(artifact.Name),
		Title:       artifact.Spec.Title,
		Description: artifact.Spec.Description,
		Type:        entity.ArtifactType(artifact.Spec.Type),
		CreatedAt:   artifact.Spec.CreatedAt.Time,
		UpdatedAt:   artifact.Spec.UpdatedAt.Time,
	}
}

func MakeArtifacts(artifacts v1.GandelmArtifactList) entity.Artifacts {
	results := make(entity.Artifacts, 0, len(artifacts.Items))
	for _, artifact := range artifacts.Items {
		results = append(results, MakeArtifact(&artifact))
	}

	return results
}
