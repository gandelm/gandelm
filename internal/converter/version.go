package converter

import (
	v1 "github.com/gandelm/gandelm/api/v1"
	versionv1 "github.com/gandelm/gandelm/generated/protocol/version/v1"
	"github.com/gandelm/gandelm/internal/core/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MakeVersionPb(version *entity.Version) *versionv1.Version {
	return &versionv1.Version{
		Id:            version.ID,
		BasePriority:  version.BasePriority,
		BaseArtifacts: []*versionv1.BaseArtifact{},
		CreatedAt:     timestamppb.New(version.CreatedAt),
		UpdatedAt:     timestamppb.New(version.UpdatedAt),
	}
}

func MakeVersionsPb(versions entity.Versions) []*versionv1.Version {
	results := make([]*versionv1.Version, 0, len(versions))
	for _, version := range versions {
		results = append(results, MakeVersionPb(version))
	}

	return results
}

func MakeVersion(version *v1.GandelmVersion) *entity.Version {
	return &entity.Version{
		ID:            version.Name,
		BaseArtifacts: map[string]string{},
		BasePriority:  version.Spec.BasePriority,
		CreatedAt:     version.Spec.CreatedAt.Time,
		UpdatedAt:     version.Spec.UpdatedAt.Time,
	}
}

func MakeVersions(versionList *v1.GandelmVersionList) entity.Versions {
	results := make(entity.Versions, 0, len(versionList.Items))
	for _, version := range versionList.Items {
		results = append(results, MakeVersion(&version))
	}

	return results
}
