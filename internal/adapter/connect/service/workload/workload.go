package workload

import (
	"context"

	"connectrpc.com/connect"
	workloadv1 "github.com/gandelm/gandelm/generated/protocol/workload/v1"
	"github.com/gandelm/gandelm/generated/protocol/workload/v1/workloadv1connect"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"github.com/gandelm/gandelm/internal/provider"
)

var _ workloadv1connect.WorkloadServiceHandler = (*WorkloadService)(nil)

func NewWorkloadService(container container.Containerer) *WorkloadService {
	return &WorkloadService{
		workloadRORepository:   provider.NewWorkloadRORepository(container),
		deploymentRORepository: provider.NewDeploymentRORepository(container),
	}
}

type WorkloadService struct {
	workloadRORepository   repository.WorkloadRORepository
	deploymentRORepository repository.DeploymentRORepository
}

// Get implements workloadv1connect.WorkloadServiceHandler.
func (w *WorkloadService) Get(ctx context.Context, req *connect.Request[workloadv1.GetRequest]) (*connect.Response[workloadv1.GetResponse], error) {
	workload, err := w.workloadRORepository.Find(ctx, req.Msg.CatalogId, req.Msg.WorkloadId)
	if err != nil {
		return nil, err
	}

	deploymentList, err := w.deploymentRORepository.FindAll(ctx, req.Msg.CatalogId)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&workloadv1.GetResponse{
		Workload:    converter.MakeWorkloadPb(workload),
		Deployments: converter.MakeDeploymentsPb(deploymentList.Items),
	}), nil
}

// List implements workloadv1connect.WorkloadServiceHandler.
func (w *WorkloadService) List(ctx context.Context, req *connect.Request[workloadv1.ListRequest]) (*connect.Response[workloadv1.ListResponse], error) {
	workloads, err := w.workloadRORepository.FindAll(ctx, req.Msg.CatalogId)
	if err != nil {
		return nil, err
	}

	deploymentList, err := w.deploymentRORepository.FindAll(ctx, req.Msg.CatalogId)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&workloadv1.ListResponse{
		Workloads:   converter.MakeWorkloadsPb(workloads),
		Deployments: converter.MakeDeploymentsPb(deploymentList.Items),
	}), nil
}
