package repository

import (
	"context"

	"github.com/gandelm/gandelm/internal/core/domain/entity"
)

type ArtifactRORepository interface {
	Find(ctx context.Context, id entity.ArtifactID) (entity.Artifacts, error)
	FindAll(ctx context.Context) (entity.Artifacts, error)
}

type ArtifactRWRepository interface {
	ArtifactRORepository
	Create(ctx context.Context, artifact *entity.Artifact) error
	Update(ctx context.Context, artifact *entity.Artifact) error
	Delete(ctx context.Context, id entity.ArtifactID) error
}
