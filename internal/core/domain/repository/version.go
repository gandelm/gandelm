package repository

import (
	"context"

	"github.com/gandelm/gandelm/internal/core/domain/entity"
)

type VersionRORepository interface {
	Find(ctx context.Context, id string) (*entity.Version, error)
	FindAll(ctx context.Context) (entity.Versions, error)
}

type VersionRWRepository interface {
	VersionRORepository
	Create(ctx context.Context, version *entity.Version) error
	// Update(ctx context.Context, version *entity.Version) error
	Delete(ctx context.Context, name string) error
}
