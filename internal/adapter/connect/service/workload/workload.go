package workload

import (
	"context"

	"connectrpc.com/connect"
	workloadv1 "github.com/gandelm/gandelm/generated/protocol/workload/v1"
	"github.com/gandelm/gandelm/generated/protocol/workload/v1/workloadv1connect"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/provider"
)

var _ workloadv1connect.WorkloadServiceHandler = (*WorkloadService)(nil)

func NewWorkloadService(container container.Containerer) *WorkloadService {
	return &WorkloadService{container: container}
}

type WorkloadService struct {
	container container.Containerer
}

// Get implements workloadv1connect.WorkloadServiceHandler.
func (w *WorkloadService) Get(ctx context.Context, req *connect.Request[workloadv1.GetRequest]) (*connect.Response[workloadv1.GetResponse], error) {
	workload, err := provider.NewWorkloadRORepository(w.container).Find(ctx, req.Msg.CatalogId, req.Msg.WorkloadId)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&workloadv1.GetResponse{
		Workload: converter.MakeWorkloadPb(workload),
	}), nil
}

// List implements workloadv1connect.WorkloadServiceHandler.
func (w *WorkloadService) List(ctx context.Context, req *connect.Request[workloadv1.ListRequest]) (*connect.Response[workloadv1.ListResponse], error) {
	workloads, err := provider.NewWorkloadRORepository(w.container).FindAll(ctx, req.Msg.CatalogId)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&workloadv1.ListResponse{
		Workloads: converter.MakeWorkloadsPb(workloads),
	}), nil
}
