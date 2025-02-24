package catalog

import (
	"context"
	"strings"

	"connectrpc.com/connect"
	v1 "github.com/gandelm/gandelm/api/v1"
	catalogv1 "github.com/gandelm/gandelm/generated/protocol/catalog/v1"
	"github.com/gandelm/gandelm/internal/container"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
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

	id := strings.ReplaceAll(version, ".", "-") + "-" + name
	id = strings.ToLower(id)

	catalog := v1.GandelmCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Name:       id,
			Namespace:  c.container.Config().Namespace(),
			Finalizers: []string{"gandelm.com/finalizer"},
		},
	}

	spec := v1.GandelmCatalogSpec{
		ID:          id,
		Name:        name,
		Version:     version,
		Description: description,
		Priority:    priority,
	}
	catalog.Spec = spec

	if err := c.container.Kubernetes().Create(ctx, &catalog); err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.CreateResponse{
		Catalog: &catalogv1.Catalog{
			Id:          id,
			Name:        name,
			Version:     version,
			Description: description,
			Priority:    priority,
		},
	}), nil
}

// Delete implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Delete(ctx context.Context, request *connect.Request[catalogv1.DeleteRequest]) (*connect.Response[catalogv1.DeleteResponse], error) {
	if err := c.container.Kubernetes().Delete(ctx, &v1.GandelmCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Msg.GetCatalogId(),
			Namespace: c.container.Config().Namespace(),
		},
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.DeleteResponse{}), nil
}

// Get implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Get(ctx context.Context, request *connect.Request[catalogv1.GetRequest]) (*connect.Response[catalogv1.GetResponse], error) {
	catalog := &v1.GandelmCatalog{}
	if err := c.container.Kubernetes().Get(ctx, types.NamespacedName{
		Namespace: c.container.Config().Namespace(),
		Name:      request.Msg.GetCatalogId(),
	}, catalog); err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.GetResponse{
		Catalog: &catalogv1.Catalog{
			Id:          catalog.Spec.ID,
			Name:        catalog.Spec.Name,
			Version:     catalog.Spec.Version,
			Description: catalog.Spec.Description,
			Priority:    catalog.Spec.Priority,
		},
	}), nil
}

// List implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) List(ctx context.Context, request *connect.Request[catalogv1.ListRequest]) (*connect.Response[catalogv1.ListResponse], error) {
	catalogs := &v1.GandelmCatalogList{}
	if err := c.container.Kubernetes().List(ctx, catalogs); err != nil {
		return nil, err
	}

	response := &catalogv1.ListResponse{Catalogs: []*catalogv1.Catalog{}}
	for _, catalog := range catalogs.Items {
		response.Catalogs = append(response.Catalogs, &catalogv1.Catalog{
			Id:          catalog.Spec.ID,
			Name:        catalog.Spec.Name,
			Version:     catalog.Spec.Version,
			Description: catalog.Spec.Description,
			Priority:    catalog.Spec.Priority,
		})
	}

	return connect.NewResponse(response), nil
}

// Update implements catalogv1connect.CatalogServiceHandler.
func (c *CatalogService) Update(ctx context.Context, request *connect.Request[catalogv1.UpdateRequest]) (*connect.Response[catalogv1.UpdateResponse], error) {
	priority := request.Msg.GetCatalog().GetPriority()
	description := request.Msg.GetCatalog().GetDescription()
	version := request.Msg.GetCatalog().GetVersion()
	name := request.Msg.GetCatalog().GetName()

	id := strings.ReplaceAll(version, ".", "-") + "-" + name
	id = strings.ToLower(id)

	catalog := v1.GandelmCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Name:      id,
			Namespace: c.container.Config().Namespace(),
		},
	}

	spec := v1.GandelmCatalogSpec{
		ID:          id,
		Name:        name,
		Version:     version,
		Description: description,
		Priority:    priority,
	}
	catalog.Spec = spec

	if err := c.container.Kubernetes().Create(ctx, &catalog); err != nil {
		return nil, err
	}

	return connect.NewResponse(&catalogv1.UpdateResponse{
		Catalog: &catalogv1.Catalog{
			Id:          id,
			Name:        name,
			Version:     version,
			Description: description,
			Priority:    priority,
		},
	}), nil
}
