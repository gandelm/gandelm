package catalog

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/gandelm/gandelm/api/v1"
	catalogv1 "github.com/gandelm/gandelm/generated/protocol/catalog/v1"
	"github.com/gandelm/gandelm/internal/container"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CatalogService struct {
	container container.Containerer
}

// Create implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Create(ctx context.Context, request *connect.Request[catalogv1.CreateRequest]) (*connect.Response[catalogv1.CreateResponse], error) {
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
func (c *CatalogService) List(ctx context.Context, request *connect.Request[catalogv1.ListRequest]) (*connect.Response[catalogv1.ListResponse], error) {
	catalog := v1.GandelmCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "1-2-0-sample",
			Namespace: c.container.Config().Namespace(),
		},
	}

	spec := v1.GandelmCatalogSpec{
		ID:          "1-2-0-sample",
		Name:        "sample",
		Version:     "1.2.0",
		Description: "This is a sample catalog",
		Endpoint:    "https://endpoint.gandelm.com",
		Entrypoint:  "https://entrypoint.gandelm.com",
	}
	catalog.Spec = spec

	if err := c.container.Kubernetes().Create(ctx, &catalog); err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.ListResponse{}), nil
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
