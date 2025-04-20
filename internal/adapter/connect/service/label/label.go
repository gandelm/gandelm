package label

import (
	"context"

	"connectrpc.com/connect"
	labelv1 "github.com/gandelm/gandelm/generated/protocol/label/v1"
	"github.com/gandelm/gandelm/generated/protocol/label/v1/labelv1connect"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/gandelm/gandelm/internal/core/domain/repository"
	"github.com/gandelm/gandelm/internal/provider"
)

var _ labelv1connect.LabelServiceHandler = (*LabelService)(nil)

func NewLabelService(container container.Containerer) *LabelService {
	return &LabelService{
		labelRepository: provider.NewLabelRWRepository(container),
	}
}

type LabelService struct {
	labelRepository repository.LabelRWRepository
}

// Create implements labelv1connect.LabelServiceHandler.
func (a *LabelService) Create(ctx context.Context, request *connect.Request[labelv1.CreateRequest]) (*connect.Response[labelv1.CreateResponse], error) {
	id := request.Msg.GetId()
	title := request.Msg.GetTitle()
	desc := request.Msg.GetDescription()
	color := request.Msg.GetColor()

	err := a.labelRepository.Create(ctx, &entity.Label{
		ID:          id,
		Title:       title,
		Description: desc,
		Color:       color,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&labelv1.CreateResponse{
		Label: &labelv1.Label{
			Id:          id,
			Title:       title,
			Description: desc,
			Color:       color,
		},
	}), nil
}

// Delete implements labelv1connect.LabelServiceHandler.
func (a *LabelService) Delete(ctx context.Context, request *connect.Request[labelv1.DeleteRequest]) (*connect.Response[labelv1.DeleteResponse], error) {
	if err := a.labelRepository.Delete(ctx, request.Msg.GetLabelId()); err != nil {
		return nil, err
	}

	return connect.NewResponse(&labelv1.DeleteResponse{}), nil
}

// Get implements labelv1connect.LabelServiceHandler.
func (a *LabelService) Get(ctx context.Context, request *connect.Request[labelv1.GetRequest]) (*connect.Response[labelv1.GetResponse], error) {
	label, err := a.labelRepository.Find(ctx, request.Msg.GetLabelId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&labelv1.GetResponse{
		Label: converter.MakeLabelPb(label),
	}), nil
}

// List implements labelv1connect.LabelServiceHandler.
func (a *LabelService) List(ctx context.Context, _ *connect.Request[labelv1.ListRequest]) (*connect.Response[labelv1.ListResponse], error) {
	labels, err := a.labelRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&labelv1.ListResponse{
		Labels: converter.MakeLabelsPb(labels),
	}), nil
}
