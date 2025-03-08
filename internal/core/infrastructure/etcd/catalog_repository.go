package etcd

import (
	"context"

	v1 "github.com/gandelm/gandelm/api/v1"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
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

	return c.db.Create(ctx, &model)
}

// Delete implements repository.CatalogRWRepository.
func (c *CatalogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return c.db.Delete(ctx, &v1.GandelmCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Name:      id.String(),
			Namespace: c.container.Config().Namespace(),
		},
	})
}

// Update implements repository.CatalogRWRepository.
func (c *CatalogRepository) Update(ctx context.Context, catalog *entity.Catalog) error {
	model := v1.GandelmCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Name:      catalog.ID.String(),
			Namespace: c.container.Config().Namespace(),
		},
	}

	model.Spec = v1.GandelmCatalogSpec{
		ID:          catalog.WorkloadID,
		Name:        catalog.Name,
		Version:     catalog.Version,
		Description: catalog.Description,
		Priority:    uint32(catalog.Priority),
	}

	return c.db.Update(ctx, &model)
}

// Find implements repository.CatalogRORepository.
func (c *CatalogRepository) Find(ctx context.Context, id uuid.UUID) (*entity.Catalog, error) {
	catalog := &v1.GandelmCatalog{}
	if err := c.db.Get(ctx, types.NamespacedName{
		Namespace: c.container.Config().Namespace(),
		Name:      id.String(),
	}, catalog); err != nil {
		return nil, err
	}

	return converter.MakeCatalog(catalog), nil
}

// FindAll implements repository.CatalogRORepository.
func (c *CatalogRepository) FindAll(ctx context.Context) (entity.Catalogs, error) {
	catalogs := v1.GandelmCatalogList{}
	if err := c.db.List(ctx, &catalogs); err != nil {
		return nil, err
	}

	return converter.MakeCatalogs(catalogs), nil
}
