package repository

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
)

type DeploymentRORepository interface {
	FindAll(ctx context.Context, namespace string) (*appsv1.DeploymentList, error)
}
