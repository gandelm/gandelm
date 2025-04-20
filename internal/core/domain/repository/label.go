package repository

import (
	"context"

	"github.com/gandelm/gandelm/internal/core/domain/entity"
)

type LabelRORepository interface {
	Find(ctx context.Context, id string) (*entity.Label, error)
	FindAll(ctx context.Context) (entity.Labels, error)
}

type LabelRWRepository interface {
	LabelRORepository
	Create(ctx context.Context, label *entity.Label) error
	// Update(ctx context.Context, label *entity.Label) error
	Delete(ctx context.Context, name string) error
}
