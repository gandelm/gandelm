package version

import (
	"context"
	"time"

	"connectrpc.com/connect"
	versionv1 "github.com/gandelm/gandelm/generated/protocol/version/v1"
	"github.com/gandelm/gandelm/generated/protocol/version/v1/versionv1connect"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"github.com/gandelm/gandelm/internal/provider"
)

var _ versionv1connect.VersionServiceHandler = (*VersionService)(nil)

func NewVersionService(container container.Containerer) *VersionService {
	return &VersionService{
		versionRepository: provider.NewVersionRWRepository(container),
	}
}

type VersionService struct {
	versionRepository repository.VersionRWRepository
}

// Create implements versionv1connect.VersionServiceHandler.
func (v *VersionService) Create(ctx context.Context, request *connect.Request[versionv1.CreateRequest]) (*connect.Response[versionv1.CreateResponse], error) {
	id := request.Msg.GetId()
	// baseArtifacts := request.Msg.GetBaseArtifacts()
	basePriority := request.Msg.GetBasePriority()

	now := time.Now()
	version := &entity.Version{
		ID:           id,
		BasePriority: basePriority,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := v.versionRepository.Create(ctx, version); err != nil {
		return nil, err
	}

	return connect.NewResponse(&versionv1.CreateResponse{
		Version: &versionv1.Version{},
	}), nil
}

// Delete implements versionv1connect.VersionServiceHandler.
func (v *VersionService) Delete(context.Context, *connect.Request[versionv1.DeleteRequest]) (*connect.Response[versionv1.DeleteResponse], error) {
	panic("unimplemented")
}

// Get implements versionv1connect.VersionServiceHandler.
func (v *VersionService) Get(context.Context, *connect.Request[versionv1.GetRequest]) (*connect.Response[versionv1.GetResponse], error) {
	panic("unimplemented")
}

// List implements versionv1connect.VersionServiceHandler.
func (v *VersionService) List(ctx context.Context, request *connect.Request[versionv1.ListRequest]) (*connect.Response[versionv1.ListResponse], error) {
	versions, err := v.versionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&versionv1.ListResponse{
		Versions: converter.MakeVersionsPb(versions),
	}), nil
}
