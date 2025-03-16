package etcd

import (
	"context"

	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ repository.DeploymentRORepository = (*DeploymentRORepository)(nil)

type DeploymentRORepository struct {
	db        client.Client
	container container.Containerer
}

func NewDeploymentRORepository(container container.Containerer) *DeploymentRORepository {
	return &DeploymentRORepository{
		db:        container.Kubernetes(),
		container: container,
	}
}

// FindAll implements repository.DeploymentRORepository.
func (d *DeploymentRORepository) FindAll(ctx context.Context, namespace string) (*appsv1.DeploymentList, error) {
	list := &appsv1.DeploymentList{}
	if err := d.db.List(ctx, list, client.InNamespace(namespace)); err != nil {
		return nil, err
	}

	return list, nil
}
