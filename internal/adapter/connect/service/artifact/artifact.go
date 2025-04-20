package artifact

import (
	"context"

	"connectrpc.com/connect"
	artifactv1 "github.com/gandelm/gandelm/generated/protocol/artifact/v1"
	"github.com/gandelm/gandelm/generated/protocol/artifact/v1/artifactv1connect"
	"github.com/gandelm/gandelm/internal/container"
)

var _ artifactv1connect.ArtifactServiceHandler = (*ArtifactService)(nil)

func NewArtifactService(container container.Containerer) *ArtifactService {
	return &ArtifactService{container: container}
}

type ArtifactService struct {
	container container.Containerer
}

// Create implements artifactv1connect.ArtifactServiceHandler.
func (a *ArtifactService) Create(context.Context, *connect.Request[artifactv1.CreateRequest]) (*connect.Response[artifactv1.CreateResponse], error) {
	panic("unimplemented")
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
func (a *ArtifactService) List(context.Context, *connect.Request[artifactv1.ListRequest]) (*connect.Response[artifactv1.ListResponse], error) {
	panic("unimplemented")
}

// Update implements artifactv1connect.ArtifactServiceHandler.
func (a *ArtifactService) Update(context.Context, *connect.Request[artifactv1.UpdateRequest]) (*connect.Response[artifactv1.UpdateResponse], error) {
	panic("unimplemented")
}
