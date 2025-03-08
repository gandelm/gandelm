package etcd

import (
	"context"

	v1 "github.com/gandelm/gandelm/api/v1"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ repository.CatalogRORepository = (*CatalogRepository)(nil)
var _ repository.CatalogRWRepository = (*CatalogRepository)(nil)

type CatalogRepository struct {
	db        client.Client
	container container.Containerer
}

func NewCatalogRepository(container container.Containerer) *CatalogRepository {
	return &CatalogRepository{
		container: container,
		db:        container.Kubernetes(),
	}
}

// Create implements repository.CatalogRWRepository.
func (c *CatalogRepository) Create(ctx context.Context, catalog *entity.Catalog) error {
	model := v1.GandelmCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Name:       catalog.ID.String(),
			Namespace:  c.container.Config().Namespace(),
			Finalizers: []string{"gandelm.com/finalizer"},
		},
	}

	model.Spec = v1.GandelmCatalogSpec{
		ID:          catalog.WorkloadID,
		Name:        catalog.Name,
		Version:     catalog.Version,
		Description: catalog.Description,
		Priority:    uint32(catalog.Priority),
	}

	if err := c.db.Create(ctx, &model); err != nil {
		return err
	}

	return nil
}

// Delete implements repository.CatalogRWRepository.
func (c *CatalogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// Update implements repository.CatalogRWRepository.
func (c *CatalogRepository) Update(ctx context.Context, catalog *entity.Catalog) error {
	panic("unimplemented")
}

// Find implements repository.CatalogRORepository.
func (c *CatalogRepository) Find(ctx context.Context, id uuid.UUID) (*entity.Catalog, error) {
	panic("unimplemented")
}

// FindAll implements repository.CatalogRORepository.
func (c *CatalogRepository) FindAll(ctx context.Context) (entity.Catalogs, error) {
	panic("unimplemented")
}
