package artifact

import (
	"context"

	"connectrpc.com/connect"
	artifactv1 "github.com/gandelm/gandelm/generated/protocol/artifact/v1"
	"github.com/gandelm/gandelm/generated/protocol/artifact/v1/artifactv1connect"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"github.com/gandelm/gandelm/internal/provider"
)

var _ artifactv1connect.ArtifactServiceHandler = (*ArtifactService)(nil)

func NewArtifactService(container container.Containerer) *ArtifactService {
	return &ArtifactService{
		repo: provider.NewArtifactRepository(container),
	}
}

type ArtifactService struct {
	repo repository.ArtifactRWRepository
}

// Create implements artifactv1connect.ArtifactServiceHandler.
func (a *ArtifactService) Create(ctx context.Context, req *connect.Request[artifactv1.CreateRequest]) (*connect.Response[artifactv1.CreateResponse], error) {
	id := req.Msg.GetId()
	title := req.Msg.GetTitle()
	desc := req.Msg.GetDescription()
	t := req.Msg.GetType()

	err := a.repo.Create(ctx, &entity.Artifact{
		ID:          entity.ArtifactID(id),
		Title:       title,
		Description: desc,
		Type:        entity.ArtifactType(t),
	})

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&artifactv1.CreateResponse{
		Artifact: &artifactv1.Artifact{
			Id:          id,
			Title:       title,
			Description: desc,
			Type:        t,
		},
	}), nil
}

// Delete implements artifactv1connect.ArtifactServiceHandler.
func (a *ArtifactService) Delete(context.Context, *connect.Request[artifactv1.DeleteRequest]) (*connect.Response[artifactv1.DeleteResponse], error) {
	panic("unimplemented")
}

// Get implements artifactv1connect.ArtifactServiceHandler.
func (a *ArtifactService) Get(context.Context, *connect.Request[artifactv1.GetRequest]) (*connect.Response[artifactv1.GetResponse], error) {
	panic("unimplemented")
}

// List implements artifactv1connect.ArtifactServiceHandler.
func (a *ArtifactService) List(ctx context.Context, req *connect.Request[artifactv1.ListRequest]) (*connect.Response[artifactv1.ListResponse], error) {
	artifacts, err := a.repo.FindAll(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&artifactv1.ListResponse{
		Artifacts: converter.MakeArtifactsPb(artifacts),
	}), nil
}

// Update implements artifactv1connect.ArtifactServiceHandler.
func (a *ArtifactService) Update(context.Context, *connect.Request[artifactv1.UpdateRequest]) (*connect.Response[artifactv1.UpdateResponse], error) {
	panic("unimplemented")
}
