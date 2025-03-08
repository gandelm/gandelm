package repository

import (
	"context"

	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/google/uuid"
)

type CatalogRORepository interface {
	FindAll(ctx context.Context) (entity.Catalogs, error)
	Find(ctx context.Context, id uuid.UUID) (*entity.Catalog, error)
}

type CatalogRWRepository interface {
	CatalogRORepository
	Create(ctx context.Context, catalog *entity.Catalog) error
	Update(ctx context.Context, catalog *entity.Catalog) error
	Delete(ctx context.Context, id uuid.UUID) error
}
