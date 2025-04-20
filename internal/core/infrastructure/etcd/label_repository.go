package etcd

import (
	"context"

	v1 "github.com/gandelm/gandelm/api/v1"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ repository.LabelRORepository = (*LabelRORepository)(nil)

type LabelRORepository struct {
	db        client.Client
	container container.Containerer
}

func NewLabelRORepository(container container.Containerer) *LabelRORepository {
	return &LabelRORepository{
		container: container,
		db:        container.Kubernetes(),
	}
}

var _ repository.LabelRWRepository = (*LabelRepository)(nil)

type LabelRepository struct {
	*LabelRORepository
}

func NewLabelRepository(container container.Containerer) *LabelRepository {
	return &LabelRepository{
		LabelRORepository: NewLabelRORepository(container),
	}
}

// Find implements repository.LabelRORepository.
func (l *LabelRORepository) Find(ctx context.Context, id string) (*entity.Label, error) {
	label := &v1.GandelmLabel{}
	if err := l.db.Get(ctx, client.ObjectKey{
		Namespace: l.container.Config().Namespace(),
		Name:      id,
	}, label); err != nil {
		return nil, err
	}

	return converter.MakeLabel(label), nil
}

// FindAll implements repository.LabelRORepository.
func (l *LabelRORepository) FindAll(ctx context.Context) (entity.Labels, error) {
	labels := &v1.GandelmLabelList{}
	if err := l.db.List(ctx, labels, client.InNamespace(l.container.Config().Namespace())); err != nil {
		return nil, err
	}

	return converter.MakeLabels(labels), nil
}

// Create implements repository.LabelRWRepository.
func (l *LabelRepository) Create(ctx context.Context, label *entity.Label) error {
	model := v1.GandelmLabel{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: l.container.Config().Namespace(),
			Name:      label.ID,
		},
	}

	model.Spec = v1.GandelmLabelSpec{
		Title:       label.Title,
		Description: label.Description,
		Color:       label.Color,
	}

	return l.db.Create(ctx, &model)
}

// Delete implements repository.LabelRWRepository.
func (l *LabelRepository) Delete(ctx context.Context, labelID string) error {
	return l.db.Delete(ctx, &v1.GandelmLabel{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: l.container.Config().Namespace(),
			Name:      labelID,
		},
	})
}
