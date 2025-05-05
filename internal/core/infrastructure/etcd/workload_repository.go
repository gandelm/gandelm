package etcd

import (
	"context"

	v1 "github.com/gandelm/gandelm/api/v1"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ repository.WorkloadRORepository = (*WorkloadRORepository)(nil)

type WorkloadRORepository struct {
	db        client.Client
	container container.Containerer
}

// Find implements repository.WorkloadRORepository.
func (w *WorkloadRORepository) Find(ctx context.Context, namespace string, id entity.WorkloadID) (*entity.Workload, error) {
	workload := &v1.GandelmWorkload{}
	if err := w.db.Get(ctx, types.NamespacedName{
		Namespace: namespace,
		Name:      id,
	}, workload); err != nil {
		return nil, err
	}

	return converter.MakeWorkload(workload), nil
}

// FindAll implements repository.WorkloadRORepository.
func (w *WorkloadRORepository) FindAll(ctx context.Context, namespace string) (entity.Workloads, error) {
	list := &v1.GandelmWorkloadList{}
	if err := w.db.List(ctx, list, client.InNamespace(namespace)); err != nil {
		return nil, err
	}

	return converter.MakeWorkloads(list.Items), nil
}

func NewWorkloadRORepository(container container.Containerer) *WorkloadRORepository {
	return &WorkloadRORepository{
		container: container,
		db:        container.Kubernetes(),
	}
}
