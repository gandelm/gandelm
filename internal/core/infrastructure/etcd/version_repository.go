package etcd

import (
	"context"
	"sort"

	v1 "github.com/gandelm/gandelm/api/v1"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ repository.VersionRORepository = (*VersionRORepository)(nil)

type VersionRORepository struct {
	db        client.Client
	container container.Containerer
}

func NewVersionRORepository(container container.Containerer) *VersionRORepository {
	return &VersionRORepository{
		container: container,
		db:        container.Kubernetes(),
	}
}

// Find implements repository.VersionRORepository.
func (v *VersionRORepository) Find(ctx context.Context, id string) (*entity.Version, error) {
	panic("unimplemented")
}

// FindAll implements repository.VersionRORepository.
func (v *VersionRORepository) FindAll(ctx context.Context) (entity.Versions, error) {
	versions := &v1.GandelmVersionList{}
	if err := v.db.List(ctx, versions, client.InNamespace(v.container.Config().Namespace())); err != nil {
		return nil, err
	}

	sort.Slice(versions.Items, func(i, j int) bool { return versions.Items[i].Spec.BasePriority < versions.Items[j].Spec.BasePriority })

	return converter.MakeVersions(versions), nil
}

var _ repository.VersionRWRepository = (*VersionRepository)(nil)

type VersionRepository struct {
	*VersionRORepository
}

func NewVersionRepository(container container.Containerer) *VersionRepository {
	return &VersionRepository{
		VersionRORepository: NewVersionRORepository(container),
	}
}

// Create implements repository.VersionRWRepository.
func (v *VersionRepository) Create(ctx context.Context, version *entity.Version) error {
	model := v1.GandelmVersion{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: v.container.Config().Namespace(),
			Name:      version.ID,
		},
	}

	model.Spec = v1.GandelmVersionSpec{
		BaseArtifacts: make(map[string]string),
		BasePriority:  version.BasePriority,
		CreatedAt:     metav1.NewTime(version.CreatedAt),
		UpdatedAt:     metav1.NewTime(version.UpdatedAt),
	}

	return v.db.Create(ctx, &model)
}

// Delete implements repository.VersionRWRepository.
func (v *VersionRepository) Delete(ctx context.Context, name string) error {
	panic("unimplemented")
}
