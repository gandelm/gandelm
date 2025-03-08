package catalog

import (
	"context"

	"connectrpc.com/connect"
	catalogv1 "github.com/gandelm/gandelm/generated/protocol/catalog/v1"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/converter"
	"github.com/gandelm/gandelm/internal/core/usecase"
	"github.com/gandelm/gandelm/internal/provider"
	"github.com/google/uuid"
)

func NewCatalogService(container container.Containerer) *CatalogService {
	return &CatalogService{container: container}
}

type CatalogService struct {
	container container.Containerer
}

// Create implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Create(ctx context.Context, request *connect.Request[catalogv1.CreateRequest]) (*connect.Response[catalogv1.CreateResponse], error) {
	priority := request.Msg.GetPriority()
	description := request.Msg.GetDescription()
	version := request.Msg.GetVersion()
	name := request.Msg.GetName()

	o, err := provider.NewCatalogCreator(c.container).Execute(ctx, &usecase.CatalogCreateInput{
		Priority:    int(priority),
		Description: description,
		Version:     version,
		Name:        name,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.CreateResponse{
		Catalog: converter.MakeCatalogPb(o.Catalog),
	}), nil
}

// Delete implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Delete(ctx context.Context, request *connect.Request[catalogv1.DeleteRequest]) (*connect.Response[catalogv1.DeleteResponse], error) {
	id := request.Msg.GetCatalogId()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if err := provider.NewCatalogRWRepository(c.container).Delete(ctx, uuid); err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.DeleteResponse{}), nil
}

// Get implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Get(ctx context.Context, request *connect.Request[catalogv1.GetRequest]) (*connect.Response[catalogv1.GetResponse], error) {
	id := request.Msg.GetCatalogId()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	catalog, err := provider.NewCatalogRORepository(c.container).Find(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.GetResponse{
		Catalog: converter.MakeCatalogPb(catalog),
	}), nil
}

// List implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) List(ctx context.Context, request *connect.Request[catalogv1.ListRequest]) (*connect.Response[catalogv1.ListResponse], error) {
	catalogs, err := provider.NewCatalogRORepository(c.container).FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.ListResponse{
		Catalogs: converter.MakeCatalogsPb(catalogs),
	}), nil
}

// Update implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Update(ctx context.Context, request *connect.Request[catalogv1.UpdateRequest]) (*connect.Response[catalogv1.UpdateResponse], error) {
	catalogpb := request.Msg.GetCatalog()
	uuid, err := uuid.Parse(catalogpb.GetId())
	if err != nil {
		return nil, err
	}

	o, err := provider.NewCatalogUpdator(c.container).Execute(ctx, &usecase.CatalogUpdateInput{
		ID:          uuid,
		Priority:    int(catalogpb.GetPriority()),
		Description: catalogpb.GetDescription(),
		Version:     catalogpb.GetVersion(),
		Name:        catalogpb.GetName(),
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.UpdateResponse{
		Catalog: converter.MakeCatalogPb(o.Catalog),
	}), nil
}
