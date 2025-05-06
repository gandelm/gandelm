package converter

import (
	v1 "github.com/gandelm/gandelm/api/v1"
	catalogv1 "github.com/gandelm/gandelm/generated/protocol/catalog/v1"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MakeCatalogPb(catalog *entity.Catalog) *catalogv1.Catalog {
	return &catalogv1.Catalog{
		Id:          catalog.ID.String(),
		WorkloadId:  catalog.WorkloadID,
		Name:        catalog.Name,
		Version:     catalog.Version,
		Description: catalog.Description,
		Priority:    uint32(catalog.Priority),
		Labels:      catalog.Labels,
		CreatedAt:   timestamppb.New(catalog.CreatedAt),
		UpdatedAt:   timestamppb.New(catalog.UpdatedAt),
	}
}

func MakeCatalogsPb(catalogs entity.Catalogs) []*catalogv1.Catalog {
	catalogsPb := make([]*catalogv1.Catalog, 0, len(catalogs))
	for _, catalog := range catalogs {
		catalogsPb = append(catalogsPb, MakeCatalogPb(catalog))
	}

	return catalogsPb
}

func MakeCatalog(catalog *v1.GandelmCatalog) *entity.Catalog {
	return &entity.Catalog{
		ID:           uuid.MustParse(catalog.Name),
		WorkloadID:   catalog.Spec.ID,
		Name:         catalog.Spec.Name,
		Version:      catalog.Spec.Version,
		Description:  catalog.Spec.Description,
		Priority:     int(catalog.Spec.Priority),
		Labels:       catalog.Spec.Labels,
		ArtifactTags: catalog.Spec.ArtifactTags,
		CreatedAt:    catalog.Spec.CreatedAt.Time,
		UpdatedAt:    catalog.Spec.UpdatedAt.Time,
	}
}

func MakeCatalogs(catalogs v1.GandelmCatalogList) entity.Catalogs {
	results := make(entity.Catalogs, 0, len(catalogs.Items))
	for _, catalog := range catalogs.Items {
		results = append(results, MakeCatalog(&catalog))
	}

	return results
}
