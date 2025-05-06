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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ repository.ArtifactRORepository = (*ArtifactRORepository)(nil)

type ArtifactRORepository struct {
	db        client.Client
	container container.Containerer
}

func NewArtifactRORepository(container container.Containerer) *ArtifactRORepository {
	return &ArtifactRORepository{
		container: container,
		db:        container.Kubernetes(),
	}
}

var _ repository.ArtifactRWRepository = (*ArtifactRepository)(nil)

type ArtifactRepository struct {
	*ArtifactRORepository
}

func NewArtifactRepository(container container.Containerer) *ArtifactRepository {
	return &ArtifactRepository{
		ArtifactRORepository: NewArtifactRORepository(container),
	}
}

// Find implements repository.ArtifactRORepository.
func (a *ArtifactRORepository) Find(ctx context.Context, id entity.ArtifactID) (*entity.Artifact, error) {
	artifact := &v1.GandelmArtifact{}
	if err := a.db.Get(ctx, client.ObjectKey{
		Name:      string(id),
		Namespace: a.container.Config().Namespace(),
	}, artifact); err != nil {
		return nil, err
	}

	return converter.MakeArtifact(artifact), nil
}

// FindAll implements repository.ArtifactRORepository.
func (a *ArtifactRORepository) FindAll(ctx context.Context) (entity.Artifacts, error) {
	artifacts := v1.GandelmArtifactList{}
	if err := a.db.List(ctx, &artifacts); err != nil {
		return nil, err
	}

	sort.Slice(artifacts.Items, func(i, j int) bool {
		return artifacts.Items[i].Name < artifacts.Items[j].Name
	})

	return converter.MakeArtifacts(artifacts), nil
}

// Create implements repository.ArtifactRWRepository.
func (a *ArtifactRepository) Create(ctx context.Context, artifact *entity.Artifact) error {
	model := v1.GandelmArtifact{
		ObjectMeta: metav1.ObjectMeta{
			Name:      string(artifact.ID),
			Namespace: a.container.Config().Namespace(),
		},
	}

	now := time.Now()

	model.Spec = v1.GandelmArtifactSpec{
		Title:       artifact.Title,
		Description: artifact.Description,
		Type:        int(artifact.Type),
		CreatedAt:   metav1.Time{Time: now},
		UpdatedAt:   metav1.Time{Time: now},
	}

	return a.db.Create(ctx, &model)
}

// Delete implements repository.ArtifactRWRepository.
func (a *ArtifactRepository) Delete(ctx context.Context, id entity.ArtifactID) error {
	panic("unimplemented")
}

// Update implements repository.ArtifactRWRepository.
func (a *ArtifactRepository) Update(ctx context.Context, artifact *entity.Artifact) error {
	panic("unimplemented")
}
