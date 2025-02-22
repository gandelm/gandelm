package catalog

import (
	"context"

	"connectrpc.com/connect"
	catalogv1 "github.com/gandelm/gandelm/generated/protocol/catalog/v1"
	"github.com/gandelm/gandelm/internal/container"
)

type CatalogService struct {
	container container.Containerer
}

// Create implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Create(context.Context, *connect.Request[catalogv1.CreateRequest]) (*connect.Response[catalogv1.CreateResponse], error) {
	panic("unimplemented")
}

// Delete implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Delete(context.Context, *connect.Request[catalogv1.DeleteRequest]) (*connect.Response[catalogv1.DeleteResponse], error) {
	panic("unimplemented")
}

// Get implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Get(context.Context, *connect.Request[catalogv1.GetRequest]) (*connect.Response[catalogv1.GetResponse], error) {
	panic("unimplemented")
}

// List implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) List(context.Context, *connect.Request[catalogv1.ListRequest]) (*connect.Response[catalogv1.ListResponse], error) {
	panic("unimplemented")
}

// Update implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Update(context.Context, *connect.Request[catalogv1.UpdateRequest]) (*connect.Response[catalogv1.UpdateResponse], error) {
	panic("unimplemented")
}

func NewCatalogService(
	container container.Containerer,
) *CatalogService {
	return &CatalogService{
		container: container,
	}
}
