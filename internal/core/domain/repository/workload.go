package repository

import (
	"context"

	"github.com/gandelm/gandelm/internal/core/domain/entity"
)

type WorkloadRORepository interface {
	FindAll(ctx context.Context, namespace string) (entity.Workloads, error)
	Find(ctx context.Context, namespace string, id entity.WorkloadID) (*entity.Workload, error)
}
