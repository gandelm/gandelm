package etcd

import (
	"context"
	"sort"
	"time"

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

var _ repository.CatalogRORepository = (*CatalogRORepository)(nil)

type CatalogRORepository struct {
	db        client.Client
	container container.Containerer
}

func NewCatalogRORepository(container container.Containerer) *CatalogRORepository {
	return &CatalogRORepository{
		container: container,
		db:        container.Kubernetes(),
	}
}

var _ repository.CatalogRWRepository = (*CatalogRepository)(nil)

type CatalogRepository struct {
	*CatalogRORepository
}

func NewCatalogRepository(container container.Containerer) *CatalogRepository {
	return &CatalogRepository{
		CatalogRORepository: NewCatalogRORepository(container),
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

	now := time.Now()

	model.Spec = v1.GandelmCatalogSpec{
		ID:           catalog.WorkloadID,
		Name:         catalog.Name,
		Version:      catalog.Version,
		Description:  catalog.Description,
		Priority:     uint32(catalog.Priority),
		Labels:       catalog.Labels,
		ArtifactTags: catalog.ArtifactTags,
		CreatedAt:    metav1.Time{Time: now},
		UpdatedAt:    metav1.Time{Time: now},
	}

	return c.db.Create(ctx, &model)
}

// Update implements repository.CatalogRWRepository.
func (c *CatalogRepository) Update(ctx context.Context, catalog *entity.Catalog) error {
	model := v1.GandelmCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Name:      catalog.ID.String(),
			Namespace: c.container.Config().Namespace(),
		},
	}

	now := time.Now()

	model.Spec = v1.GandelmCatalogSpec{
		ID:           catalog.WorkloadID,
		Name:         catalog.Name,
		Version:      catalog.Version,
		Description:  catalog.Description,
		ArtifactTags: catalog.ArtifactTags,
		Priority:     uint32(catalog.Priority),
		UpdatedAt:    metav1.Time{Time: now},
	}

	return c.db.Update(ctx, &model)
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

// Find implements repository.CatalogRORepository.
func (c *CatalogRORepository) Find(ctx context.Context, id uuid.UUID) (*entity.Catalog, error) {
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
func (c *CatalogRORepository) FindAll(ctx context.Context) (entity.Catalogs, error) {
	catalogs := v1.GandelmCatalogList{}
	if err := c.db.List(ctx, &catalogs); err != nil {
		return nil, err
	}

	sort.SliceStable(catalogs.Items, func(i, j int) bool {
		return catalogs.Items[i].Spec.Priority > catalogs.Items[j].Spec.Priority
	})

	return converter.MakeCatalogs(catalogs), nil
}
