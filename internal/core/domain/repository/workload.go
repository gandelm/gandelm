package repository

import (
	"context"

	"github.com/gandelm/gandelm/internal/core/domain/entity"
)

type WorkloadRORepository interface {
	FindAll(ctx context.Context) (entity.Workloads, error)
	Find(ctx context.Context, id entity.WorkloadID) (*entity.Workload, error)
}

type WorkloadRWRepository interface {
	WorkloadRORepository
	Create(ctx context.Context, workload *entity.Workload) error
	Update(ctx context.Context, workload *entity.Workload) error
	Delete(ctx context.Context, id entity.WorkloadID) error
}
