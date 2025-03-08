package usecase

import (
	"context"

	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"github.com/google/uuid"
)

type CatalogUpdator interface {
	Execute(ctx context.Context, input *CatalogUpdateInput) (*CatalogUpdateOutput, error)
}

type CatalogUpdateInput struct {
	ID          uuid.UUID
	Priority    int
	Description string
	Version     string
	Name        string
}

type CatalogUpdateOutput struct {
	Catalog *entity.Catalog
}

type CatalogUpdate struct {
	catalogRWRepository repository.CatalogRWRepository
}

func NewCatalogUpdate(
	catalogRWRepository repository.CatalogRWRepository,
) *CatalogUpdate {
	return &CatalogUpdate{
		catalogRWRepository: catalogRWRepository,
	}
}

func (c *CatalogUpdate) Execute(ctx context.Context, input *CatalogUpdateInput) (*CatalogUpdateOutput, error) {
	catalog := &entity.Catalog{
		ID:          input.ID,
		Name:        input.Name,
		Version:     input.Version,
		Description: input.Description,
		Priority:    input.Priority,
		WorkloadID:  entity.NewWorkloadID(input.Version, input.Name),
	}

	if err := c.catalogRWRepository.Update(ctx, catalog); err != nil {
		return nil, err
	}

	return &CatalogUpdateOutput{
		Catalog: catalog,
	}, nil
}
