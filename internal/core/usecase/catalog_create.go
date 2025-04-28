package usecase

import (
	"context"

	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"github.com/google/uuid"
)

type CatalogCreator interface {
	Execute(ctx context.Context, input *CatalogCreateInput) (*CatalogCreateOutput, error)
}

type CatalogCreateInput struct {
	Priority    int
	Description string
	Version     string
	Name        string
	Labels      []string
}

type CatalogCreateOutput struct {
	Catalog *entity.Catalog
}

type CatalogCreate struct {
	catalogRWRepository repository.CatalogRWRepository
}

func NewCatalogCreate(
	catalogRWRepository repository.CatalogRWRepository,
) *CatalogCreate {
	return &CatalogCreate{
		catalogRWRepository: catalogRWRepository,
	}
}

func (c *CatalogCreate) Execute(ctx context.Context, input *CatalogCreateInput) (*CatalogCreateOutput, error) {
	catalog := &entity.Catalog{
		ID:          uuid.New(),
		Name:        input.Name,
		Version:     input.Version,
		Description: input.Description,
		Priority:    input.Priority,
		Labels:      input.Labels,
		WorkloadID:  entity.NewWorkloadID(input.Version, input.Name),
	}

	if err := c.catalogRWRepository.Create(ctx, catalog); err != nil {
		return nil, err
	}

	return &CatalogCreateOutput{
		Catalog: catalog,
	}, nil
}
